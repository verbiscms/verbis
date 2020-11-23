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

// UserHandler defines methods for Users to interact with the server
type UserHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	GetRoles(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
	ResetPassword(g *gin.Context)
}

// UserController defines the handler for Users
type UserController struct {
	store  *models.Store
	config config.Configuration
}

// newUser - Construct
func newUser(m *models.Store, config config.Configuration) *UserController {
	return &UserController{
		store:  m,
		config: config,
	}
}

// Get all users
//
// Returns 200 if the users were obtained successfully.
// Returns 500 if there was an error getting the users.
// Returns 400 if there was conflict or the request was invalid.
func (c *UserController) Get(g *gin.Context) {
	const op = "UserHandler.Get"

	params := http.NewParams(g).Get()
	users, total, err := c.store.User.Get(params)
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
	users.HideCredentials()

	pagination := http.NewPagination().Get(params, total)

	Respond(g, 200, "Successfully obtained users", users, pagination)
}

// Get By ID
//
// Returns 200 if the user was obtained.
// Returns 500 if there as an error obtaining the user.
// Returns 400 if the ID wasn't passed or failed to convert.
func (c *UserController) GetById(g *gin.Context) {
	const op = "UserHandler.GetById"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "Pass a valid number to obtain the user by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := c.store.User.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	user.HideCredentials()

	Respond(g, 200, "Successfully obtained user with ID: "+strconv.Itoa(id), user)
}

// Get Roles
//
// Returns 200 if the user roles were obtained.
// Returns 500 if there as an error obtaining the user roles.
func (c *UserController) GetRoles(g *gin.Context) {
	const op = "UserHandler.GetRoles"

	roles, err := c.store.User.GetRoles()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained user roles", roles)
}

// Create
//
// Returns 200 if the user was created.
// Returns 500 if there was an error creating the user.
// Returns 400 if the the validation failed or a user already exists.
func (c *UserController) Create(g *gin.Context) {
	const op = "UserHandler.Create"

	var u domain.UserCreate
	if err := g.ShouldBindJSON(&u); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := c.store.User.Create(&u)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created user with ID: "+strconv.Itoa(user.Id), user)
}

// Update
//
// Returns 200 if the user was updated.
// Returns 500 if there was an error updating the user.
// Returns 400 if the the validation failed or the user wasn't found.
func (c *UserController) Update(g *gin.Context) {
	const op = "UserHandler.Update"

	var u domain.User
	if err := g.ShouldBindJSON(&u); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to update the user", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	u.Id = id

	updatedUser, err := c.store.User.Update(&u)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated user with ID: "+strconv.Itoa(u.Id), updatedUser)
}

// Delete
//
// Returns 200 if the user was deleted.
// Returns 500 if there was an error deleting the user.
// Returns 400 if the the user wasn't found or no ID was passed.
func (c *UserController) Delete(g *gin.Context) {
	const op = "UserHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to delete a user", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = c.store.User.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted user with ID: "+strconv.Itoa(id), nil)
}

// ResetPassword
//
// Returns 200 if the reset password was successful.
// Returns 500 if there was an error resetting the user failed.
// Returns 400 if the the user wasn't found, no ID was passed or validation failed.
func (c *UserController) ResetPassword(g *gin.Context) {
	const op = "UserHandler.ResetPassword"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to update a user's password", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
	}

	user, err := c.store.User.GetById(id)
	if err != nil {
		Respond(g, 200, "No user has been found with the ID:" + strconv.Itoa(id), err)
		return
	}

	var reset domain.UserPasswordReset
	reset.DBPassword = user.Password
	if err := g.ShouldBindJSON(&reset); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = c.store.User.ResetPassword(id, reset)
	if errors.Code(err) == errors.INVALID {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated password for the user with ID: "+strconv.Itoa(id), nil)
}
