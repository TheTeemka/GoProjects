package handlers

import (
	"os"

	"github.com/TheTeemka/GoProjects/hw_6/middlewares"
	"github.com/TheTeemka/GoProjects/hw_6/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/TheTeemka/GoProjects/hw_6/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(userHandler *AuthHandler, attHandler *AttendanceHandler, scheduleHandler *ScheduleHandler, jwtService *services.JWTService) *echo.Echo {
	e := echo.New()
	e.GET("/health", HealthCheck)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}\x1b[32m${status}\t\x1b[0m\t\x1b[36m${method}\x1b[0m ${uri} ${latency_human} ${error}\n",
		Output: os.Stdout,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) {
			return true, nil
		},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	authMiddleware := middlewares.AuthMiddleware(jwtService)
	{
		attroup := e.Group("/attendance")
		attroup.POST("/subject", attHandler.CreateAttendance)
		attroup.GET("/attendanceByStudentId/:student_id", attHandler.GetAllAttendanceByStudentID)
		attroup.GET("/attendanceBySubjectId/:subject_id", attHandler.GetAllAttendanceBySubjectID)
	}

	{
		schGroup := e.Group("/schedule")
		schGroup.GET("/student/:id", scheduleHandler.GetForStudent)
		schGroup.GET("/schedule/group/:id", scheduleHandler.GetForGroup)
		schGroup.GET("/all_class_schedule", scheduleHandler.GetForAll)
	}

	{
		authGroup := e.Group("/auth")
		authGroup.POST("/register", userHandler.CreateUser)
		authGroup.POST("/login", userHandler.Login)
		authGroup.POST("/logout", userHandler.Logout, authMiddleware)
		authGroup.GET("/users/me", userHandler.GetMe, authMiddleware)
		authGroup.POST("/refresh", userHandler.RefreshAccessToken)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.HTTPErrorHandler = ErrorHandler
	return e
}
