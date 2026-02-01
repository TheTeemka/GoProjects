package handlers

import (
	"context"
	"strconv"

	"github.com/TheTeemka/GoProjects/hw/models"

	"github.com/labstack/echo/v4"
)

type IScheduleService interface {
	GetForStudent(ctx context.Context, studentID int) ([]*models.ScheduleDTO, error)
	GetForGroup(ctx context.Context, groupID int) ([]*models.ScheduleDTO, error)
	GetAll(ctx context.Context) ([]*models.ScheduleDTO, error)
	CreateSchedule(ctx context.Context, req *models.CreateScheduleRequest) error
	DeleteSchedule(ctx context.Context, id int) error
}

// DeleteSchedule godoc
// @Summary Delete a schedule entry
// @Tags Schedule
// @Param id path int true "Schedule ID"
// @Success 204
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /schedules/{id} [delete]
func (sh *ScheduleHandler) DeleteSchedule(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]any{"err": err.Error()})
	}

	if err := sh.service.DeleteSchedule(context.Background(), id); err != nil {
		return c.JSON(500, map[string]any{"err": err.Error()})
	}

	return c.NoContent(204)
}

type ScheduleHandler struct {
	service IScheduleService
}

func NewScheduleHandler(service IScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		service: service,
	}
}

// GetForStudent godoc
// @Summary Get schedule for a student
// @Tags Schedule
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {array} models.ScheduleDTO
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
		return err
	}

	return c.JSON(200, schedule)

}

// GetForGroup godoc
// @Summary Get schedule for a group
// @Tags Schedule
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {array} models.ScheduleDTO
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
		return err
	}

	if len(schedule) == 0 {
		schedule = make([]*models.ScheduleDTO, 0)
	}
	return c.JSON(200, schedule)
}

// GetForAll godoc
// @Summary Get schedule for all classes
// @Tags Schedule
// @Produce json
// @Success 200 {array} models.ScheduleDTO
// @Failure 500 {object} map[string]any
// @Router /schedule/all_class_schedule [get]
func (sh *ScheduleHandler) GetForAll(c echo.Context) error {
	schedule, err := sh.service.GetAll(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(200, schedule)
}

// CreateSchedule godoc
// @Summary Create a class schedule entry
// @Tags Schedule
// @Accept json
// @Produce json
// @Param request body models.CreateScheduleRequest true "Create schedule request"
// @Success 201 {object} models.ScheduleDTO
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /schedule [post]
func (sh *ScheduleHandler) CreateSchedule(c echo.Context) error {
	var req models.CreateScheduleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]any{"err": err.Error()})
	}

	err := sh.service.CreateSchedule(context.Background(), &req)
	if err != nil {
		return err
	}

	return c.NoContent(201)
}
