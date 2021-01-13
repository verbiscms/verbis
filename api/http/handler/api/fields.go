package api

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

// Fields defines the handler for Fields
type Fields struct {
	store  *models.Store
	config config.Configuration
}

// newFields - Construct
func NewFields(m *models.Store, config config.Configuration) *Fields {
	return &Fields{
		store:  m,
		config: config,
	}
}

// Get - Filter fields and get layouts based on query params.
//
// Returns 200 if login was successful.
// Returns 500 if the layouts failed to be obtained.
func (c *Fields) Get(g *gin.Context) {
	const op = "FieldHandler.Get"

	resource := g.Query("resource")

	userId, err := strconv.Atoi(g.Query("user_id"))
	if err != nil || userId == 0 {
		owner, err := c.store.User.GetOwner()
		if err != nil {
			Respond(g, 500, errors.Message(err), err)
		}
		userId = owner.Id
	}

	categoryId, err := strconv.Atoi(g.Query("category_id"))
	if err != nil {
		categoryId = 0
		//Respond(g, 400, "Field search failed, wrong type passed to category id", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})x
	}

	post := domain.Post{
		Id:                0,
		Slug:              "",
		Title:             "",
		Status:            "",
		Resource:          &resource,
		PageTemplate:      g.Query("page_template"),
		PageLayout:        g.Query("layout"),
		CodeInjectionHead: nil,
		CodeInjectionFoot: nil,
		UserId:            userId,
	}

	// Get the author associated with the post
	author, err := c.store.User.GetById(post.UserId)
	if err != nil {
		author = domain.User{}
	}

	// Get the categories associated with the post
	category, err := c.store.Categories.GetById(categoryId)
	if err != nil {
		category = domain.Category{}
	}

	fields := c.store.Fields.GetLayout(post, author, &category)

	Respond(g, 200, "Successfully obtained fields", fields)
}
