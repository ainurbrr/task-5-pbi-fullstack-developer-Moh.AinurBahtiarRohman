package user

type RegisterUserInput struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Phone_Number string `json:"phone_number" binding:"required"`
	Password     string `json:"password" binding:"required,min=6"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}