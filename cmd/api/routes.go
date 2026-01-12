package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xerdin442/api-practice/internal/api/handlers"
	"github.com/xerdin442/api-practice/internal/api/middleware"
)

func (app *application) routes() http.Handler {
	r := gin.Default()
	h := handlers.New(app.services, app.cache)

	v1 := r.Group("api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/login", h.Login)
		auth.POST("/logout", middleware.AuthMiddleware(app.cache), h.Logout)
	}

	users := v1.Group("/users")
	users.Use(middleware.AuthMiddleware(app.cache))
	{
		users.GET("/profile", h.GetProfile)
	}

	events := v1.Group("/events")
	events.Use(middleware.AuthMiddleware(app.cache))
	{
		events.POST("", h.CreateEvent)
		events.GET("/:id", h.GetEvent)
		events.GET("", h.ListEvents)
		events.PUT("/:id", h.UpdateEvent)
		events.DELETE("/:id", h.DeleteEvent)
	}

	attendees := v1.Group("/events/:id/attendees")
	attendees.Use(middleware.AuthMiddleware(app.cache))
	{
		attendees.GET("/", h.GetEventAttendees)
		attendees.POST("/rsvp", h.ReserveTicket)
		attendees.POST("/revoke", h.RevokeTicket)
	}

	return r
}
