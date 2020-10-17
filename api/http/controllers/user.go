package controllers

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserHandler defines methods for Users to interact with the server
type UserHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	GetRoles(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// UserController defines the handler for Users
type UserController struct {
	model models.UserRepository
}

// newUser - Construct
func newUser(m models.UserRepository) *UserController {
	return &UserController{
		model: m,
	}
}

// Get all users
func (c *UserController) Get(g *gin.Context) {
	const op = "UserHandler.Get"

	params := http.GetParams(g)
	users, total, err := c.model.Get(params)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	// Remove the token and password
	for k, _ := range users {
		users[k].Password = ""
		users[k].Token = ""
	}

	pagination := http.GetPagination(params, total)

	Respond(g, 200,"Successfully obtained users", users, pagination)
}

// Get By ID
// Returns errors.INVALID if the Id is not a string or passed.
func (c *UserController) GetById(g *gin.Context) {
	const op = "UserHandler.GetById"

	paramId := g.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		Respond(g, 400,  "Pass a valid number to obtain the user by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := c.model.GetById(id)
	if err != nil {
		Respond(g, 200, errors.Message(err), err)
		return
	}

	user.Password = ""
	user.Token = ""

	Respond(g, 200, "Successfully obtained user with ID: " + strconv.Itoa(id), user)
}

// Get Roles
func (c *UserController) GetRoles(g *gin.Context) {
	const op = "UserHandler.GetRoles"

	roles, err := c.model.GetRoles()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained user roles", roles)
}

// Create
// Returns errors.INVALID if validation failed.
func (c *UserController) Create(g *gin.Context) {
	const op = "UserHandler.Create"

	var user domain.User
	if err := g.ShouldBindJSON(&user); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := c.model.Create(&user)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created user with ID: " + strconv.Itoa(user.Id), user)
}

// Update
// Returns errors.INVALID if validation failed or the Id is not a string or passed.
func (c *UserController) Update(g *gin.Context) {
	const op = "UserHandler.Update"

	var u domain.User
	if err := g.ShouldBindJSON(&u); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400,"A valid ID is required to update the user", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	u.Id = id

	updatedUser, err := c.model.Update(&u)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated user with ID " + strconv.Itoa(u.Id), updatedUser)
}

// Delete
// Returns errors.INVALID if the Id is not a string or passed
func (c *UserController) Delete(g *gin.Context) {
	const op = "UserHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400,"A valid ID is required to delete a post", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	}

	err = c.model.Delete(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted user with ID " + strconv.Itoa(id), nil)
}
