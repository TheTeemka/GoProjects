package models

import "errors"

var (
	ErrTokenNotFound       = errors.New("token not found")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")

	ErrInvalidUserRole    = errors.New("user role not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrPasswordMismatch   = errors.New("password mismatch")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
