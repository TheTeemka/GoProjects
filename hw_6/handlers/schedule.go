package handlers

import (
	"context"
	"strconv"

	"github.com/TheTeemka/GoProjects/hw_6/services"

	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	service *services.ScheduleService
}

func NewScheduleHandler(service *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		service: service,
	}
}

// GetForStudent godoc
// @Summary Get schedule for a student
// @Tags Schedule
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {array} models.Schedule
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /schedule/student/{id} [get]
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

// GetForGroup godoc
// @Summary Get schedule for a group
// @Tags Schedule
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {array} models.Schedule
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /schedule/schedule/group/{id} [get]
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

// GetForAll godoc
// @Summary Get schedule for all classes
// @Tags Schedule
// @Produce json
// @Success 200 {array} models.Schedule
// @Failure 500 {object} map[string]any
// @Router /schedule/all_class_schedule [get]
func (sh *ScheduleHandler) GetForAll(c echo.Context) error {
	schedule, err := sh.service.GetAll(context.Background())
	if err != nil {
		return c.JSON(500, map[string]any{"err": err.Error()})
	}

	return c.JSON(200, schedule)
}
