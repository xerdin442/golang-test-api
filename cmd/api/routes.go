package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xerdin442/api-practice/internal/api/handlers"
	"github.com/xerdin442/api-practice/internal/api/middleware"
)

func (app *application) routes() http.Handler {
	r := gin.New()
	r.Use(middleware.CustomRequestLogger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RateLimiter())

	v1 := r.Group("api/v1")
	h := handlers.New(app.services, app.cache)

	auth := v1.Group("/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/login", h.Login)
		auth.POST("/logout", middleware.JwtGuard(app.cache), h.Logout)
	}

	users := v1.Group("/users")
	users.Use(middleware.JwtGuard(app.cache))
	{
		users.GET("/profile", h.GetProfile)
	}

	events := v1.Group("/events")
	events.Use(middleware.JwtGuard(app.cache))
	{
		events.POST("", h.CreateEvent)
		events.GET("/:id", h.GetEvent)
		events.GET("", h.ListEvents)
		events.PUT("/:id", h.UpdateEvent)
		events.DELETE("/:id", h.DeleteEvent)
	}

	attendees := v1.Group("/events/:id/attendees")
	attendees.Use(middleware.JwtGuard(app.cache))
	{
		attendees.GET("/", h.GetEventAttendees)
		attendees.POST("/rsvp", h.ReserveTicket)
		attendees.POST("/revoke", h.RevokeTicket)
	}

	return r
}
