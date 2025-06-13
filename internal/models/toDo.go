package models

type GetToDoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	PlannedDate string `json:"plannedDate"`
}

type PostToDoRequest struct {
	ID          string `json:"id"` //Will ich eigentlich automatisch generieren
	Title       string `json:"title"`
	Done        bool   `json:"done"`
	PlannedDate string `json:"plannedDate"`
}
