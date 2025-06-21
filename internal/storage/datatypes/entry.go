package datatypes

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entry struct {
	ID          string `gorm:"primaryKey;unique"`
	Title       string
	Done        bool
	PlannedDate string
}

func (e *Entry) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}
