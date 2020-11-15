package controllers

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FieldHandler defines methods for fields to interact with the server
type FieldHandler interface {
	Get(g *gin.Context)
}

// FieldController defines the handler for Fields
type FieldController struct {
	fieldsModel 		models.FieldsRepository
	userModel 			models.UserRepository
	categoriesModel		models.CategoryRepository
}

// newFields - Construct
func newFields(f models.FieldsRepository, u models.UserRepository, c models.CategoryRepository) *FieldController {
	return &FieldController{
		fieldsModel: f,
		userModel: u,
		categoriesModel: c,
	}
}

// Filter fields
func (c *FieldController) Get(g *gin.Context) {
	const op = "FieldHandler.Get"

	post := domain.Post{
		Id:             0,
		Slug:           "",
		Title:          "",
		Status:         "",
		Resource:       nil,
		PageTemplate:   "",
		Layout:         "",
		Fields:         nil,
		CodeInjectHead: nil,
		CodeInjectFoot: nil,
		UserId:         0,
		CreatedAt:      nil,
		UpdatedAt:      nil,
		SeoMeta:        domain.PostSeoMeta{},
	}

	// Get the request query
	query := g.Request.URL.Query()

	// Check for page template
	if pt, ok := query["page_template"]; ok {
		post.PageTemplate = pt[0]
	}

	// Check for layout
	if la, ok := query["layout"]; ok {
		post.Layout = la[0]
	}

	// Check for resource
	if re, ok := query["resource"]; ok {
		resource := re[0]
		post.Resource = &resource
	}

	// Check for user ID
	// TODO: clean up here
	if u, ok := query["user_id"]; ok{
		id, err := strconv.Atoi(u[0])
		if err != nil {
			owner, _ := c.userModel.GetOwner()
			post.UserId = owner.Id
		}
		post.UserId = id
	} else {
		owner, _ := c.userModel.GetOwner()
		post.UserId = owner.Id
	}

	// Get the author associated with the post
	author, _ := c.userModel.GetById(post.UserId)

	// Get the categories associated with the post
	categories, err := c.categoriesModel.GetByPost(post.Id)
	if err != nil {
		categories = nil
	}

	fields, err := c.fieldsModel.GetLayout(post, author, categories)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained fields", fields)
}