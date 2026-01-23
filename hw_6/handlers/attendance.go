package handlers

import (
	"strconv"

	"github.com/TheTeemka/GoProjects/hw_6/models"
	"github.com/TheTeemka/GoProjects/hw_6/services"
	"github.com/labstack/echo/v4"
)

type AttendanceHandler struct {
	attendanceService *services.AttendanceService
}

func NewAttendanceHandler(attendanceService *services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

// CreateAttendance godoc
// @Summary Create attendance record
// @Tags Attendance
// @Accept json
// @Produce json
// @Param request body models.CreateAttendanceRequest true "Create attendance request"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/subject [post]
func (ah *AttendanceHandler) CreateAttendance(c echo.Context) error {
	var req models.CreateAttendanceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	dto, err := req.ToCreateAttendanceDTO()
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := ah.attendanceService.CreateAttendance(dto); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, map[string]string{"message": "Attendance record created successfully"})
}

// GetAllAttendanceByStudentID godoc
// @Summary Get all attendance records for a student
// @Tags Attendance
// @Produce json
// @Param student_id path int true "Student ID"
// @Success 200 {array} models.AttendanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/attendanceByStudentId/{student_id} [get]
func (ah *AttendanceHandler) GetAllAttendanceByStudentID(c echo.Context) error {
	studentIDStr := c.Param("student_id")

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid student ID"})
	}

	attendanceRecords, err := ah.attendanceService.GetAllAttendanceByStudentID(studentID)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	var resp []*models.AttendanceResponse
	for _, record := range attendanceRecords {
		resp = append(resp, record.ToAttendanceResponse())
	}

	return c.JSON(200, resp)
}

// GetAllAttendanceBySubjectID godoc
// @Summary Get all attendance records for a subject
// @Tags Attendance
// @Produce json
// @Param subject_id path int true "Subject ID"
// @Success 200 {array} models.AttendanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/attendanceBySubjectId/{subject_id} [get]
func (ah *AttendanceHandler) GetAllAttendanceBySubjectID(c echo.Context) error {
	subjectIDStr := c.Param("subject_id")

	subjectID, err := strconv.Atoi(subjectIDStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid subject ID"})
	}

	attendanceRecords, err := ah.attendanceService.GetAllAttendanceBySubjectID(subjectID)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	var resp []*models.AttendanceResponse
	for _, record := range attendanceRecords {
		resp = append(resp, record.ToAttendanceResponse())
	}

	return c.JSON(200, resp)
}
