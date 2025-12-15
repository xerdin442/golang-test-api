package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xerdin442/api-practice/internal/api/handlers"
)

func (app *application) routes() http.Handler {
	r := gin.Default()
	h := handlers.New(app.services)

	v1 := r.Group("api/v1")
	{
		v1.POST("/events", h.CreateEvent)
		v1.GET("/events")
		v1.GET("/events/:id")
		v1.PUT("/events/:id")
		v1.DELETE("/events/:id")
	}

	return r
}
