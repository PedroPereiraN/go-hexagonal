package helpers

import (
	"encoding/hex"
	"fmt"
	"github.com/PedroPereiraN/go-hexagonal/domain"
)

func VerifyOldPassword(user domain.UserDomain, newPassword string) (bool, error) {
  passwordByte, err := hex.DecodeString(user.Password)

  if err != nil {
    return true, err
  }

  if string(passwordByte) == newPassword {
    return true, nil
  }

  fmt.Println(newPassword)
  fmt.Println(string(passwordByte))

  return false, nil
}
