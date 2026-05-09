package main

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserUpdate handler explicitly maps only permitted fields
func UserUpdate(req UpdateUserRequest) {
    // Only update DB fields present in req
}