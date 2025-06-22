package models

type EntryDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	PlannedDate string `json:"plannedDate"`
}

type CreateEntryDto struct {
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	PlannedDate string `json:"plannedDate"`
}
