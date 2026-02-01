package handlers

import (
	"context"
	"strconv"

	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/TheTeemka/GoProjects/hw/services"
	"github.com/labstack/echo/v4"
)

type GroupHandler struct {
	service *services.GroupService
}

func NewGroupHandler(service *services.GroupService) *GroupHandler {
	return &GroupHandler{service: service}
}

// CreateGroup godoc
// @Summary Create a new group
// @Tags Groups
// @Accept json
// @Produce json
// @Param request body models.CreateGroupRequest true "Create group request"
// @Success 201 {string} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups [post]
func (h *GroupHandler) CreateGroup(c echo.Context) error {
	var req models.CreateGroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if _, err := h.service.CreateGroup(context.Background(), &req); err != nil {
		return err
	}

	return c.NoContent(201)
}

// GetGroupByID godoc
// @Summary Get group by ID
// @Tags Groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} models.GroupDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups/{id} [get]
func (h *GroupHandler) GetGroupByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid id"})
	}

	dto, err := h.service.GetGroupByID(context.Background(), id)
	if err != nil {
		return err
	}

	if dto == nil {
		return c.JSON(404, map[string]string{"error": "group not found"})
	}

	return c.JSON(200, dto)
}
