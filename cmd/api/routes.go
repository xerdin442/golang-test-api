package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xerdin442/api-practice/internal/api/handlers"
	"github.com/xerdin442/api-practice/internal/api/middleware"
)

func (app *application) routes() http.Handler {
	r := gin.Default()
	h := handlers.New(app.services)

	v1 := r.Group("api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/login", h.Login)
	}

	users := v1.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.POST("/logout", h.Logout)
		users.GET("/profile", h.GetProfile)
		users.PUT("/profile", h.UpdateProfile)
	}

	events := v1.Group("/events")
	events.Use(middleware.AuthMiddleware())
	{
		events.POST("/", h.CreateEvent)
		events.GET("/")
		events.GET("/:id")
		events.PUT("/:id")
		events.DELETE("/:id")
	}

	return r
}
