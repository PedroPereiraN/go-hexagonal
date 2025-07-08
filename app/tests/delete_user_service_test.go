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

func TestUserServiceDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := service.NewUserService(repository)

	t.Run("user_not_found", func(t *testing.T) {

		userId := uuid.New()

		repository.EXPECT().List(userId).Return(domain.UserDomain{}, errors.New("User not found"))

		id, err := service.Delete(userId)

		assert.EqualValues(t, uuid.Nil, id)

		assert.EqualError(t, err, "User not found")
	})

	t.Run("repository_error", func(t *testing.T) {

		userId := uuid.New()

		foundUser := domain.UserDomain{
			Id: uuid.New(),
			Name: "Found User",
			Email: "email@test.com",
			Phone: "00000000000",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		}

		repository.EXPECT().List(userId).Return(foundUser, nil)
		repository.EXPECT().Delete(userId).Return(uuid.Nil, errors.New("Repository error"))

		id, err := service.Delete(userId)

		assert.EqualValues(t, uuid.Nil, id)

		assert.EqualError(t, err, "Repository error")
	})

	t.Run("user_deleted", func(t *testing.T) {

		userId := uuid.New()

		foundUser := domain.UserDomain{
			Id: uuid.New(),
			Name: "Found User",
			Email: "email@test.com",
			Phone: "00000000000",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		}

		repository.EXPECT().List(userId).Return(foundUser, nil)
		repository.EXPECT().Delete(userId).Return(userId, nil)
		id, err := service.Delete(userId)

		assert.EqualValues(t, userId, id)
		assert.NoError(t, err)
	})
}
