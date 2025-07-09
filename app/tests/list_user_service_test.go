package test

import (
	"errors"
	"testing"
	"time"

	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/services"
	"github.com/PedroPereiraN/go-hexagonal/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := service.NewUserService(repository)

	t.Run("user_not_found", func(t *testing.T) {

		userId := uuid.New()

		repository.EXPECT().List(userId).Return(domain.UserDomain{}, errors.New("User not found"))
		uDomain, err := service.List(userId)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)

		assert.EqualError(t, err, "User not found")
	})

	t.Run("user_found", func(t *testing.T) {

		foundUser := domain.UserDomain{
			Id: uuid.New(),
			Name: "Found User",
			Email: "email@test.com",
			Phone: "00000000000",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		}

		repository.EXPECT().List(foundUser.Id).Return(foundUser, nil)
		uDomain, err := service.List(foundUser.Id)

		assert.EqualValues(t, foundUser, uDomain)
		assert.NoError(t, err)
	})
}
