package handlers

import (
	"github.com/TheTeemka/GoProjects/hw_5/middlewares"
	"github.com/TheTeemka/GoProjects/hw_5/services"
	"github.com/labstack/echo"
)

func RegisterRoutes(userHandler *UserHandler, attHandler *AttendanceHandler, scheduleHandler *ScheduleHandler, jwtService *services.JWTService) *echo.Echo {
	e := echo.New()

	authMiddleware := middlewares.AuthMiddleware(jwtService)
	{
		attroup := e.Group("/attendance")
		attroup.POST("/attendance/subject", attHandler.CreateAttendance)
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

	return e
}
