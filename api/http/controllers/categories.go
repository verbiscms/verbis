package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CategoryHandler defines methods for categories to interact with the server
type CategoryHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// CategoriesController defines the handler for Categories
type CategoriesController struct {
	store *models.Store
	config    config.Configuration
}

// newCategories - Construct
func newCategories(m *models.Store, config config.Configuration) *CategoriesController {
	return &CategoriesController{
		store: m,
		config:    config,
	}
}

// Get all categories
//
// Returns 200 if there are no categories or success.
// Returns 500 if there was an error getting the categories.
// Returns 400 if there was conflict or the request was invalid.
func (c *CategoriesController) Get(g *gin.Context) {
	const op = "CategoryHandler.Get"

	params := http.GetParams(g)
	categories, total, err := c.store.Categories.Get(params)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	pagination := http.GetPagination(params, total)

	Respond(g, 200, "Successfully obtained categories", categories, pagination)
}

// Get By ID
//
// Returns 200 if the category was obtained.
// Returns 500 if there as an error obtaining the category.
// Returns 400 if the ID wasn't passed or failed to convert.
func (c *CategoriesController) GetById(g *gin.Context) {
	const op = "CategoryHandler.GetById"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "Pass a valid number to obtain the category by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	category, err := c.store.Categories.GetById(id)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained category with ID "+string(rune(id)), category)
}

// Create
//
// Returns 200 if the category was created.
// Returns 400 if the the validation failed.
// Returns 500 if there was an error creating the category.
func (c *CategoriesController) Create(g *gin.Context) {
	const op = "CategoryHandler.Create"

	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newCategory, err := c.store.Categories.Create(&category)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created category with ID "+string(rune(newCategory.Id)), newCategory)
}

// Update
//
// Returns 200 if the category was updated.
// Returns 500 if there was an error updating the category.
// Returns 400 if the the validation failed or the category wasn't found.
func (c *CategoriesController) Update(g *gin.Context) {
	const op = "CategoryHandler.Update"

	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500, "A valid ID is required to update the category", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	category.Id = id

	err = c.store.Categories.Update(&category)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated category with ID "+string(rune(category.Id)), category)
}

// Delete
//
// Returns 200 if the category was deleted.
// Returns 500 if there was an error deleting the category.
// Returns 400 if the the category wasn't found or no ID was passed.
func (c *CategoriesController) Delete(g *gin.Context) {
	const op = "CategoryHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to delete a category", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	}

	err = c.store.Categories.Delete(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted category with ID "+string(rune(id)), nil)
}
