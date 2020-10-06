package controllers

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// PostHandler defines methods for Posts to interact with the server
type PostHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// PostsController defines the handler for Posts
type PostsController struct {
	postModel   	models.PostsRepository
	fieldsModel 	models.FieldsRepository
	userModel 		models.UserRepository
	categoriesModel	models.CategoryRepository
}

// newPosts - Construct
func newPosts(m models.PostsRepository, f models.FieldsRepository, u models.UserRepository, c models.CategoryRepository) *PostsController {
	return &PostsController{
		postModel: m,
		fieldsModel: f,
		userModel: u,
		categoriesModel: c,
	}
}

// Get all posts
func (c *PostsController) Get(g *gin.Context) {
	const op = "PostHandler.Get"

	params := http.GetParams(g)
	posts, err := c.postModel.Get(params)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	}
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	// Loop over all posts and obtain data
	var returnData []domain.PostData
	for _, post := range posts {
		formatted, err := c.Format(g, post)
		if err != nil {
			Respond(g, 500, errors.Message(err), err)
			return
		} else {
			returnData = append(returnData, formatted)
		}
	}

	// Get the total number of posts for response
	totalAmount, err := c.postModel.Total()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	pagination := http.GetPagination(params, totalAmount)

	Respond(g, 200, "Successfully obtained posts", returnData, *pagination)
}

// Get By ID
// Returns errors.INVALID if the Id is not a string or passed.
func (c *PostsController) GetById(g *gin.Context) {
	const op = "PostHandler.GetById"

	paramId := g.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		Respond(g, 400,  "Pass a valid number to obtain the post by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	post, err := c.postModel.GetById(id)
	if err != nil {
		Respond(g, 200, errors.Message(err), err)
		return
	}

	formatPost, err := c.Format(g, post)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained post with the ID: " + paramId, formatPost)
}

// Create
// Returns errors.INVALID if validation failed.
func (c *PostsController) Create(g *gin.Context) {
	const op = "PostHandler.Create"

	var post domain.PostCreate
	if err := g.ShouldBindJSON(&post); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newPost, err := c.postModel.Create(&post)
	if errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	formatPost, err := c.Format(g, newPost)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 201, "Successfully created post with the ID: " + strconv.Itoa(newPost.Id), formatPost)
}

// Update
// Returns errors.INVALID if validation failed or the Id is not a string or passed.
func (c *PostsController) Update(g *gin.Context) {
	const op = "PostHandler.Update"

	var post domain.PostCreate
	if err := g.ShouldBindJSON(&post); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400,"A valid ID is required to update the post", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	post.Id = id

	updatedPost, err := c.postModel.Update(&post)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	formatPost, err := c.Format(g, updatedPost)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated post with the ID: " + strconv.Itoa(updatedPost.Id), formatPost)
}

// Delete
// Returns errors.INVALID if the Id is not a string or passed
func (c *PostsController) Delete(g *gin.Context) {
	const op = "PostHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400,"A valid ID is required to delete a post", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	}

	err = c.postModel.Delete(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted post with the ID: " + strconv.Itoa(id), nil)
}

// Format
func (c *PostsController) Format(g *gin.Context, post domain.Post) (domain.PostData, error) {

	// Get the author associated with the post
	author, err := c.userModel.GetById(post.UserId)
	if err != nil {
		return domain.PostData{}, err
	}

	// Get the categories associated with the post
	categories, err := c.categoriesModel.GetByPost(post.Id)
	if err != nil {
		return domain.PostData{}, err
	}

	// Get the layout associated with the post
	layout, err := c.fieldsModel.GetLayout(post, author, categories)
	if err != nil {
		return domain.PostData{}, err
	}

	return domain.PostData{
		Post:       post,
		Layout:     layout,
		Author:     domain.PostAuthor(author),
		Categories: c.processCategories(categories),
	}, nil
}

// Process Categories
func (c *PostsController) processCategories(dc []domain.Category) []domain.PostCategory {
	var postCategories []domain.PostCategory

	if len(dc) == 0 {
		return make([]domain.PostCategory, 0)
	} else {
		for _, v := range dc {
			postCategories = append(postCategories, domain.PostCategory{
				Id:           v.Id,
				Slug:         v.Slug,
				Name:         v.Name,
				Description:  v.Description,
				Hidden:       v.Hidden,
				ParentId:     v.ParentId,
				PageTemplate: v.PageTemplate,
				UpdatedAt:    time.Time{},
				CreatedAt:    time.Time{},
			})
		}
	}

	return postCategories
}