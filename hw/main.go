package main

import (
	"context"
	"log"

	"github.com/TheTeemka/GoProjects/hw/config"
	"github.com/TheTeemka/GoProjects/hw/database"
	"github.com/TheTeemka/GoProjects/hw/handlers"
	"github.com/TheTeemka/GoProjects/hw/repository"
	"github.com/TheTeemka/GoProjects/hw/services"
)

func main() {
	cfg := config.GetConfig()

	conn := database.OpenConnection(cfg.DB.String())
	defer conn.Close(context.Background())
	log.Println("Database connected")

	sqlDB := database.PGXConnToSQLDB(conn)
	database.GooseMigrate(sqlDB, "./database/migrations")
	log.Println("Database migrated")

	jwtService := services.NewJWTService(cfg.JWT.SecretKey, cfg.JWT.TTL)

	tokenRepo := repository.NewTokenRepository(conn)
	refreshTokenService := services.NewTokenService(tokenRepo, cfg.RefreshToken.Size, cfg.RefreshToken.Type, cfg.RefreshToken.TTL)

	scheduleRepo := repository.NewScheduleRepository(conn)
	scheduleService := services.NewScheduleService(scheduleRepo)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)

	attendanceRepo := repository.NewAttendanceRepository(conn)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)

	studentRepo := repository.NewStudentRepository(conn)
	studentService := services.NewStudentsService(studentRepo)
	studentHandler := handlers.NewStudentsHandler(studentService)

	groupRepo := repository.NewGroupRepository(conn)
	groupService := services.NewGroupService(groupRepo)
	groupHandler := handlers.NewGroupHandler(groupService)

	userRepo := repository.NewUserRepository(conn)
	userService := services.NewAuthService(userRepo, jwtService, refreshTokenService)
	userHandler := handlers.NewUserHandler(userService)

	e := handlers.RegisterRoutes(userHandler, attendanceHandler, scheduleHandler, studentHandler, groupHandler, jwtService)

	log.Println("Starting server on port", cfg.Port)
	if err := e.Start(cfg.Port); err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("Finished execution")
}
