package domain

import (
	"strings"
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(
	id uuid.UUID,
	name string,
	email string,
	phone string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) (UserDomain, error) {
	uDomain := UserDomain{
		Name: name,
		Email: email,
		Phone: phone,
		Password: password,
	}

	if id == uuid.Nil {
		uDomain.Id = uuid.New()
	} else {
		uDomain.Id = id
	}

	if createdAt.IsZero() {
		uDomain.CreatedAt = time.Now()
	} else {
		uDomain.CreatedAt = createdAt
	}

	if !updatedAt.IsZero() {
		uDomain.UpdatedAt = updatedAt
	}

	if !deletedAt.IsZero() {
		uDomain.DeletedAt = deletedAt
	}

	if uDomain.IsBcryptHash(password) && password != "" {
		uDomain.Password = password
	}

	if !uDomain.IsBcryptHash(password) && password != "" {
		err := uDomain.EncryptPassword(password)

		if err != nil {
			return UserDomain{}, err
		}
	}

	return uDomain, nil
}

type UserDomain struct {
	Id uuid.UUID
	Name string
	Email string
	Phone string
	Password string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdatedAt time.Time
}

func (user *UserDomain) EncryptPassword(password string) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return nil
}

func (user *UserDomain) IsBcryptHash(s string) bool {
	return strings.HasPrefix(s, "$2a$") || strings.HasPrefix(s, "$2b$") || strings.HasPrefix(s, "$2y$")
}
