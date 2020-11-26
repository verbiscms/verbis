package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
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
	store  *models.Store
	config config.Configuration
}

// newFields - Construct
func newFields(m *models.Store, config config.Configuration) *FieldController {
	return &FieldController{
		store:  m,
		config: config,
	}
}

// Get - Filter fields and get layouts based on query params.
//
// Returns 200 if login was successful.
// Returns 500 if the layouts failed to be obtained.
func (c *FieldController) Get(g *gin.Context) {
	const op = "FieldHandler.Get"

	resource := g.Query("resource")
	user, err := strconv.Atoi(g.Query("user_id"))
	if err != nil {
		Respond(g, 400, "Field search failed, wrong type passed to user id", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if user == 0 {
		owner, err := c.store.User.GetOwner()
		if err != nil {
			Respond(g, 500, errors.Message(err), err)
		}
		user = owner.Id
	}

	post := domain.Post{
		Id:             0,
		Slug:           "",
		Title:          "",
		Status:         "",
		Resource:       &resource,
		PageTemplate:   g.Query("page_template"),
		Layout:         g.Query("layout"),
		Fields:         nil,
		CodeInjectHead: nil,
		CodeInjectFoot: nil,
		UserId:         user,
	}

	// Get the author associated with the post
	author, err := c.store.User.GetById(post.UserId)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	// Get the categories associated with the post
	categories, err := c.store.Categories.GetByPost(post.Id)
	if err != nil {
		categories = nil
	}

	fields, err := c.store.Fields.GetLayout(post, author, categories)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained fields", fields)
}
