package db

import "time"

// User represents API user including admin and regular user
type User struct {
	key        string
	Name       string
	Password   string
	CreatedOn  time.Time
	ModifiedOn time.Time
}

// CreateUserRequest represents request body of Create User
type CreateUserRequest struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}
