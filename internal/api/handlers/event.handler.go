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
	userID := c.MustGet("userID").(int32)

	var req dto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	userID := c.MustGet("userID").(int32)

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var req dto.UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := h.services.Event.UpdateEvent(c.Request.Context(), req, int32(eventID), userID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidDate), errors.Is(err, service.ErrEventNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case errors.Is(err, service.ErrOwnerRestrictedAction):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

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
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := h.services.Event.GetEvent(c.Request.Context(), int32(eventID))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEventNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event by ID"})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func (h *RouteHandler) DeleteEvent(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	if err := h.services.Event.DeleteEvent(c.Request.Context(), int32(eventID), userID); err != nil {
		switch {
		case errors.Is(err, service.ErrEventNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case errors.Is(err, service.ErrOwnerRestrictedAction):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func (h *RouteHandler) ReserveTicket(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	if err := h.services.Event.ReserveTicket(c.Request.Context(), userID, int32(eventID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reserve event ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event ticket reserved successfully"})
}

func (h *RouteHandler) RevokeTicket(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	if err := h.services.Event.RevokeTicket(c.Request.Context(), userID, int32(eventID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke event ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event ticket revoked successfully"})
}

func (h *RouteHandler) GetEventAttendees(c *gin.Context) {
	userID := c.MustGet("userID").(int32)

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	attendees, err := h.services.Event.GetEventAttendees(c.Request.Context(), userID, int32(eventID))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEventNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case errors.Is(err, service.ErrOwnerRestrictedAction):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event attendees"})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"attendees": attendees})
}
