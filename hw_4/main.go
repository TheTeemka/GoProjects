package main

import (
	"context"
	"log"

	"github.com/temirlanbayangazy/GoProjects/hw_4/database"
	"github.com/temirlanbayangazy/GoProjects/hw_4/handlers"
	"github.com/temirlanbayangazy/GoProjects/hw_4/repository"
	"github.com/temirlanbayangazy/GoProjects/hw_4/services"

	"github.com/labstack/echo"
)

func main() {
	const (
		port     = ":8020"
		dbString = "postgres://dean:dean_password@localhost:5432/university"
	)

	conn := database.OpenConnection(dbString)
	defer conn.Close(context.Background())
	log.Println("Database connected")

	e := echo.New()

	scheduleRepo := repository.NewScheduleRepository(conn)
	scheduleService := services.NewScheduleService(scheduleRepo)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)
	scheduleHandler.RegisterRoutees(e)

	attendanceRepo := repository.NewAttendanceRepository(conn)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)
	attendanceHandler.RegisterRoutees(e)

	log.Println("Starting server on port", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("Finished execution")
}
