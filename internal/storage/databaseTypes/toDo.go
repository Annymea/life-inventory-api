package datatypes

type ToDo struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Done        bool
	PlannedDate string
}
