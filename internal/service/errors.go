package service

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("User already exists with that email")
	ErrInvalidPassword    = errors.New("Invalid password")
	ErrInvalidEmail       = errors.New("Invalid email address")
)

var (
	ErrInvalidDate           = errors.New("Event start date cannot be in the past")
	ErrOwnerRestrictedAction = errors.New("This resource can only be accessed by the event owner")
)
