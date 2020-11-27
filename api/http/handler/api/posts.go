package api

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// PostHandler defines methods for Posts to interact with the server
type PostHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Posts defines the handler for Posts
type Posts struct {
	store  *models.Store
	config config.Configuration
}

// newPosts - Construct
func NewPosts(m *models.Store, config config.Configuration) *Posts {
	return &Posts{
		store:  m,
		config: config,
	}
}

// Get all posts, obtain resource param to pass to the get
// function.
//
// Returns 200 if there are no posts or success.
// Returns 400 if there was conflict or the request was invalid.
// Returns 500 if there was an error getting or formatting the posts.
func (c *Posts) Get(g *gin.Context) {
	const op = "PostHandler.Get"

	params := http.NewParams(g).Get()
	posts, total, err := c.store.Posts.Get(params, g.Query("resource"))
	if errors.Code(err) == errors.NOTFOUND {
		handler.Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	postData, err := c.store.Posts.FormatMultiple(posts)
	if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	pagination := http.NewPagination().Get(params, total)

	handler.Respond(g, 200, "Successfully obtained posts", postData, pagination)
}

// Get By ID
//
// Returns 200 if the posts were obtained.
// Returns 400 if the ID wasn't passed or failed to convert.
// Returns 500 if there as an error obtaining or formatting the post.
func (c *Posts) GetById(g *gin.Context) {
	const op = "PostHandler.GetById"

	paramId := g.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		handler.Respond(g, 400, "Pass a valid number to obtain the post by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	post, err := c.store.Posts.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		handler.Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	formatPost, err := c.store.Posts.Format(post)
	if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully obtained post with ID: "+paramId, formatPost)
}

// Create
//
// Returns 200 if the post was created.
// Returns 500 if there was an error creating or formatting the post.
// Returns 400 if the the validation failed or there was a conflict with the post.
func (c *Posts) Create(g *gin.Context) {
	const op = "PostHandler.Create"

	var post domain.PostCreate
	if err := g.ShouldBindJSON(&post); err != nil {
		handler.Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newPost, err := c.store.Posts.Create(&post)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	formatPost, err := c.store.Posts.Format(newPost)
	if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully created post with ID: "+strconv.Itoa(newPost.Id), formatPost)
}

// Update
//
// Returns 200 if the post was updated.
// Returns 500 if there was an error updating or formatting the post.
// Returns 400 if the the validation failed, there was a conflict, or the post wasn't found.
func (c *Posts) Update(g *gin.Context) {
	const op = "PostHandler.Update"

	var post domain.PostCreate
	if err := g.ShouldBindJSON(&post); err != nil {
		handler.Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		handler.Respond(g, 400, "A valid ID is required to update the post", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	post.Id = id

	updatedPost, err := c.store.Posts.Update(&post)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	formatPost, err := c.store.Posts.Format(updatedPost)
	if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully updated post with ID: "+strconv.Itoa(updatedPost.Id), formatPost)
}

// Delete
//
// Returns 200 if the post was deleted.
// Returns 500 if there was an error deleting the post.
// Returns 400 if the the post wasn't found or no ID was passed.
func (c *Posts) Delete(g *gin.Context) {
	const op = "PostHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		handler.Respond(g, 400, "A valid ID is required to delete a post", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = c.store.Posts.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully deleted post with ID: "+strconv.Itoa(id), nil)
}
