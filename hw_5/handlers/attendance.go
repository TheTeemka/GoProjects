package handlers

import (
	"strconv"

	"github.com/TheTeemka/GoProjects/hw_5/models"
	"github.com/TheTeemka/GoProjects/hw_5/services"
	"github.com/labstack/echo"
)

type AttendanceHandler struct {
	attendanceService *services.AttendanceService
}

func NewAttendanceHandler(attendanceService *services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

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
