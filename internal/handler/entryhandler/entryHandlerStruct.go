package entryhandler

import "gorm.io/gorm"

type EntryHandler struct {
	DB *gorm.DB
}

func NewEntryHandler(db *gorm.DB) *EntryHandler {
	return &EntryHandler{db}
}
