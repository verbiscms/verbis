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

		// Middleware
		//operator.Use(middleware.OperatorTokenCheck(d))

		// Session
		api.GET("/session", h.Auth.CheckSession)

		// Fields
		api.GET("/fields", h.Fields.List)

		// Roles
		api.GET("/roles", h.Roles.List)

		// Misc
		api.GET("/layouts", h.Themes.Layouts)
		api.GET("/templates", h.Themes.Templates)
		api.GET("/config", h.Themes.Config)

		// Cache
		api.POST("/cache", middleware.Authorise(domain.Permissions.Settings, domain.UpdateMethod), h.Cache.Clear)

		// Themes
		themes := api.Group("/themes")
		themes.GET("", middleware.Authorise(domain.Permissions.Settings, domain.ViewMethod), h.Themes.List)
		themes.GET("/:name", middleware.Authorise(domain.Permissions.Settings, domain.ViewMethod), h.Themes.Find)
		themes.POST("/:name", middleware.Authorise(domain.Permissions.Settings, domain.UpdateMethod), h.Themes.Activate)

		// Posts
		posts := api.Group("/posts")
		posts.GET("", middleware.Authorise(domain.Permissions.Posts, domain.ViewMethod), h.Posts.List)
		posts.GET("/:id", middleware.Authorise(domain.Permissions.Posts, domain.ViewMethod), h.Posts.Find)
		posts.POST("/", middleware.Authorise(domain.Permissions.Posts, domain.CreateMethod), h.Posts.Create)
		posts.PUT("/:id", middleware.Authorise(domain.Permissions.Posts, domain.UpdateMethod), h.Posts.Update)
		posts.DELETE("/:id", middleware.Authorise(domain.Permissions.Posts, domain.DeleteMethod), h.Posts.Delete)

		// Categories
		categories := api.Group("/categories")
		categories.GET("", middleware.Authorise(domain.Permissions.Categories, domain.ViewMethod), h.Categories.List)
		categories.GET("/:id", middleware.Authorise(domain.Permissions.Categories, domain.ViewMethod), h.Categories.Find)
		categories.POST("/", middleware.Authorise(domain.Permissions.Categories, domain.CreateMethod), h.Categories.Create)
		categories.PUT("/:id", middleware.Authorise(domain.Permissions.Categories, domain.UpdateMethod), h.Categories.Update)
		categories.DELETE("/:id", middleware.Authorise(domain.Permissions.Categories, domain.DeleteMethod), h.Categories.Delete)

		// Media
		media := api.Group("/media")
		media.GET("", middleware.Authorise(domain.Permissions.Media, domain.ViewMethod), h.Media.List)
		media.GET("/:id", middleware.Authorise(domain.Permissions.Media, domain.ViewMethod), h.Media.Find)
		media.POST("", middleware.Authorise(domain.Permissions.Media, domain.CreateMethod), h.Media.Upload)
		media.PUT("/:id", middleware.Authorise(domain.Permissions.Media, domain.UpdateMethod), h.Media.Update)
		media.DELETE("/:id", middleware.Authorise(domain.Permissions.Media, domain.DeleteMethod), h.Media.Delete)

		// Users
		users := api.Group("/users")
		users.GET("", middleware.Authorise(domain.Permissions.Users, domain.ViewMethod), h.Users.List)
		users.GET("/:id", middleware.Authorise(domain.Permissions.Users, domain.ViewMethod), h.Users.Find)
		users.POST("", middleware.Authorise(domain.Permissions.Users, domain.CreateMethod), h.Users.Create)
		users.PUT("/:id", middleware.Authorise(domain.Permissions.Users, domain.UpdateMethod), h.Users.Update)
		users.DELETE("/:id", middleware.Authorise(domain.Permissions.Users, domain.DeleteMethod), h.Users.Delete)
		users.POST("/:id/reset-password", middleware.Authorise(domain.Permissions.Users, domain.UpdateMethod), h.Users.ResetPassword)

		// Settings
		options := api.Group("/options")
		options.GET("", middleware.Authorise(domain.Permissions.Settings, domain.ViewMethod), h.Options.List)
		options.GET("/:name", middleware.Authorise(domain.Permissions.Settings, domain.ViewMethod), h.Options.Find)
		options.POST("", middleware.Authorise(domain.Permissions.Settings, domain.CreateMethod), h.Options.UpdateCreate)

		// Forms
		forms := api.Group("/forms")
		forms.GET("", middleware.Authorise(domain.Permissions.Forms, domain.ViewMethod), h.Forms.List)
		forms.GET("/:id", middleware.Authorise(domain.Permissions.Forms, domain.ViewMethod), h.Forms.Find)
		forms.POST("/:uuid", h.Forms.Send)

		// Redirects
		redirects := api.Group("/redirects")
		redirects.GET("", middleware.Authorise(domain.Permissions.Settings, domain.ViewMethod), h.Redirects.List)
		redirects.GET("/:id", middleware.Authorise(domain.Permissions.Settings, domain.ViewMethod), h.Redirects.Find)
		redirects.POST("", middleware.Authorise(domain.Permissions.Settings, domain.CreateMethod), h.Redirects.Create)
		redirects.PUT("/:id", middleware.Authorise(domain.Permissions.Settings, domain.UpdateMethod), h.Redirects.Update)
		redirects.DELETE("/:id", middleware.Authorise(domain.Permissions.Settings, domain.DeleteMethod), h.Redirects.Delete)

		// Storage
		storage := api.Group("/storage")
		storage.GET("/config", middleware.Authorise(domain.Permissions.Integrations, domain.ViewMethod), h.Storage.Config)
		storage.POST("", middleware.Authorise(domain.Permissions.Integrations, domain.UpdateMethod), h.Storage.Save)
		storage.POST("/migrate", middleware.Authorise(domain.Permissions.Integrations, domain.UpdateMethod), h.Storage.Migrate)
		storage.POST("/bucket", middleware.Authorise(domain.Permissions.Integrations, domain.UpdateMethod), h.Storage.CreateBucket)
		storage.GET("/bucket/:name", middleware.Authorise(domain.Permissions.Integrations, domain.ViewMethod), h.Storage.ListBuckets)
		storage.DELETE("/bucket/:name", middleware.Authorise(domain.Permissions.Integrations, domain.DeleteMethod), h.Storage.DeleteBucket)

		// System
		system := api.Group("/system")
		system.POST("/update", middleware.Authorise(domain.Permissions.Settings, domain.UpdateMethod), h.System.Update)
	}
}
