package databasetypes

type ToDo struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Done        bool
	PlannedDate string
}
