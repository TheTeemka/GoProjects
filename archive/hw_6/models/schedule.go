package models

type Schedule struct {
	ID        int    `json:"id"`
	Subject   string `json:"subject"`
	DayOfWeek string `json:"day_of_week"`
	Time      string `json:"time"`
	GroupID   int    `json:"group_id"`
}
