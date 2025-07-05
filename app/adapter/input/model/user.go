package model

type CreateUserModel struct {
	Email string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
  Name string `json:"name" binding:"required,min=3,max=100"`
	Phone string `json:"phone" binding:"required,min=11,max=11"`
}

type UpdateUserModel struct {
	Email string `json:"email" binding:"omitempty,email"`
  Name string `json:"name" binding:"omitempty,min=3,max=100"`
	Phone string `json:"phone" binding:"omitempty,min=11,max=11"`
}

type UpdateUserPasswordModel struct {
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
}

type UserLoginModel struct {
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Email string `json:"email" binding:"required,email"`
}
