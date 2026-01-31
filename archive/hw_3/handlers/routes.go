package handlers

import "github.com/labstack/echo"

func RegisterRoutes(sh *ScheduleHandler) *echo.Echo {
	e := echo.New()
	e.GET("/student/:id", sh.GetForStudent)
	e.GET("/schedule/group/:id", sh.GetForGroup)
	e.GET("/all_class_schedule", sh.GetForAll)

	return e
}
