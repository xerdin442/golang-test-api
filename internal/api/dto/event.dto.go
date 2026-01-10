package dto

import "time"

type CreateEventRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location" binding:"required"`
	Datetime    time.Time `json:"datetime" binding:"required"`
}

type UpdateEventRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Datetime    time.Time `json:"datetime"`
}
