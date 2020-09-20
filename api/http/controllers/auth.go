package controllers

import (
	"cms/api/config"
	"cms/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AuthController struct {
	authModel models.AuthRepository
	sessionModel models.SessionRepository
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
func newAuth(m models.AuthRepository, s models.SessionRepository) *AuthController {
	return &AuthController{
		authModel: m,
		sessionModel: s,
	}
}

// Login the user
func (c *AuthController) Login(g *gin.Context) {
	var u login
	if err := g.ShouldBindJSON(&u); err != nil {
		Respond(g, 401, "Validation failed", err)
		return
	}

	// Auth user
	lu, err := c.authModel.Authenticate(u.Email, u.Password)
	if err != nil {
		Respond(g, 401, err.Error(), nil)
		return
	}

	// Store session
	sessionToken, err := c.sessionModel.Create(lu.Id, lu.Email)
	cookie := http.Cookie{
		Name: "verbis-session",
		Value: sessionToken,
		HttpOnly: true,
		MaxAge: 43200,
		Path: "/",
		Secure: false,
	}
	http.SetCookie(g.Writer, &cookie)

	Respond(g, 200, "Successfully logged in & session started", lu)
}

// Logout the user
func (c *AuthController) Logout(g *gin.Context) {
	token := g.Request.Header.Get("token")

	userId, err := c.authModel.Logout(token)
	if err != nil {
		Respond(g, 401, "Cannot log the given user out", err.Error())
		return
	}

	if err := c.sessionModel.Delete(userId); err != nil {
		log.Error(err)
	}

	cookie := http.Cookie{
		Name: "verbis-session",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
		Path: "/",
		Secure: false,
	}
	http.SetCookie(g.Writer, &cookie)

	return
}

// Verify email
func (c *AuthController) VerifyEmail(g *gin.Context) {
	token := g.Param("token")

	err := c.authModel.VerifyEmail(token)
	if err != nil {
		fmt.Println(err)
		NoPageFound(g)
		return
	}

	g.Redirect(301, config.Admin.Path)
}

// Reset password
func (c *AuthController) ResetPassword(g *gin.Context) {
	var rp resetPassword
	if err := g.ShouldBindJSON(&rp); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	fmt.Println(rp)

	if err := c.authModel.ResetPassword(rp.Token, rp.Password); err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully reset password", nil)
}

// Verify email
func (c *AuthController) VerifyPasswordToken(g *gin.Context) {
	token := g.Param("token")
	err := c.authModel.VerifyPasswordToken(token)
	if err != nil {
		Respond(g, 404, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully logged in & session started", nil)
}

// Send reset password email & generate token
func (c *AuthController) SendResetPassword(g *gin.Context) {
	var srp sendResetPassword
	if err := g.ShouldBindJSON(&srp); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	err := c.authModel.SendResetPassword(srp.Email)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
	}

	Respond(g, 200, "A fresh verification link has been sent to your email", nil)
}

