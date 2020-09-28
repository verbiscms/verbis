package controllers

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
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
	controller Controller
	model      models.CategoryRepository
	server     *server.Server
}

// newCategories - Construct
func newCategories(m models.CategoryRepository) *CategoriesController {
	return &CategoriesController{
		model: m,
	}
}

// Get all categories
func (c *CategoriesController) Get(g *gin.Context) {
	const op = "CategoriesController.Get"

	params := http.GetParams(g)
	categories, err := c.model.Get(params)

	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	}

	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained categories", categories)
}

// Get By ID
func (c *CategoriesController) GetById(g *gin.Context) {
	const op = "CategoriesController.GetById"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400,"Pass a valid number to obtain the category by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	category, err := c.model.GetById(id)
	if err != nil {
		Respond(g, 200, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained category with ID " + string(rune(id)), category)
}

// Create
func (c *CategoriesController) Create(g *gin.Context) {
	const op = "CategoriesController.Create"

	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newCategory, err := c.model.Create(&category)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created category with ID " + string(rune(newCategory.Id)), newCategory)
}

// Update
// Returns errors.INVALID if validation failed or the Id is not a string or passed.
func (c *CategoriesController) Update(g *gin.Context) {
	const op = "CategoriesController.Update"

	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500,"A valid ID is required to update the category", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	category.Id = id

	err = c.model.Update(&category)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated category with ID " + string(rune(category.Id)), category)
}

// Delete
// Returns errors.INVALID if the Id is not a string or passed
func (c *CategoriesController) Delete(g *gin.Context) {
	const op = "CategoriesController.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500,"A valid ID is required to delete a category", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	}

	err = c.model.Delete(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted category with ID " + string(rune(id)), nil)
}