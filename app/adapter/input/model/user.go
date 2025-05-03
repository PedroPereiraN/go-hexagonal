package model

type UserRequestModel struct {
  Email string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
  Name string `json:"name" binding:"required,min=3,max=100"`
}

type UserUpdateModel struct {
  Email string `json:"email,omitempty" binding:"omitempty,email"`
  Password string `json:"password,omitempty" binding:"omitempty,min=6,containsany=!@#$%*"`
  Name string `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
}
