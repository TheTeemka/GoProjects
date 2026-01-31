package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/TheTeemka/GoProjects/hw/services"
	"github.com/labstack/echo/v4"
)

type StudentsHandler struct {
	service *services.StudentsService
}

func NewStudentsHandler(service *services.StudentsService) *StudentsHandler {
	return &StudentsHandler{service: service}
}

// CreateStudent godoc
// @Summary Create a new student
// @Tags Students
// @Accept json
// @Produce json
// @Param request body models.CreateStudentRequest true "Create student request"
// @Success 201 {string} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students [post]
func (h *StudentsHandler) CreateStudent(c echo.Context) error {
	var req models.CreateStudentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.service.CreateStudent(context.Background(), &req); err != nil {
		return err
	}

	return c.NoContent(201)
}

// GetStudentByID godoc
// @Summary Get student by ID
// @Tags Students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.StudentDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [get]
func (h *StudentsHandler) GetStudentByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid id"})
	}

	dto, err := h.service.GetStudentByID(context.Background(), id)
	if err != nil {
		return err
	}

	if dto == nil {
		return c.JSON(404, map[string]string{"error": "student not found"})
	}

	return c.JSON(200, dto)
}

// GetStudentsByGroupID godoc
// @Summary Get students by group ID
// @Tags Students
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {array} models.StudentDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/group/{id} [get]
func (h *StudentsHandler) GetStudentsByGroupID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid id"})
	}

	dtos, err := h.service.GetStudentsByGroupID(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(200, dtos)
}

// ListStudents godoc
// @Summary List students with optional filters
// @Tags Students
// @Produce json
// @Param name query string false "Name filter"
// @Param email query string false "Email filter"
// @Param group_id query int false "Group ID filter"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.StudentDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students [get]
func (h *StudentsHandler) ListStudents(c echo.Context) error {
	q := c.QueryParams()
	filter := &models.StudentFilter{}

	if v := q.Get("name"); v != "" {
		filter.Name = &v
	}
	if v := q.Get("email"); v != "" {
		filter.Email = &v
	}
	if v := q.Get("group_id"); v != "" {
		if gid, err := strconv.Atoi(v); err == nil {
			filter.GroupID = &gid
		} else {
			return c.JSON(400, map[string]string{"error": "invalid group_id"})
		}
	}
	if v := q.Get("limit"); v != "" {
		if lim, err := strconv.ParseUint(v, 10, 64); err == nil {
			filter.Limit = lim
		} else {
			return c.JSON(400, map[string]string{"error": "invalid limit"})
		}
	}
	if v := q.Get("offset"); v != "" {
		if off, err := strconv.ParseUint(v, 10, 64); err == nil {
			filter.Offset = off
		} else {
			return c.JSON(400, map[string]string{"error": "invalid offset"})
		}
	}

	entities, err := h.service.ListStudents(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.JSON(200, entities)
}

// UpdateStudent godoc
// @Summary Update student
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param request body models.UpdateStudentRequest true "Update student request"
// @Success 200 {string} string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [put]
func (h *StudentsHandler) UpdateStudent(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid id"})
	}

	var req models.UpdateStudentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.service.UpdateStudent(context.Background(), id, &req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "student updated"})
}

// DeleteStudent godoc
// @Summary Delete student
// @Tags Students
// @Produce json
// @Param id path int true "Student ID"
// @Success 204 {string} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [delete]
func (h *StudentsHandler) DeleteStudent(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid id"})
	}

	if err := h.service.DeleteStudent(context.Background(), id); err != nil {
		return err
	}

	return c.NoContent(204)
}
