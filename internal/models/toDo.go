package models

type ToDo struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	PlannedDate string `json:"plannedDate"`
}
