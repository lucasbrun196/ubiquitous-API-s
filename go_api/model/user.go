package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	User     string `json:"user"`
	Password string `json:"password"`
}
