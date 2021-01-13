package wordpress

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/importer"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/google/uuid"
	"github.com/gookit/color"
	"net/url"
)

var (
	resource = "posts"
	template = "test"
	fieldUuid = "39ca0ea0-c911-4eaa-b6e0-67dfd99e1225"
)

type Convertor interface {

}

type Convert struct {
	XML     WpXml
	Failed  []FailedImport
	store   *models.Store
	authors []domain.User
	owner domain.User
}

type FailedImport struct {
	Item  Item
	Error error
}

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
		Failed: nil,
		store:  s,
		owner: owner,
	}, nil
}

func (c *Convert) Import() {
	go c.populateAuthors()
	c.populatePosts()
}

func (c *Convert) parseImages(content string) (string, error) {

	parsed, err := importer.ParseHTML(content, func(url string) (string, error) {
		return c.DownloadFile(url)
	})

	if err != nil {
		return "", err
	}

	return parsed, nil
}

// DownloadFile
func (c *Convert) DownloadFile(url string) (string, error) {
	file, err := importer.DownloadFile(url)
	if err != nil {
		return "", err
	}

	media, err := c.store.Media.Upload(file, c.owner.Token)
	if err != nil {
		return "", err
	}

	return media.Url, nil
}

// populatePosts
func (c *Convert) populatePosts() {

	for _, item := range c.XML.Channel.Items {

		link, err := parseLink(item.Link)
		if err != nil {
			c.fail(item, err)
		}

		userId, err := c.findAuthor(item)
		if err != nil {
			c.fail(item, err)
		}

		content, err := c.parseImages(item.Content)
		if err != nil {
			c.fail(item, err)
		}

		uuid, err := uuid.Parse(fieldUuid)
		if err != nil {
			c.fail(item, err)
		}

		color.Green.Println(item.Categories)

		post := domain.PostCreate{
			Post: domain.Post{
				Slug:              link,
				Title:             item.Title,
				Status:            getStatus(item.Status),
				Resource:          &resource,
				PageTemplate:      template,
				PageLayout:        "",
				CodeInjectionHead: nil,
				CodeInjectionFoot: nil,
				UserId:            userId,
				PublishedAt:       &item.PubDatetime,
				CreatedAt:         &item.PostDatetime,
				UpdatedAt:         &item.PostDatetime,
				SeoMeta:           domain.PostSeoMeta{},
			},
			Author:   0,
			Category: nil,
			Fields:   []domain.PostField{
				{
					UUID:          uuid,
					Type:          "richtext",
					Name:          "content",
					OriginalValue:  domain.FieldValue(content),
				},
			},
		}

		_, err = c.store.Posts.Create(&post)
		if err != nil {
			c.fail(item, err)
		}
	}
}


func getCategory(item Item) {

}

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
		if !exists {
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

			}

			c.authors = append(c.authors, u)
			// TODO: Send email with new password
		}
	}
}

func (c *Convert) fail(item Item, err error) {
	c.Failed = append(c.Failed, FailedImport{
		Item:  item,
		Error: err,
	})
}

func getStatus(status string) string {
	if status == "publish" {
		return "published"
	}
	return status
}

func parseLink(link string) (string, error) {
	const op = "WordpressConvertor.parseLink"
	u, err := url.Parse(link)
	if err != nil {
		return "", &errors.Error{Code: errors.INVALID, Message: "Unable to parse post link", Operation: op, Err: err}
	}
	return u.Path, nil
}
