package domain

import (
  "github.com/google/uuid"
  "crypto/md5"
	"encoding/hex"
)

type UserDomain struct {
  Id uuid.UUID
  Name string
  Password string
  Email string
}

func (user *UserDomain) EncryptPassword(password string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func (user *UserDomain) CreateUser(newUser UserDomain) {
  user.Id = uuid.New()
  user.Password = user.EncryptPassword(newUser.Password)
  user.Email = newUser.Email
  user.Name = newUser.Name
}
