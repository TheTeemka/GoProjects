package main

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	"github.com/TheTeemka/GoProjects/hw_5/database"
	"github.com/TheTeemka/GoProjects/hw_5/handlers"
	"github.com/TheTeemka/GoProjects/hw_5/repository"
	"github.com/TheTeemka/GoProjects/hw_5/services"
)

func main() {
	const (
		port     = ":8020"
		dbString = "postgres://dean:dean_password@localhost:5432/university"
		jwtTTL   = 30 * time.Minute
	)
	var (
		secretKey = decodeBase64("N33BOcxkBlFmYmk3imxTvlIWp6MwgKc83Xv+hw+11ns=")
	)

	conn := database.OpenConnection(dbString)
	defer conn.Close(context.Background())
	log.Println("Database connected")

	jwtService := services.NewJWTService(secretKey, jwtTTL)

	scheduleRepo := repository.NewScheduleRepository(conn)
	scheduleService := services.NewScheduleService(scheduleRepo)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)

	attendanceRepo := repository.NewAttendanceRepository(conn)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)

	userRepo := repository.NewUserRepository(conn)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, jwtService)

	e := handlers.RegisterRoutes(userHandler, attendanceHandler, scheduleHandler, jwtService)
	log.Println("Starting server on port", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("Finished execution")
}

func decodeBase64(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalf("Failed to decode base64 string: %v", err)
	}
	return data
}
