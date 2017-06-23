package model

// User represents an user
type User struct {
	ID       string `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"is_admin"`
}

// UserList type
type UserList []User
