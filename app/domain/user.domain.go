package domain

import (
  "github.com/google/uuid"
	"encoding/hex"
)

type UserDomain struct {
  Id uuid.UUID
  Name string
  Password string
  Email string
}

func (user *UserDomain) EncryptPassword() {
  passwordToBytes := []byte(user.Password)
	user.Password = hex.EncodeToString(passwordToBytes)
}

func (user *UserDomain) AddId() {
  user.Id = uuid.New()
}
