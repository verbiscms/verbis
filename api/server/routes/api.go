// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/http/handler"
	"github.com/verbiscms/verbis/api/http/middleware"
	"github.com/verbiscms/verbis/api/http/sockets"
	"github.com/verbiscms/verbis/api/server"
)

// apiRoutes API facing routes.
func apiRoutes(d *deps.Deps, s *server.Server) {
	api := s.Group(app.HTTPAPIRoute)
	{
		// API Middleware
		api.Use(middleware.CORS())
		api.Use(middleware.EmptyBody())

		if !d.Installed {
			h := handler.NewInstall(d)
			// Preflight
			api.POST("/install/validate/:step", h.System.ValidateInstall)
			// Install
			api.POST("/install", h.System.Install)
			return
		}

		h := handler.NewAPI(d)

		api.Use(middleware.TokenCheck(d.Store.User))

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

		// Operator
		operator := api.Group("")
		// Middleware
		//operator.Use(middleware.OperatorTokenCheck(d))

		// System
		operator.POST("/update", h.System.Update)

		// Session
		operator.GET("/session", h.Auth.CheckSession)

		// Fields
		operator.GET("/fields", h.Fields.List)

		// Roles
		operator.GET("/roles", h.Roles.List)

		// Misc
		operator.GET("/layouts", h.Themes.Layouts)
		operator.GET("/templates", h.Themes.Templates)
		operator.GET("/config", h.Themes.Config)

		// Cache
		operator.POST("/cache", h.Cache.Clear)

		// Themes
		themes := api.Group("/themes")
		{
			themes.GET("", h.Themes.List)
			themes.GET("/:name", h.Themes.Find)
			themes.POST("/:name", h.Themes.Activate)
		}

		// Posts
		posts := api.Group("/posts")
		{
			posts.GET("", h.Posts.List)
			posts.GET("/:id", h.Posts.Find)
			posts.POST("/", h.Posts.Create)
			posts.PUT("/:id", h.Posts.Update)
			posts.DELETE("/:id", h.Posts.Delete)
		}

		// Categories
		categories := api.Group("/categories")
		{
			categories.GET("", h.Categories.List)
			categories.GET("/:id", h.Categories.Find)
			categories.POST("/", h.Categories.Create)
			categories.PUT("/:id", h.Categories.Update)
			categories.DELETE("/:id", h.Categories.Delete)
		}

		// Media
		media := api.Group("/media")
		{
			media.GET("", h.Media.List)
			media.GET("/:id", h.Media.Find)
			media.POST("", h.Media.Upload)
			media.PUT("/:id", h.Media.Update)
			media.DELETE("/:id", h.Media.Delete)
		}

		// Users
		users := api.Group("/users")
		{
			users.GET("", h.Users.List)
			users.GET("/:id", h.Users.Find)
			users.PUT("/:id", h.Users.Update)
			users.POST("", h.Users.Create)
			users.DELETE("/:id", h.Users.Delete)
			users.POST("/:id/reset-password", h.Users.ResetPassword)
		}

		// Settings
		settings := api.Group("/options")
		{
			settings.GET("", middleware.Authorise("settings", domain.ListMethod), h.Options.List)
			settings.GET("/:name", h.Options.Find)
			settings.POST("", h.Options.UpdateCreate)
		}

		// Forms
		forms := api.Group("/forms")
		{
			forms.GET("", h.Forms.List)
			forms.GET("/:id", h.Forms.Find)
			forms.POST("/:uuid", h.Forms.Send)
		}

		// Redirects
		redirects := api.Group("/redirects")
		{
			redirects.GET("", h.Redirects.List)
			redirects.GET("/:id", h.Redirects.Find)
			redirects.POST("", h.Redirects.Create)
			redirects.PUT("/:id", h.Redirects.Update)
			redirects.DELETE("/:id", h.Redirects.Delete)
		}

		// Storage
		storage := api.Group("/storage")
		{
			storage.POST("", h.Storage.Save)
			storage.GET("/config", h.Storage.Config)
			storage.POST("/migrate", h.Storage.Migrate)
			storage.POST("/bucket", h.Storage.CreateBucket)
			storage.GET("/bucket/:name", h.Storage.ListBuckets)
			storage.DELETE("/bucket/:name", h.Storage.DeleteBucket)
		}
	}
}
