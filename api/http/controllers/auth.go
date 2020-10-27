package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authModel models.AuthRepository
	userModel models.UserRepository
	config 	  config.Configuration
}

type AuthHandler interface {
	Login(g *gin.Context)
	Logout(g *gin.Context)
	ResetPassword(g *gin.Context)
	VerifyEmail(g *gin.Context)
	VerifyPasswordToken(g *gin.Context)
	SendResetPassword(g *gin.Context)
}

// Construct
func newAuth(m models.AuthRepository, u models.UserRepository, config config.Configuration) *AuthController {
	return &AuthController{
		authModel: m,
		userModel: u,
		config:    config,
	}
}

// Login the user
// Returns errors.INVALID if validation failed
func (c *AuthController) Login(g *gin.Context) {
	const op = "AuthHandler.Login"

	type login struct {
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var u login
	if err := g.ShouldBindJSON(&u); err != nil {
		Respond(g, 400,  "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	// Auth user
	lu, err := c.authModel.Authenticate(u.Email, u.Password)
	if err != nil {
		Respond(g, 401, errors.Message(err), err)
		return
	}

	// Get the user & role
	user, err := c.userModel.GetById(lu.Id)
	if err != nil {
		Respond(g, 401, errors.Message(err), err)
		return
	}

	// Remove the password
	user.Password = ""

	// Set the verbis cookie
	g.SetCookie("verbis-session", user.Token, 172800, "/", "", false, true)

	Respond(g, 200, "Successfully logged in & session started", user)
}

// Logout the user
func (c *AuthController) Logout(g *gin.Context) {
	const op = "AuthHandler.Logout"

	token := g.Request.Header.Get("token")
	_, err := c.authModel.Logout(token)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	g.SetCookie("verbis-session", "", -1, "/", "", false, true)

	return
}

// Verify email
func (c *AuthController) VerifyEmail(g *gin.Context) {
	const op = "AuthHandler.VerifyEmail"

	token := g.Param("token")
	err := c.authModel.VerifyEmail(token)
	if err != nil {
		NoPageFound(g)
		return
	}

	g.Redirect(301, c.config.Admin.Path)
}

// Reset password
func (c *AuthController) ResetPassword(g *gin.Context) {
	const op = "AuthHandler.ResetPassword"

	type resetPassword struct {
		NewPassword			string			`json:"new_password" binding:"required,min=8,max=60"`
		ConfirmPassword		string			`json:"confirm_password" binding:"eqfield=NewPassword,required"`
		Token 				string 			`db:"token" json:"token" binding:"required"`
	}

	var rp resetPassword
	if err := g.ShouldBindJSON(&rp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if err := c.authModel.ResetPassword(rp.Token, rp.NewPassword); err != nil {
		Respond(g, 400, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully reset password", nil)
}

// VerifyPasswordToken
// Returns errors.INVALID if validation failed.
func (c *AuthController) VerifyPasswordToken(g *gin.Context) {
	const op = "AuthHandler.VerifyPasswordToken"

	token := g.Param("token")
	err := c.authModel.VerifyPasswordToken(token)
	if err != nil {
		Respond(g, 404, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully logged in & session started", nil)
}

// SendResetPassword reset password email & generate token
func (c *AuthController) SendResetPassword(g *gin.Context) {
	const op = "AuthHandler.SendResetPassword"

	type sendResetPassword struct {
		Email string `json:"email" binding:"required,email"`
	}

	var srp sendResetPassword
	if err := g.ShouldBindJSON(&srp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err := c.authModel.SendResetPassword(srp.Email)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	}
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "A fresh verification link has been sent to your email", nil)
}

