package model

// User represents an user
type User struct {
	ID       string `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}
