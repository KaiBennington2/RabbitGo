package dto

// CreateUserDTO
// DTO to create a user type entity.
type CreateUserDTO struct {
	ID   int    `json:"id"`
	Code string `json:"user_code"`
	Name string `json:"name"`
}

// UpdateUserPutDTO
// DTO to update the body of a user-type entity.
type UpdateUserPutDTO CreateUserDTO
