package service

import "errors"

var (
	ErrInvalidDate = errors.New("Event start date cannot be in the past")
)
