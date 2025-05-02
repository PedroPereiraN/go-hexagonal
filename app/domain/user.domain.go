package domain

import "github.com/google/uuid"

type UserDomain struct {
  Id uuid.UUID
  Name string
  Password string
  Email string
}
