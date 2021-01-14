package wordpress

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/importer"
	"github.com/ainsleyclark/verbis/api/models"
	"mime/multipart"
)

// TODO: This needs to be dynamic
var (
	resource  = "posts"
	layout    = "main"
	template  = "test"
	fieldUuid = "39ca0ea0-c911-4eaa-b6e0-67dfd99e1225"
)

type Convert struct {
	XML     WpXml
	Failed  Failures
	store   *models.Store
	authors []domain.User
	owner   domain.User
}

// Failed import defines the errors that occured when importing
// multiple entities into Verbis.
type Failures struct {
	Posts FailedPosts
	Authors FailedAuthors
}

type FailedPosts []FailedPost

type FailedPost struct {
	Post  Item
	Media []FailedMedia
	Error error
}

// Append
//
// Accepts the failed post (Item), the array of FailedMedia that could
// not be parsed or uploaded. And the original error to be sent
// back after the import.
func (f FailedPosts) Append(item Item, media []FailedMedia, err error) {
	f = append(f, FailedPost{
		Post:  item,
		Media: media,
		Error: err,
	})
}

type FailedMedia struct {
	Url   string
	Error error
}

type FailedAuthors []FailedAuthor

func (f FailedAuthors) Append(fName string, lName string, email string, err error) {
	f = append(f, FailedAuthor{
		FirstName: fName,
		LastName: lName,
		Email:    email,
	})
}

type FailedAuthor struct {
	FirstName string
	LastName  string
	Email     string
	Error     error
}

// New - Construct
func New(xmlPath string, s *models.Store) (*Convert, error) {
	wp := NewWordpressXml()
	err := wp.ReadFile(xmlPath)
	if err != nil {
		return nil, err
	}

	owner, err := s.User.GetOwner()
	if err != nil {
		return nil, err
	}

	return &Convert{
		XML:    wp,
		Failed: Failures{},
		store:  s,
		owner:  owner,
	}, nil
}

// Import the XML file into Wordpress by populating Authors
// and Posts
func (c *Convert) Import() {
	c.populateAuthors()
	c.populatePosts()
}

// populatePosts
func (c *Convert) populatePosts() {
	const op = "WordpressConvertor.populatePosts"

	for _, item := range c.XML.Channel.Items {

		link, err := importer.ParseLink(item.Link)
		if err != nil {
			c.Failed.Posts.Append(item, nil, err)
		}

		uuid, err := importer.ParseUUID(fieldUuid)
		if err != nil {
			c.Failed.Posts.Append(item, nil, err)
		}

		userId, err := c.findAuthor(item)
		if err != nil {
			c.Failed.Posts.Append(item, nil, err)
		}

		content, failed, err := c.parseImages(item.Content)
		if err != nil {
			c.Failed.Posts.Append(item, failed, err)
		}

		post := domain.PostCreate{
			Post: domain.Post{
				Slug:         link,
				Title:        item.Title,
				Status:       getStatus(item.Status),
				Resource:     &resource,
				PageTemplate: template,
				PageLayout:   layout,
				UserId:       userId,
				PublishedAt:  &item.PubDatetime,
				CreatedAt:    &item.PostDatetime,
				UpdatedAt:    &item.PostDatetime,
				SeoMeta:      c.getSeoMeta(item.Title, item.Meta),
			},
			Author: userId,
			Fields: []domain.PostField{
				{
					UUID:          uuid,
					Type:          "richtext",
					Name:          "content",
					OriginalValue: domain.FieldValue(content),
				},
			},
		}

		category, err := c.getCategory(item.Categories)
		if err != nil && errors.Code(err) != errors.NOTFOUND {
			c.Failed.Posts.Append(item, nil, err)
		}

		if err == nil {
			post.Category = &category.Id
		}

		_, err = c.store.Posts.Create(&post)
		if err != nil {
			c.Failed.Posts.Append(item,nil, err)
		}
	}
}

