package service

import (
	"errors"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/ports/output"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("super-secret")

func NewUserService(repository port.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

type UserService interface {
  Create(domain.UserDomain) (uuid.UUID, error)
	List(uuid.UUID) (domain.UserDomain, error)
	ListAll() ([]domain.UserDomain, error)
	Delete(uuid.UUID) (uuid.UUID, error)
	Update(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
	UpdatePassword(uuid.UUID, string) (uuid.UUID, error)
	Login(string, string) (string, error)
}

type userService struct {
	repository port.UserRepository
}

func (service * userService) Create(dto domain.UserDomain) (uuid.UUID, error) {
	uDomain, err := domain.CreateUser(
		dto.Id,
		dto.Name,
		dto.Email,
		dto.Phone,
		dto.Password,
		dto.CreatedAt,
		dto.UpdatedAt,
		dto.DeletedAt,
	)

	if err != nil {
		return uuid.Nil, err
	}

	//check phone
	_, err = service.repository.FindUserByPhone(uDomain.Phone)

	// first we check if an error exists
	// if the error exists and is NOT because the value could not be found we return the error
	// if the error is because the value doesn't exists we ignore it
	if err != nil && err.Error() != "sql: no rows in result set" {
		return uuid.Nil, err
	}

	// second we check if errors doesnt exists, because if the error doesn't exists an user with this phone already exists
	// so we cant let this user be created
	if err == nil {
		return uuid.Nil, errors.New("Phone is already registered")
	}

	// check email
	_, err = service.repository.FindUserByEmail(uDomain.Email)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return uuid.Nil, err
	}

	if err == nil {
		return uuid.Nil, errors.New("Email is already registered")
	}

	result, err := service.repository.Create(uDomain)

	if err != nil {
		return uuid.Nil, err
	}

	return result, err
}

func (service *userService) List(id uuid.UUID) (domain.UserDomain, error) {
	userData, err := service.repository.List(id)

	if err != nil {
		return domain.UserDomain{}, err
	}

	uDomain, err := domain.CreateUser(
		userData.Id,
		userData.Name,
		userData.Email,
		userData.Phone,
		userData.Password,
		userData.CreatedAt,
		userData.UpdatedAt,
		userData.DeletedAt,
	)

	if err != nil {
		return domain.UserDomain{}, err
	}

	return uDomain, nil
}

func (service *userService) ListAll() ([]domain.UserDomain, error) {
	var users []domain.UserDomain
	usersData, err := service.repository.ListAll()

	if err != nil {
		return []domain.UserDomain{}, err
	}

	for _, userData := range usersData {
		uDomain, err := domain.CreateUser(
			userData.Id,
			userData.Name,
			userData.Email,
			userData.Phone,
			userData.Password,
			userData.CreatedAt,
			userData.UpdatedAt,
			userData.DeletedAt,
		)

		if err != nil {
			return []domain.UserDomain{}, err
		}

		users = append(users, uDomain)
	}

	return users, nil
}

func (service *userService) Delete(id uuid.UUID) (uuid.UUID, error) {
	_, err := service.repository.List(id)

	if err != nil {
		return uuid.Nil, err
	}

	userId, err := service.repository.Delete(id)

  if err != nil {
    return uuid.Nil, err
  }

  return userId, nil

}

func (service *userService) Update(id uuid.UUID, dto domain.UserDomain) (uuid.UUID, error) {
	_, err := service.repository.List(id)

	if err != nil {
		return uuid.Nil, err
	}

	uDomain, err := domain.CreateUser(
		dto.Id,
		dto.Name,
		dto.Email,
		dto.Phone,
		dto.Password,
		dto.CreatedAt,
		dto.UpdatedAt,
		dto.DeletedAt,
	)

	if err != nil {
		return uuid.Nil, err
	}

	//check phone
	_, err = service.repository.FindUserByPhone(uDomain.Phone)

	// first we check if an error exists
	// if the error exists and is NOT because the value could not be found we return the error
	// if the error is because the value doesn't exists we ignore it
	if err != nil && err.Error() != "sql: no rows in result set" {
		return uuid.Nil, err
	}

	// second we check if errors doesnt exists, because if the error doesn't exists an user with this phone already exists
	// so we cant let this user be created
	if err == nil {
		return uuid.Nil, errors.New("Phone is already registered")
	}

	// check email
	_, err = service.repository.FindUserByEmail(uDomain.Email)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return uuid.Nil, err
	}

	if err == nil {
		return uuid.Nil, errors.New("Email is already registered")
	}

	userId, err := service.repository.Update(id, uDomain)

	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func (service *userService) UpdatePassword(id uuid.UUID, password string) (uuid.UUID, error) {
	uDomain, err := domain.CreateUser(
		uuid.Nil,
		"",
		"",
		"",
		password,
		time.Time{},
		time.Time{},
		time.Time{},
	)

	if err != nil {
		return uuid.Nil, err
	}

	userId, err := service.repository.UpdatePassword(id, uDomain)

  if err != nil {
    return uuid.Nil, err
  }

  return userId, nil
}

func (service *userService) Login(email string, password string) (string, error) {

	user, err := service.repository.FindUserByEmail(email)

	if err != nil {
    return "", err
  }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Wrong password")
	}

	claims := jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
