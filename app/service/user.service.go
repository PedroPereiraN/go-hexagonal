package service

import (
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/port/output"
	"github.com/google/uuid"
)

func NewUserService(repository port.UserRepository) UserService {
  return &userService{
    repository: repository,
  }
}

type UserService interface {
  Create(domain.UserDomain) (string, error)
  List(uuid.UUID) (domain.UserDomain, error)
  ListAll() ([]domain.UserDomain, error)
  Update(uuid.UUID, domain.UserDomain) (string, error)
  Delete(uuid.UUID) (string, error)
  UpdatePassword(uuid.UUID, string) (string, error)
}

type userService struct {
  repository port.UserRepository
}

func (service *userService) Create(dto domain.UserDomain) (string, error) {
  uDomain := domain.UserDomain{
    Name: dto.Name,
    Email: dto.Email,
    Password: dto.Password,
  }
  uDomain.EncryptPassword()
  uDomain.AddId()

  newUserId, err := service.repository.Create(uDomain)

  if err != nil {
    return "", err
  }

  successMessage := "User created successfully: " + newUserId.String()

  return successMessage, nil
}

func (service *userService) List(id uuid.UUID) (domain.UserDomain, error) {
  user, err := service.repository.List(id)

  if err != nil {
    return domain.UserDomain{}, err
  }

  return user, nil
}

func (service *userService) ListAll() ([]domain.UserDomain, error) {
  users, err := service.repository.ListAll()

  if err != nil {
    return []domain.UserDomain{}, err
  }

  return users, nil
}

func (service *userService) Update(id uuid.UUID, dto domain.UserDomain) (string, error) {
  uDomain := domain.UserDomain{
    Name: dto.Name,
    Email: dto.Email,
  }

  newUserId, err := service.repository.Update(id, uDomain)

  if err != nil {
    return "", err
  }

  successMessage := "User edited successfully: " + newUserId.String()

  return successMessage, nil
}

func (service *userService) Delete(id uuid.UUID) (string, error) {
  userId, err := service.repository.Delete(id)

  if err != nil {
    return "", err
  }

  successMessage := "User deleted successfully: " + userId.String()

  return successMessage, nil
}

func (service *userService) UpdatePassword(id uuid.UUID, password string) (string, error) {
  uDomain := domain.UserDomain{
    Password: password,
  }

  uDomain.EncryptPassword()

  newUserId, err := service.repository.UpdatePassword(id, uDomain.Password)

  if err != nil {
    return "", err
  }

  successMessage := "User password edited successfully: " + newUserId.String()

  return successMessage, nil
}
