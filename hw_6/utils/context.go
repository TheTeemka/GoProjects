package utils

import (
	"github.com/TheTeemka/GoProjects/hw_6/models"
	"github.com/labstack/echo/v4"
)

const (
	userIDKey = "userID"
)

func SetUserClaims(ctx echo.Context, userID *models.UserClaims) {
	ctx.Set(userIDKey, userID)
}

func GetUserClaims(ctx echo.Context) (*models.UserClaims, bool) {
	val := ctx.Get(userIDKey)
	if val == nil {
		return nil, false
	}

	userClaims, ok := val.(*models.UserClaims)
	if !ok {
		return nil, false
	}

	return userClaims, ok
}
