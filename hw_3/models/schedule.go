package models

type Schedule struct {
	ID        int
	Subject   string
	DayOfWeek string
	Time      string
	GroupID   int
}
