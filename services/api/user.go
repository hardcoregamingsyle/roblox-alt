package main

// F-24: Explicit DTO to prevent Mass Assignment
type UpdateUserRequest struct {
	DisplayName string `json:"displayName"`
	Bio         string `json:"bio"`
}

func UpdateProfile(req UpdateUserRequest) {
    // Only bind allowed fields
}