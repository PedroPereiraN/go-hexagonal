package service

import(
  "github.com/PedroPereiraN/go-hexagonal/port/output"
  "github.com/PedroPereiraN/go-hexagonal/domain"
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
}

type userService struct {
  repository port.UserRepository
}

func (service *userService) Create(dto domain.UserDomain) (string, error) {
  uDomain := domain.UserDomain{
    Name: dto.Name,
    Email: dto.Email,
    Password: dto.Password,
    Id: uuid.New(),
  }

  newUserId, err := service.repository.Create(uDomain)

  if err != nil {
    return "", err
  }

  successMessage := "Usuário criado com sucesso: " + newUserId.String()

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
