package port

import (
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(domain.UserDomain) (uuid.UUID, error)
	FindUserByPhone(phone string) (domain.UserDomain, error)
	FindUserByEmail(phone string) (domain.UserDomain, error)
	List(uuid.UUID) (domain.UserDomain, error)
	ListAll() ([]domain.UserDomain, error)
	Delete(uuid.UUID) (uuid.UUID, error)
	Update(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
	UpdatePassword(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
}