// parseImages
//
// Accepts a HTML document as a string and uses the ParseHTML function to
// loop over the images, upload them and modify the contents of the HTML
// file If a media item failed to be uploaded to the media library
// or a the file could not be downloaded (such as a 404) the
// media item will be appended to the FailedMedia array.
//
// Returns the modified HTML file, the FailedMedia array and an error
// if there was a problem parsing the HTML.
func (c *Convert) parseImages(content string) (string, []FailedMedia, error) {
	var failed []FailedMedia
	parsed, err := importer.ParseHTML(content, func(file *multipart.FileHeader, url string, err error) (string, error) {
		if err != nil {
			failed = append(failed, FailedMedia{Url: url, Error: err})
		}

		media, err := c.store.Media.Upload(file, c.owner.Token)
		if err != nil {
			failed = append(failed, FailedMedia{Url: url, Error: err})
		}

		return media.Url, nil
	})

	if err != nil {
		return "", failed, err
	}

	return parsed, failed, nil
}

// getCategory
//
// Converts a 'Wordpress' category into a domain.Category
//
// Returns found category if it already exists.
// Returns newly created category if it doesnt exist.
// Returns errors.NOTFOUND if not category is attached to the post.
func (c *Convert) getCategory(categories []Category) (domain.Category, error) {
	const op = "WordpressConvertor.getCategory"

	if len(categories) == 0 {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "No category is attached to the post type.", Operation: op, Err: fmt.Errorf("no category found")}
	}

	wp := categories[0]

	if c.store.Categories.ExistsBySlug(wp.URLSlug) {
		return c.store.Categories.GetBySlug(wp.URLSlug)
	}

	return c.store.Categories.Create(&domain.Category{
		Slug:     wp.URLSlug,
		Name:     wp.DisplayName,
		Resource: resource,
	})
}

// getSeoMeta
//
// Constructs domain.PostSeoMeta and attaches meta titles and
// meta descriptions if the 'Yoast' plugin exists in
// 'Wordpress'.
func (c *Convert) getSeoMeta(title string, meta []Meta) domain.PostSeoMeta {
	m := domain.PostSeoMeta{
		Meta: &domain.PostMeta{
			Title: title,
			Twitter: domain.PostTwitter{
				Title: title,
			},
			Facebook: domain.PostFacebook{
				Title: title,
			},
		},
	}

	for _, v := range meta {
		if v.MetaKey == "_yoast_wpseo_metadesc" {
			m.Meta.Description = v.MetaValue
			m.Meta.Twitter.Description = v.MetaValue
			m.Meta.Facebook.Description = v.MetaValue
		}
	}

	return m
}

// findAuthor
//
// Looks through the array of authors attached to the Convert
// struct and returns the Author ID.
//
// Returns owner ID if there was an error obtaining the Wordpress
// authors or no author exists in the Convert authors array.
func (c *Convert) findAuthor(item Item) (int, error) {
	author, err := c.XML.AuthorForLogin(item.Creator)
	if err != nil {
		return c.owner.Id, nil
	}

	for _, v := range c.authors {
		if v.Id == author.AuthorID {
			return v.Id, nil
		}
	}

	return c.owner.Id, nil
}

func (c *Convert) populateAuthors() {
	for _, v := range c.XML.Channel.Authors {
		exists := c.store.User.ExistsByEmail(v.AuthorEmail)
		if exists {
			continue
		}

		password := encryption.CreatePassword()
		user := &domain.UserCreate{
			User: domain.User{
				UserPart: domain.UserPart{
					FirstName: v.AuthorFirstName,
					LastName:  v.AuthorLastName,
					Email:     v.AuthorEmail,
					Role: domain.UserRole{
						Id: 2,
					},
				},
			},
			Password:        password,
			ConfirmPassword: password,
		}

		u, err := c.store.User.Create(user)
		if err != nil {
			c.Failed.Authors.Append(v.AuthorFirstName, v.AuthorLastName, v.AuthorEmail, err)
		}

		c.authors = append(c.authors, u)
		// TODO: Send email with new password
	}
}

// getStatus
//
// Converts the Wordpress status to Verbis specific status's.
func getStatus(status string) string {
	if status == "publish" {
		return "published"
	}
	return status
}
