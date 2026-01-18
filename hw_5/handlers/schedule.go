package handlers

import (
	"context"
	"strconv"

	"github.com/TheTeemka/GoProjects/hw_5/services"

	"github.com/labstack/echo"
)

type ScheduleHandler struct {
	service *services.ScheduleService
}

func NewScheduleHandler(service *services.ScheduleService) *ScheduleHandler {
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
