package models

type ToDoDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	PlannedDate string `json:"plannedDate"`
}

//Sobald ich die ID autogeneriere
//type CreateToDoDTO struct {
//	Title       string `json:"title"`
//	Done        bool   `json:"done"`
//	PlannedDate string `json:"plannedDate"`
//}
