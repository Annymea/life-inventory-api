package models

type UserDto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
