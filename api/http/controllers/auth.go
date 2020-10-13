package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authModel models.AuthRepository
	sessionModel models.SessionRepository
	userModel models.UserRepository
}

type AuthHandler interface {
	Login(g *gin.Context)
	Logout(g *gin.Context)
	ResetPassword(g *gin.Context)
	VerifyEmail(g *gin.Context)
	VerifyPasswordToken(g *gin.Context)
	SendResetPassword(g *gin.Context)
}

type login struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type sendResetPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type showPasswordReset struct {
	Token string `db:"token" json:"token" binding:"required"`
}

type resetPassword struct {
	Password string	`db:"password" json:"password" binding:"required,min=8,max=60,alphanum"`
	Token string `db:"token" json:"token" binding:"required"`
}

// Construct
func newAuth(m models.AuthRepository, s models.SessionRepository, u models.UserRepository) *AuthController {
	return &AuthController{
		authModel: m,
		sessionModel: s,
		userModel: u,
	}
}

// Login the user
// Returns errors.INVALID if validation failed
func (c *AuthController) Login(g *gin.Context) {
	const op = "AuthHandler.Login"

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

	// Store session
	sessionToken, err := c.sessionModel.Create(user.Id, lu.Email)

	Respond(g, 200, "Successfully logged in & session started", gin.H{
		"user": user,
		"session": sessionToken,
	})
}

// Logout the user
func (c *AuthController) Logout(g *gin.Context) {
	const op = "AuthHandler.Logout"

	token := g.Request.Header.Get("token")
	userId, err := c.authModel.Logout(token)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	}
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	if err := c.sessionModel.Delete(userId); err != nil {
		Respond(g, 500, errors.Message(err), err)
	}

	//cookie := http.Cookie{
	//	Name: "verbis-session",
	//	Value: "",
	//	HttpOnly: true,
	//	MaxAge: -1,
	//	Path: "/",
	//	Secure: false,
	//}
	//http.SetCookie(g.Writer, &cookie)

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

	g.Redirect(301, config.Admin.Path)
}

// Reset password
func (c *AuthController) ResetPassword(g *gin.Context) {
	const op = "AuthHandler.ResetPassword"

	var rp resetPassword
	if err := g.ShouldBindJSON(&rp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if err := c.authModel.ResetPassword(rp.Token, rp.Password); err != nil {
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

	var srp sendResetPassword
	if err := g.ShouldBindJSON(&srp); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err := c.authModel.SendResetPassword(srp.Email)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	}
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "A fresh verification link has been sent to your email", nil)
}

