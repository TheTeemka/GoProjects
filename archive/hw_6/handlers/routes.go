package handlers

import (
	"github.com/TheTeemka/GoProjects/hw_6/middlewares"
	"github.com/TheTeemka/GoProjects/hw_6/services"
	"github.com/labstack/echo/v4"

	_ "github.com/TheTeemka/GoProjects/hw_6/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(userHandler *UserHandler, attHandler *AttendanceHandler, scheduleHandler *ScheduleHandler, jwtService *services.JWTService) *echo.Echo {
	e := echo.New()

	e.GET("/health", HealthCheck)

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
		authGroup.GET("/user/login", userHandler.Login)
		authGroup.GET("/users/me", userHandler.GetMe, authMiddleware)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
