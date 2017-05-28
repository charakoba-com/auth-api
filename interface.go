package authapi

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
