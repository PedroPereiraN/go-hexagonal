package port

import (
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
)

type UserService interface {
	Create(domain.UserDomain) (uuid.UUID, error)
	List(uuid.UUID) (domain.UserDomain, error)
	ListAll() ([]domain.UserDomain, error)
	Delete(uuid.UUID) (uuid.UUID, error)
	Update(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
	UpdatePassword(uuid.UUID, string) (uuid.UUID, error)
	Login(string, string) (string, error)
}
