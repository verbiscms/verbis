package controllers

import (
	"cms/api/domain"
	"cms/api/http"
	"cms/api/models"
	"cms/api/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoriesController struct {
	controller Controller
	model      models.CategoryRepository
	server     *server.Server
}

type CategoryHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Construct
func newCategories(m models.CategoryRepository) *CategoriesController {
	return &CategoriesController{
		model: m,
	}
}

// Get All
func (c *CategoriesController) Get(g *gin.Context) {
	params := http.GetParams(g)
	categories, err := c.model.Get(params)
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}

	successMsg := "Successfully obtained categories"
	if len(categories) == 0 {
		successMsg = "No categories available"
	}
	Respond(g, 200, successMsg, categories)
}

// Get By ID
func (c *CategoriesController) GetById(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}

	category, err := c.model.GetById(id)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully obtained category with ID " + string(id), category)
}

// Insert
func (c *CategoriesController) Create(g *gin.Context) {
	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	id, err := c.model.Create(&category)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	category.Id = id
	Respond(g, 200, "Successfully created category with ID " + string(id), category)
}

// Update
func (c *CategoriesController) Update(g *gin.Context) {
	var category domain.Category
	if err := g.ShouldBindJSON(&category); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	err := c.model.Update(&category)
	if err != nil {
		Respond(g, 500, err.Error(), nil)
	}

	Respond(g, 200, "Successfully updated category with ID " + string(category.Id), category)
}

// Delete
func (c *CategoriesController) Delete(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500, err.Error(), nil)
	}

	err = c.model.Delete(id)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully deleted category with ID " + string(id), nil)
}