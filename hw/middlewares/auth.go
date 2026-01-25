package middlewares

import (
	"strings"

	"github.com/TheTeemka/GoProjects/hw_6/services"
	"github.com/TheTeemka/GoProjects/hw_6/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtService *services.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(401, map[string]string{"error": "Unauthorized"})
			}

			fields := strings.Fields(authHeader)
			if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
				return c.JSON(401, map[string]string{"error": "Unauthorized"})
			}

			token := fields[1]
			userClaims, err := jwtService.ParseToken(token)
			if err != nil {
				return c.JSON(401, map[string]string{"error": "Unauthorized"})
			}
			utils.SetUserClaims(c, userClaims)

			return next(c)
		}
	}
}
