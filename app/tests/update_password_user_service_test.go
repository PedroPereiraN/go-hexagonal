package test

import (
	"errors"
	"testing"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/services"
	"github.com/PedroPereiraN/go-hexagonal/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserService_UpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := service.NewUserService(repository)

	t.Run("user_not_found", func(t *testing.T) {

		userId := uuid.New()
		newPassword := "password@123"

		repository.EXPECT().List(userId).Return(domain.UserDomain{}, errors.New("User not found"))
		id, err := service.UpdatePassword(userId, newPassword)

		assert.EqualValues(t, uuid.Nil, id)

		assert.EqualError(t, err, "User not found")
	})

	t.Run("repository_error", func(t *testing.T) {

		userId := uuid.New()
		newPassword := "password@123"

		repository.EXPECT().List(userId).Return(domain.UserDomain{}, nil)
		repository.EXPECT().UpdatePassword(userId, gomock.Any()).Return(uuid.Nil, errors.New("repository error"))

		id, err := service.UpdatePassword(userId, newPassword)

		assert.EqualValues(t, uuid.Nil, id)

		assert.EqualError(t, err, "repository error")
	})

	t.Run("password_update_successfuly", func(t *testing.T) {

		userId := uuid.New()
		newPassword := "password@123"

		repository.EXPECT().List(userId).Return(domain.UserDomain{}, nil)
		repository.EXPECT().UpdatePassword(userId, gomock.Any()).Return(userId, nil)

		id, err := service.UpdatePassword(userId, newPassword)

		assert.EqualValues(t, userId, id)

		assert.NoError(t, err)
	})
}
