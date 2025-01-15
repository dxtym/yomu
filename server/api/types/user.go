package types

type CreateUserRequest struct {
	Id        uint64 `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
}
