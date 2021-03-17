// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/google/uuid"
	"strings"
)

// PostsRepository defines methods for Posts to interact with the database
type PostsRepository interface {
	Get(meta params.Params, layout bool, resource string, status string) (domain.PostData, int, error)
	GetByID(id int, layout bool) (domain.PostDatum, error)
	GetBySlug(slug string) (domain.PostDatum, error)
	Create(p *domain.PostCreate) (domain.PostDatum, error)
	Update(p *domain.PostCreate) (domain.PostDatum, error)
	Delete(id int) error
	Exists(id int) bool
	ExistsBySlug(slug string) bool
	Total() (int, error)
}

// TODO
// 		- Need to validate page templates, if its not in the available templates, needs to be a 400.

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

// PostStore defines the data layer for Posts
type PostStore struct {
	*StoreCfgOld
	seoMetaModel    SeoMetaRepository
	userModel       UserRepository
	categoriesModel CategoryRepository
	fieldsModel     FieldsRepository
	options         domain.Options
}

// newPosts - Construct
func newPosts(cfg *StoreCfgOld) *PostStore {
	return &PostStore{
		StoreCfgOld:     cfg,
		seoMetaModel:    newSeoMeta(cfg),
		userModel:       newUser(cfg),
		categoriesModel: newCategories(cfg),
		fieldsModel:     newFields(cfg),
		options:         cfg.Options.GetStruct(),
	}
}

type PostRaw struct {
	domain.Post
	Author   domain.User     `db:"author"`
	Category domain.Category `db:"category"`
	Field    struct {
		Id            int        `db:"field_id"` //nolint
		PostId        int        `db:"post_id"`  //nolint
		UUID          *uuid.UUID `db:"uuid"`
		Type          string     `db:"type"`
		Name          string     `db:"name"`
		Key           string     `db:"field_key"`
		OriginalValue string     `db:"value" json:"value"`
	} `db:"field"`
}

func (s PostStore) getQuery(query string) string {
	// TODO CATEGORIES TIMESTAMPS
	return fmt.Sprintf(`SELECT posts.*, post_options.seo 'options.seo', post_options.meta 'options.meta',
       users.id as 'author.id', users.uuid as 'author.uuid', users.first_name 'author.first_name', users.last_name 'author.last_name', users.email 'author.email', users.website 'author.website', users.facebook 'author.facebook', users.twitter 'author.twitter', users.linked_in 'author.linked_in',
       users.instagram 'author.instagram', users.biography 'author.biography', users.profile_picture_id 'author.profile_picture_id', users.updated_at 'author.updated_at', users.created_at 'author.created_at',
       roles.id 'author.roles.id', roles.name 'author.roles.name', roles.description 'author.roles.description',
       pf.uuid 'field.uuid',
       CASE WHEN categories.id IS NULL THEN 0 ELSE categories.id END AS 'category.id',
       CASE WHEN categories.uuid IS NULL THEN '' ELSE categories.uuid END AS 'category.uuid',
       CASE WHEN categories.slug IS NULL THEN '' ELSE categories.slug END AS 'category.slug',
       CASE WHEN categories.name IS NULL THEN '' ELSE categories.name END AS 'category.name',
       CASE WHEN categories.description IS NULL THEN '' ELSE categories.description END AS 'category.description',
       CASE WHEN categories.resource IS NULL THEN '' ELSE categories.resource END AS 'category.resource',
       CASE WHEN categories.parent_id IS NULL THEN 0 ELSE categories.parent_id END AS 'category.parent_id',
       CASE WHEN pf.id IS NULL THEN 0 ELSE pf.id END AS 'field.field_id',
       CASE WHEN pf.type IS NULL THEN "" ELSE pf.type END AS 'field.type',
       CASE WHEN pf.field_key IS NULL THEN "" ELSE pf.field_key END AS 'field.field_key',
       CASE WHEN pf.name IS NULL THEN "" ELSE pf.name END AS 'field.name',
       CASE WHEN pf.value IS NULL THEN "" ELSE pf.value END AS 'field.value'
FROM (%s) posts
      LEFT JOIN post_options ON posts.id = post_options.post_id
      LEFT JOIN users ON posts.user_id = users.id
      INNER JOIN user_roles ON users.id = user_roles.user_id
      LEFT JOIN roles ON user_roles.role_id = roles.id
      LEFT JOIN post_categories pc on posts.id = pc.post_id
      LEFT JOIN categories on pc.category_id = categories.id
      LEFT JOIN post_fields pf on posts.id = pf.post_id`, query)
}

