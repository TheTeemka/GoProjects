package handlers

import (
	"errors"
	"net/http"

	"github.com/TheTeemka/GoProjects/hw_6/models"
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

	// Handle custom application errors with specific status codes
	switch {
	case errors.Is(err, models.ErrUserNotFound):
		c.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	case errors.Is(err, models.ErrPasswordMismatch):
		c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	case errors.Is(err, models.ErrInvalidUserRole):
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user role",
		})
	default:
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
}
