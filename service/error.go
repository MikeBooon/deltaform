package service

import "errors"

var (
	ErrUserNotFoundWithEmail = errors.New("User with supplied email not found")
)