func (s *PostStore) Get(meta params.Params, layout bool, resource, status string) (domain.PostData, int, error) {
	const op = "PostsRepository.Get"

	q := "SELECT * FROM posts"
	countQ := "SELECT COUNT(*) FROM posts"

	// Apply filters to total and original query
	filter, err := filterRows(s.DB, meta.Filters, "posts")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Get by resource
	if resource != "all" && resource != "" {
		if len(meta.Filters) > 0 {
			q += " AND"
			countQ += " AND"
		} else {
			q += " WHERE"
			countQ += " WHERE"
		}

		// If the resource is pages or a resource
		resourceQ := ""
		if resource == "pages" {
			resourceQ = " posts.resource IS NULL"
		} else {
			resourceQ = fmt.Sprintf(" posts.resource = '%s'", resource)
		}

		q += resourceQ
		countQ += resourceQ
	}

	// Get Status
	if status != "" {
		if resource != "" {
			q += " AND"
			countQ += " AND"
		} else {
			q += " WHERE"
			countQ += " WHERE"
		}
		q += fmt.Sprintf(" posts.status = '%s'", status)
		countQ += fmt.Sprintf(" posts.status = '%s'", status)
	}

	// Apply order
	if meta.OrderBy != "" {
		q += fmt.Sprintf(" ORDER BY posts.%s %s", meta.OrderBy, meta.OrderDirection)
	}

	// Apply pagination
	if !meta.LimitAll {
		q += fmt.Sprintf(" LIMIT %v OFFSET %v", meta.Limit, (meta.Page-1)*meta.Limit)
	}

	q = s.getQuery(q)

	// Apply order
	if meta.OrderBy != "" {
		q += fmt.Sprintf(" ORDER BY posts.%s %s", meta.OrderBy, meta.OrderDirection)
	}

	var rawPosts []PostRaw

	if err := s.DB.Select(&rawPosts, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get posts", Operation: op, Err: err}
	}

	// Count the total number of posts
	var total int
	if err := s.DB.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of posts", Operation: op, Err: err}
	}

	// Return not found error if no posts are available
	formattedPosts := s.format(rawPosts, layout)
	if len(formattedPosts) == 0 {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No posts available", Operation: op}
	}

	return formattedPosts, total, nil
}

// GetByID returns a post by Id
//
// Returns errors.NOTFOUND if the post was not found by the given Id.
func (s *PostStore) GetByID(id int, layout bool) (domain.PostDatum, error) {
	const op = "PostsRepository.GetByID"

	var p []PostRaw
	err := s.DB.Select(&p, s.getQuery("SELECT * FROM posts WHERE posts.id = ? LIMIT 1"), id)

	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the post with the ID: %d", id), Operation: op, Err: err}
	}

	formatted := s.format(p, layout)
	if len(formatted) == 0 {
		return domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: "post not found", Operation: op, Err: fmt.Errorf("no post found")}
	}

	return formatted[0], nil
}

// GetBySlug returns a a post by slug
//
// Returns errors.NOTFOUND if the post was not found by the given slug.
func (s *PostStore) GetBySlug(slug string) (domain.PostDatum, error) {
	const op = "PostsRepository.GetBySlug"

	var p []PostRaw
	err := s.DB.Select(&p, s.getQuery("SELECT * FROM posts WHERE posts.slug = ? LIMIT 1"), slug)

	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post with the slug %s", slug), Operation: op, Err: err}
	}

	formatted := s.format(p, false)
	if len(formatted) == 0 {
		return domain.PostDatum{}, &errors.Error{Code: errors.NOTFOUND, Message: "post not found", Operation: op, Err: fmt.Errorf("no post found")}
	}

	return formatted[0], nil
}

