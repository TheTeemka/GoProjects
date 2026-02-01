package handlers

import (
	"os"

	"github.com/TheTeemka/GoProjects/hw/middlewares"
	"github.com/TheTeemka/GoProjects/hw/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/TheTeemka/GoProjects/hw/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(userHandler *AuthHandler, attHandler *AttendanceHandler, scheduleHandler *ScheduleHandler, studentsHandler *StudentsHandler, groupHandler *GroupHandler, jwtService *services.JWTService) *echo.Echo {
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
		schGroup := e.Group("/schedules")
		schGroup.POST("", scheduleHandler.CreateSchedule)
		schGroup.GET("/student/:id", scheduleHandler.GetForStudent)
		schGroup.GET("/group/:id", scheduleHandler.GetForGroup)
		schGroup.GET("/all_class_schedule", scheduleHandler.GetForAll)
		schGroup.DELETE("/:id", scheduleHandler.DeleteSchedule)
	}

	{
		stGroup := e.Group("/students")
		stGroup.POST("", studentsHandler.CreateStudent)
		stGroup.GET("/:id", studentsHandler.GetStudentByID)
		stGroup.GET("", studentsHandler.ListStudents)
		stGroup.GET("/group/:id", studentsHandler.GetStudentsByGroupID)
		stGroup.PUT("/:id", studentsHandler.UpdateStudent)
		stGroup.DELETE("/:id", studentsHandler.DeleteStudent)
	}
	{
		grGroup := e.Group("/groups")
		grGroup.POST("", groupHandler.CreateGroup)
		grGroup.GET("/:id", groupHandler.GetGroupByID)
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
