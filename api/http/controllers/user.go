package controllers

import (
	"cms/api/domain"
	"cms/api/http"
	"cms/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type UserController struct {
	model models.UserRepository
}

type UserHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Construct
func newUser(m models.UserRepository) *UserController {
	return &UserController{
		model: m,
	}
}

// Get All
func (c *UserController) Get(g *gin.Context) {
	params := http.GetParams(g)
	users, err := c.model.Get(params)
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}

	// Remove the token and password
	for k, _ := range users {
		users[k].Password = ""
		users[k].Token = ""
	}

	totalAmount, err := c.model.Total()
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}
	pagination := http.GetPagination(params, totalAmount)

	Respond(g, 200,"Successfully obtained users", users, *pagination)
}

// Get By ID
func (c *UserController) GetById(g *gin.Context) {
	paramId := g.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		Respond(g, 500,  err.Error(), nil)
		return
	}

	user, err := c.model.GetById(id)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	user.Password = ""
	user.Token = ""

	Respond(g, 200, "Successfully obtained user with ID: " + strconv.Itoa(id), user)
}

// Create
func (c *UserController) Create(g *gin.Context) {
	var user domain.User
	if err := g.ShouldBindJSON(&user); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	user, err := c.model.Create(&user)
	if err != nil {Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully created user with ID: " + strconv.Itoa(user.Id), user)
}

// Update
func (c *UserController) Update(g *gin.Context) {
	var u domain.User
	if err := g.ShouldBindJSON(&u); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500,"An ID is required to update the user", nil)
		return
	}
	u.Id = id

	cast := domain.User(u)
	user, err := c.model.Update(&cast)
	if err != nil {
		log.Error(err)
		Respond(g, 500, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully updated user with ID " + strconv.Itoa(u.Id), user)
}

// Delete
func (c *UserController) Delete(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 500, err.Error(), nil)
	}

	err = c.model.Delete(id)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully deleted user with ID " + strconv.Itoa(id), nil)
}
