package request

type CreateNewUserRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" validate:"required,min=3,max=100,email"`
	Password string `json:"password" form:"password" validate:"required,min=3,max=100"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type RegisterNewUserRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" validate:"required,min=3,max=100,email"`
	Password string `json:"password" form:"password" validate:"required,min=3,max=100"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type UpdateUserRequest struct {
	ID       int    `validate:"required"`
	Username string `json:"username" form:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" validate:"required,min=3,max=100,email"`
	Password string `json:"password" form:"password" validate:"required,min=3,max=100"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,min=3,max=100,email"`
	Password string `json:"password" form:"password" validate:"required,min=3,max=100"`
}
