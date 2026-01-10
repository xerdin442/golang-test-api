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
	}

	events := v1.Group("/events")
	events.Use(middleware.AuthMiddleware())
	{
		events.POST("/", h.CreateEvent)
		events.GET("/", h.GetEvent)
		events.GET("/:id", h.ListEvents)
		events.PUT("/:id", h.UpdateEvent)
		events.DELETE("/:id", h.DeleteEvent)
	}

	attendees := v1.Group("/events/:id/attendees")
	attendees.Use(middleware.AuthMiddleware())
	{
		attendees.GET("/", h.GetEventAttendees)
		attendees.POST("/rsvp", h.ReserveTicket)
		attendees.POST("/revoke", h.RevokeTicket)
	}

	return r
}
