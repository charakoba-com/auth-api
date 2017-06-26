package model

// ErrorResponse is a response type returned when HTTP error is raised
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// HealthCheckResponse is a response type returned from HealthCheckHandler
type HealthCheckResponse struct {
	Message string `json:"message"`
}

// CreateUserResponse is a response type returned from CreateUserHandler
type CreateUserResponse struct {
	Message string `json:"message"`
}

// LookupUserResponse is a response type returned from LookupUserHandler
type LookupUserResponse struct {
	User User `json:"user"`
}

// UpdateUserResponse is a resposne type returned from UpdateUserHandler
type UpdateUserResponse struct {
	Message string `json:"message"`
}

// DeleteUserResponse is a response type returned from DeleteUserHandler
type DeleteUserResponse struct {
	Message string `json:"message"`
}

// ListupUserResponse is a response type returned from ListupUserHandler
type ListupUserResponse struct {
	Users UserList `json:"user"`
}

// AuthResponse is a response type returned from AuthHandler
type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// GetAlgorithmResponse is a response type returned from GetAlgorithmHandler
type GetAlgorithmResponse struct {
	Algorithm string `json:"algorithm"`
}

// GetKeyResponse is a response type returned from GetKeyHandler
type GetKeyResponse struct {
	PublicKey string `json:"publickey"`
}

// VerifyResponse is a response type returned from VerifyHandler
type VerifyResponse struct {
	Status bool `json:"status"`
}
