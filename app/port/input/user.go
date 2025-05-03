package port

import (
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
)

type UserService interface {
  Create(domain.UserDomain) (string, error)
  List(uuid.UUID) (domain.UserDomain, error)
  ListAll() ([]domain.UserDomain, error)
  Update(uuid.UUID, domain.UserDomain) (string, error)
  Delete(uuid.UUID) (string, error)
  UpdatePassword(uuid.UUID, string) (string, error)
}
