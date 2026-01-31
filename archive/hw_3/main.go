package main

import (
	"context"
	"log"
	"x/database"
	"x/handlers"
	"x/repository"
	service "x/services"
)

func main() {
	const (
		port     = ":8020"
		dbString = "postgres://dean:dean_password@localhost:5432/university"
	)

	conn := database.OpenConnection(dbString)
	defer conn.Close(context.Background())
	log.Println("Database connected")

	scheduleRepo := repository.NewScheduleRepository(conn)
	scheduleService := service.NewScheduleService(scheduleRepo)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)

	e := handlers.RegisterRoutes(scheduleHandler)

	log.Println("Starting server on port", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("Finished execution")
}
