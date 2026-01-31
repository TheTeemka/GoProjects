package handlers

import (
	"context"
	"strconv"

	service "github.com/temirlanbayangazy/GoProjects/hw_4/services"

	"github.com/labstack/echo"
)

type ScheduleHandler struct {
	service *service.ScheduleService
}

func NewScheduleHandler(service *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		service: service,
	}
}

func (sh *ScheduleHandler) GetForStudent(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]any{"err": err.Error()})
	}

	schedule, err := sh.service.GetForStudent(context.Background(), id)
	if err != nil {
		return c.JSON(500, map[string]any{"err": err.Error()})
	}

	return c.JSON(200, schedule)

}

func (sh *ScheduleHandler) GetForGroup(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]any{"err": err.Error()})
	}

	schedule, err := sh.service.GetForGroup(context.Background(), id)
	if err != nil {
		return c.JSON(500, map[string]any{"err": err.Error()})
	}

	return c.JSON(200, schedule)
}

func (sh *ScheduleHandler) GetForAll(c echo.Context) error {
	schedule, err := sh.service.GetAll(context.Background())
	if err != nil {
		return c.JSON(500, map[string]any{"err": err.Error()})
	}

	return c.JSON(200, schedule)
}

func (sh *ScheduleHandler) RegisterRoutees(e *echo.Echo) *echo.Echo {
	e.GET("/student/:id", sh.GetForStudent)
	e.GET("/schedule/group/:id", sh.GetForGroup)
	e.GET("/all_class_schedule", sh.GetForAll)

	return e
}
