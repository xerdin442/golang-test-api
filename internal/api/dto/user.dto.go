package dto

import "mime/multipart"

type SignupRequest struct {
	Email        string                `form:"email" binding:"required,email"`
	Password     string                `form:"password" binding:"required,min=8"`
	Name         string                `form:"name" binding:"required"`
	ProfileImage *multipart.FileHeader `form:"profile_image"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
