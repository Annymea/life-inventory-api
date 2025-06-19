package models

type EntryDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	PlannedDate string `json:"plannedDate"`
}

//Sobald ich die ID autogeneriere
//type CreateEntryDto struct {
//	Title       string `json:"title"`
//	Done        bool   `json:"done"`
//	PlannedDate string `json:"plannedDate"`
//}
