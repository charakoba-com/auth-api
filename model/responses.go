package model

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

// GetAlgorithmResponse is a response type returned from GetAlgorithmHandler
type GetAlgorithmResponse struct {
	Algorithm string `json:"algorithm"`
}
