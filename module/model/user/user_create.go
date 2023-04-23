package user

type UserCreate struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      uint64 `json:"age" validate:"required,min=9"`
}
