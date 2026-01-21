package service

import "errors"

var (
	ErrEmailAlreadyExists    = errors.New("User already exists with that email")
	ErrInvalidPassword       = errors.New("Invalid password")
	ErrInvalidEmail          = errors.New("Invalid email address")
	ErrFileUploadFailed      = errors.New("File upload failed")
	ErrUnsupportedImageType  = errors.New("Unsupported image type: only JPEG, PNG and HEIC are allowed")
	ErrInvalidDate           = errors.New("Event start date cannot be in the past")
	ErrOwnerRestrictedAction = errors.New("This resource can only be accessed by the event owner")
	ErrEventNotFound         = errors.New("No event exists with that ID")
)
