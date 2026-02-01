package errs

import "errors"

var (
	ErrTokenNotFound       = errors.New("token not found")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")

	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidUserRole    = errors.New("user role not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrPasswordMismatch   = errors.New("password mismatch")
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrGroupNotFound   = errors.New("group not found")
	ErrStudentNotFound = errors.New("student not found")
)