// Create a new post
// Returns errors.CONFLICT if the the post slug already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function
// could not get the newly created ID.
func (s *PostStore) Create(p *domain.PostCreate) (domain.PostDatum, error) {
	const op = "PostsRepository.Create"

	if err := s.validateURL(p.Slug); err != nil {
		return domain.PostDatum{}, err
	}

	// Check if the author is set assign to owner if not.
	p.UserId = s.checkOwner(*p)

	// TODO: Work out why sql defaults arent working!
	if p.Status == "" {
		p.Status = "draft"
	}

	// Remove any trailing slashes from slug.
	if p.Slug != "/" {
		p.Slug = strings.TrimRight(p.Slug, "/")
	}

	q := "INSERT INTO posts (uuid, slug, title, status, resource, page_template, layout, codeinjection_head, codeinjection_foot, user_id, published_at, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	c, err := s.DB.Exec(q, uuid.New().String(), p.Slug, p.Title, p.Status, p.Resource, p.PageTemplate, p.PageLayout, p.CodeInjectionHead, p.CodeInjectionFoot, p.UserId, p.PublishedAt)
	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the post with the title: %v", p.Title), Operation: op, Err: err}
	}

	id, err := c.LastInsertId()
	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created post ID with the title: %v", p.Title), Operation: op, Err: err}
	}

	post, err := s.GetByID(int(id), true)
	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created post with the title: %v", p.Title), Operation: op, Err: err}
	}

	// Update the categories based on the array of integers that
	// are passed.
	if err := s.categoriesModel.InsertPostCategory(int(id), p.Category); err != nil {
		return domain.PostDatum{}, err
	}

	// Update or create the fields
	if err := s.fieldsModel.UpdateCreate(int(id), p.Fields); err != nil {
		return domain.PostDatum{}, err
	}

	// Update the post meta
	if err := s.seoMetaModel.UpdateCreate(int(id), p.SeoMeta); err != nil {
		return domain.PostDatum{}, err
	}

	return post, nil
}

// Update a post by Id
// Returns errors.NOTFOUND if the post was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *PostStore) Update(p *domain.PostCreate) (domain.PostDatum, error) {
	const op = "PostsRepository.Update"

	oldPost, err := s.GetByID(p.Id, false)
	if err != nil {
		return domain.PostDatum{}, err
	}

	if oldPost.Slug != p.Slug {
		if err := s.validateURL(p.Slug); err != nil {
			return domain.PostDatum{}, err
		}
	}

	// Check if the author is set assign to owner if not.
	p.Author = s.checkOwner(*p)
	p.UserId = p.Author

	// Remove any trailing slashes from slug.
	if p.Slug != "/" {
		p.Slug = strings.TrimRight(p.Slug, "/")
	}

	// Update the posts table with data
	q := "UPDATE posts SET slug = ?, title = ?, status = ?, resource = ?, page_template = ?, layout = ?, codeinjection_head = ?, codeinjection_foot = ?, user_id = ?, published_at = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.DB.Exec(q, p.Slug, p.Title, p.Status, p.Resource, p.PageTemplate, p.PageLayout, p.CodeInjectionHead, p.CodeInjectionFoot, p.UserId, p.PublishedAt, p.Id)
	if err != nil {
		return domain.PostDatum{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the post wuth the title: %v", p.Title), Operation: op, Err: err}
	}

	// Update the categories based on the array of integers that
	// are passed. If the categories
	if err := s.categoriesModel.InsertPostCategory(p.Id, p.Category); err != nil {
		return domain.PostDatum{}, err
	}

	// Update or create the fields
	if err := s.fieldsModel.UpdateCreate(p.Id, p.Fields); err != nil {
		return domain.PostDatum{}, err
	}

	post, err := s.GetByID(p.Id, true)
	if err != nil {
		return domain.PostDatum{}, err
	}

	// Update the post meta
	if err := s.seoMetaModel.UpdateCreate(p.Id, p.SeoMeta); err != nil {
		return domain.PostDatum{}, err
	}

	// Clear the cache
	cache.Store.Delete(cache.GetPostKey(post.Id))

	return post, nil
}

// Delete post
// Returns errors.NOTFOUND if the post was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *PostStore) Delete(id int) error {
	const op = "PostsRepository.Delete"

	if !s.Exists(id) {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No post exists with the ID: %v", id), Operation: op, Err: fmt.Errorf("no post exists")}
	}

	if _, err := s.DB.Exec("DELETE FROM posts WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete post with the ID: %v", id), Operation: op, Err: err}
	}

	return nil
}

// Total gets the total number of posts
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *PostStore) Total() (int, error) {
	const op = "PostsRepository.Total"
	var total int
	if err := s.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&total); err != nil {
		return -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of posts", Operation: op, Err: err}
	}
	return total, nil
}

// Exists Checks if a post exists by the given slug
func (s *PostStore) Exists(id int) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT id FROM posts WHERE id = ?)", id).Scan(&exists)
	return exists
}

// Exists Checks if a post exists by the given slug
func (s *PostStore) ExistsBySlug(slug string) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT id FROM posts WHERE slug = ?)", slug).Scan(&exists)
	return exists
}

