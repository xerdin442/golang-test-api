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
		auth.POST("/signup")
		auth.POST("/login")
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
