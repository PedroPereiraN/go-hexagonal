package port

import (
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
  Create(domain.UserDomain) (uuid.UUID, error)
  List(uuid.UUID) (domain.UserDomain, error)
  ListAll() ([]domain.UserDomain, error)
  Update(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
}
