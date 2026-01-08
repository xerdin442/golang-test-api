package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xerdin442/api-practice/internal/api/dto"
	"github.com/xerdin442/api-practice/internal/service"
)

func (h *RouteHandler) CreateEvent(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Context missing user info"})
		return
	}
	userID := uid.(int32)

	var req dto.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	event, err := h.services.Event.CreateEvent(c.Request.Context(), req, userID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidDate):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		}

		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func (h *RouteHandler) UpdateEvent(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Context missing user info"})
		return
	}
	userID := uid.(int32)

	eventID, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	var req dto.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	event, err := h.services.Event.UpdateEvent(c.Request.Context(), req, int32(eventID), userID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidDate):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})
}

func (h *RouteHandler) ListEvents(c *gin.Context) {
	events, err := h.services.Event.ListEvents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func (h *RouteHandler) GetEvent(c *gin.Context) {
	eventID, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	event, err := h.services.Event.GetEvent(c.Request.Context(), int32(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event by ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func (h *RouteHandler) DeleteEvent(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Context missing user info"})
		return
	}
	userID := uid.(int32)

	eventID, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	err := h.services.Event.DeleteEvent(c.Request.Context(), int32(eventID), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