// checkOwner Checks if the author is set or if the author does not exist.
// Returns the owner ID under circumstances.
func (s *PostStore) checkOwner(p domain.PostCreate) int {
	if p.Author == 0 || !s.userModel.Exists(p.Author) {
		owner, err := s.userModel.GetOwner()
		if err != nil {
			logger.Panic(err)
		}
		return owner.Id
	}
	return p.Author
}

// validateURL checks if the url is valid for creating or updating a new
// post.
//
// Returns errors.CONFLICT if the post slug already exists
// Or the slug contains the admin path, .i.e /admin
func (s *PostStore) validateURL(slug string) error {
	const op = "PostsRepository.validateURLvalidateURL"

	if s.ExistsBySlug(slug) {
		return &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the slug %v, already exists", slug), Operation: op}
	}

	slugArr := strings.Split(slug, "/")
	if len(slugArr) > 1 {
		if strings.Contains(slugArr[1], "admin") {
			return &errors.Error{Code: errors.CONFLICT, Message: "Could not create the post, the path /admin is reserved", Operation: op}
		}
	}

	return nil
}

func (s *PostStore) find(posts domain.PostData, id int) bool {
	for _, v := range posts {
		if v.Id == id {
			return true
		}
	}
	return false
}

func (s *PostStore) format(rawPosts []PostRaw, layout bool) domain.PostData {
	var posts = make(domain.PostData, 0)

	for _, v := range rawPosts {
		if !s.find(posts, v.Id) {
			var category domain.Category
			if v.Category.Id != 0 {
				category = v.Category
			}

			p := domain.PostDatum{
				Post:     v.Post,
				Author:   v.Author.HideCredentials(),
				Category: &category,
				Fields:   make(domain.PostFields, 0),
			}

			if layout {
				p.Layout = s.fieldsModel.GetLayout(p)
			}

			p.Permalink = s.getPermalink(&p)
			p.Type = s.postType(&p)

			posts = append(posts, p)
		}

		if v.Field.UUID != nil {
			field := domain.PostField{
				Id:            v.Field.Id,
				PostId:        v.Field.PostId,
				UUID:          *v.Field.UUID,
				Type:          v.Field.Type,
				Name:          v.Field.Name,
				Key:           v.Field.Key,
				Value:         nil,
				OriginalValue: domain.FieldValue(v.Field.OriginalValue),
			}

			for fi, fv := range posts {
				if fv.Id == v.Id {
					posts[fi].Fields = append(posts[fi].Fields, field)
				}
			}
		}
	}

	return posts
}

func (s *PostStore) getPermalink(post *domain.PostDatum) string {
	permaLink := ""

	postResource := post.Resource
	hiddenCategory := true

	if post.HasResource() {
		resource, ok := s.Config.Resources[*postResource]
		if ok {
			// TODO: This should be in domain.
			permaLink += "/" + strings.ReplaceAll(resource.Slug, "/", "")
			hiddenCategory = resource.HideCategorySlug
		}
	}

	var catSlugs []string

	if post.HasCategory() && !hiddenCategory {
		catSlugs = append(catSlugs, post.Category.Slug)
		parentID := post.Category.ParentId

		for {
			if !post.Category.HasParent() {
				break
			}
			parentCategory, err := s.categoriesModel.GetByID(*parentID)
			if err != nil {
				break
			}
			catSlugs = append(catSlugs, parentCategory.Slug)
			parentID = parentCategory.ParentId
		}
	}

	for i := len(catSlugs) - 1; i >= 0; i-- {
		permaLink += "/" + catSlugs[i]
	}

	permaLink += "/" + post.Slug

	if s.options.SeoEnforceSlash {
		permaLink += "/"
	}

	return permaLink
}

func (s *PostStore) postType(post *domain.PostDatum) domain.PostType {
	if s.options.Homepage == post.Id {
		return domain.PostType{
			PageType: domain.HomeType,
		}
	}

	resource, ok := s.Config.Resources[post.Slug]
	if bool(post.IsArchive) || ok {
		return domain.PostType{
			PageType: domain.ArchiveType,
			Data:     resource,
		}
	}

	// Single with resource
	if post.HasResource() {
		// TODO this should be the resource
		return domain.PostType{
			PageType: domain.SingleType,
			Data:     *post.Resource,
		}
	}

	// Check if the slug is one assigned in categories.
	cat, err := s.categoriesModel.GetBySlug(post.Slug)
	if err == nil {
		return domain.PostType{
			PageType: domain.CategoryType,
			Data:     cat,
		}
	}

	return domain.PostType{
		PageType: domain.PageType,
	}
}
