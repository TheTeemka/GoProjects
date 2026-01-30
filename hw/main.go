package main

import (
	"context"
	"log"

	"github.com/TheTeemka/GoProjects/hw_6/config"
	"github.com/TheTeemka/GoProjects/hw_6/database"
	"github.com/TheTeemka/GoProjects/hw_6/handlers"
	"github.com/TheTeemka/GoProjects/hw_6/repository"
	"github.com/TheTeemka/GoProjects/hw_6/services"
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

	userRepo := repository.NewUserRepository(conn)
	userService := services.NewUserService(userRepo, jwtService, refreshTokenService)
	userHandler := handlers.NewUserHandler(userService)

	e := handlers.RegisterRoutes(userHandler, attendanceHandler, scheduleHandler, jwtService)
	log.Println("Starting server on port", cfg.Port)
	if err := e.Start(cfg.Port); err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("Finished execution")
}
