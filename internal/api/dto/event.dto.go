package dto

import "time"

type EventRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location" binding:"required"`
	Datetime    time.Time `json:"datetime" binding:"required"`
}
