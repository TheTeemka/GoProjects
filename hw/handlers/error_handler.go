package handlers

import (
	"errors"
	"net/http"

	"github.com/TheTeemka/GoProjects/hw/errs"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	// Handle Echo's built-in HTTPError
	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		c.JSON(httpErr.Code, map[string]interface{}{
			"error": httpErr.Message,
		})
		return
	}

	c.JSON(ErrToStatusCode(err), map[string]string{
		"err": err.Error(),
	})
}

func ErrToStatusCode(err error) int {
	newErr := errors.Unwrap(err)
	if newErr != nil {
		return ErrToStatusCode(newErr)
	}

	switch err {
	case errs.ErrUserNotFound, errs.ErrStudentNotFound, errs.ErrTokenNotFound:
		return http.StatusNotFound
	case errs.ErrPasswordMismatch, errs.ErrInvalidCredentials:
		return http.StatusUnauthorized
	case errs.ErrUserAlreadyExists:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
