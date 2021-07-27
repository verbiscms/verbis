// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/http/handler"
	"github.com/verbiscms/verbis/api/http/middleware"
	"github.com/verbiscms/verbis/api/http/sockets"
	"github.com/verbiscms/verbis/api/server"
)

// apiRoutes
//
// API facing routes.
func apiRoutes(d *deps.Deps, s *server.Server) {
	h := handler.NewAPI(d)

	// API Routes
	api := s.Group(app.HTTPAPIRoute)
	{
		// API Middleware
		api.Use(middleware.CORS())
		api.Use(middleware.EmptyBody())

		// Sockets
		api.GET("/ws", gin.WrapH(sockets.Handler()))

		// Site
		api.GET("/site", h.Site.Global)

		// Auth
		api.POST("/login", h.Auth.Login)
		api.POST("/logout", h.Auth.Logout)
		api.POST("/password/reset", h.Auth.ResetPassword)
		api.POST("/password/email", h.Auth.SendResetPassword)
		api.GET("/password/verify/:token", h.Auth.VerifyPasswordToken)
		//	api.GET("/email/verify/:token", h.Auth.VerifyEmail)

		// Forms
		forms := api.Group("/forms")
		forms.POST("/:uuid", h.Forms.Send)

		// Operator
		operator := api.Group("")
		// Middleware
		operator.Use(middleware.OperatorTokenCheck(d))

		// Update
		operator.POST("/update", h.System.Update)

		operator.GET("/session", h.Auth.CheckSession)

		// Themes
		operator.GET("/themes", h.Themes.List)
		operator.GET("/themes/:name", h.Themes.Find)

		// Themes
		operator.GET("/layouts", h.Themes.Layouts)
		operator.GET("/templates", h.Themes.Templates)
		operator.GET("/config", h.Themes.Config)
		operator.POST("/theme", h.Themes.Update)

		// Posts
		operator.GET("/posts", h.Posts.List)
		operator.GET("/posts/:id", h.Posts.Find)
		operator.POST("/posts", h.Posts.Create)
		operator.PUT("/posts/:id", h.Posts.Update)
		operator.DELETE("/posts/:id", h.Posts.Delete)

		// Categories
		operator.GET("/categories", h.Categories.List)
		operator.GET("/categories/:id", h.Categories.Find)
		operator.POST("/categories", h.Categories.Create)
		operator.PUT("/categories/:id", h.Categories.Update)
		operator.DELETE("/categories/:id", h.Categories.Delete)

		// Media
		operator.GET("/media", h.Media.List)
		operator.GET("/media/:id", h.Media.Find)
		operator.POST("/media", h.Media.Upload)
		operator.PUT("/media/:id", h.Media.Update)
		operator.DELETE("/media/:id", h.Media.Delete)

		// Users
		operator.GET("/users", h.Users.List)
		operator.GET("/users/:id", h.Users.Find)
		operator.PUT("/users/:id", h.Users.Update)
		operator.POST("/users/:id/reset-password", h.Users.ResetPassword)

		// Fields
		operator.GET("/fields", h.Fields.List)

		// Options
		operator.GET("/options", h.Options.List)
		operator.GET("/options/:name", h.Options.Find)
		operator.POST("/options", h.Options.UpdateCreate)

		// Roles
		operator.GET("/roles", h.Roles.List)

		// Redirects
		operator.GET("/redirects", h.Redirects.List)
		operator.GET("/redirects/:id", h.Redirects.Find)
		operator.POST("/redirects", h.Redirects.Create)
		operator.PUT("/redirects/:id", h.Redirects.Update)
		operator.DELETE("/redirects/:id", h.Redirects.Delete)

		// Cache
		operator.POST("/cache", h.Cache.Clear)

		// Forms
		operator.GET("/forms", h.Forms.List)
		operator.GET("/forms/:id", h.Forms.Find)

		// Storage
		operator.POST("/storage", h.Storage.Save)
		operator.GET("/storage/config", h.Storage.Config)
		operator.POST("/storage/migrate", h.Storage.Migrate)
		operator.POST("/storage/bucket", h.Storage.CreateBucket)
		operator.GET("/storage/bucket/:name", h.Storage.ListBuckets)
		operator.DELETE("/storage/bucket/:name", h.Storage.DeleteBucket)

		// Administrator
		admin := api.Group("")
		admin.Use(middleware.AdminTokenCheck(d))

		// Users
		admin.POST("/users", h.Users.Create)
		admin.DELETE("/users/:id", h.Users.Delete)
	}
}
