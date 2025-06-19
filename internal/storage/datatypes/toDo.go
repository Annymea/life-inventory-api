package datatypes

type Entry struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Done        bool
	PlannedDate string
}
