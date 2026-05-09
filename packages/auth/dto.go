package auth

// RegisterRequest uses explicit fields to prevent Mass Assignment (LF-03)
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserDTO ensures only authorized fields are exposed
type UserDTO struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}