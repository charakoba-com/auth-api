package model

// CreateUserRequest represents a request for create user
type CreateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LookupUserRequest represents a request for read user
type LookupUserRequest struct {
	ID string `json:"id"`
}

// UpdateUserRequest represents a request for update upser
type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// DeleteUserRequest represents a request for delete user
type DeleteUserRequest struct {
	ID string `json:"id"`
}

// AuthRequest represents a request for authenticate user
type AuthRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
