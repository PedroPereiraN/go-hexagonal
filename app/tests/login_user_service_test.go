package test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/services"
	"github.com/PedroPereiraN/go-hexagonal/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserServiceLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := service.NewUserService(repository)

	t.Run("email_not_found", func(t *testing.T) {

		userEmail := "test@email.com"
		newPassword := "password@123"

		repository.EXPECT().FindUserByEmail(userEmail).Return(domain.UserDomain{}, errors.New("User not found"))
		token, err := service.Login(userEmail, newPassword)

		assert.EqualValues(t, "", token)

		assert.EqualError(t, err, "User not found")
	})

	t.Run("wrong_password", func(t *testing.T) {

		userEmail := "test@email.com"
		newPassword := "password@123"

		repository.EXPECT().FindUserByEmail(userEmail).Return(domain.UserDomain{Password: "@123"}, nil)
		token, err := service.Login(userEmail, newPassword)

		assert.EqualValues(t, "", token)

		assert.EqualError(t, err, "Wrong password")
	})

	t.Run("login_success", func(t *testing.T) {

		userEmail := "test@email.com"
		newPassword := "password@123"

		uDomain, err := domain.CreateUser(uuid.Nil, "", userEmail, "00000000000", newPassword, time.Time{}, time.Time{}, time.Time{})

		if err != nil {
			fmt.Println("Error while trying to create a new user for testing")
		}

		repository.EXPECT().FindUserByEmail(userEmail).Return(uDomain, nil)
		_, err = service.Login(userEmail, newPassword)

		assert.NoError(t, err)
	})
}
