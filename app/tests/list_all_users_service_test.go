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

func TestUserService_ListAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := service.NewUserService(repository)

	t.Run("repository_error", func(t *testing.T) {

		repository.EXPECT().ListAll().Return([]domain.UserDomain{}, errors.New("Service error"))
		uDomain, err := service.ListAll()

		assert.EqualValues(t, []domain.UserDomain{}, uDomain)

		assert.EqualError(t, err, "Service error")
	})

	t.Run("users_found", func(t *testing.T) {

		var foundUsers []domain.UserDomain

		example1 := domain.UserDomain{
			Id: uuid.New(),
			Name: "Found User",
			Email: "email@test.com",
			Phone: "00000000000",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		}

		foundUsers = append(foundUsers, example1)

		repository.EXPECT().ListAll().Return(foundUsers, nil)
		users, err := service.ListAll()

		assert.EqualValues(t, foundUsers, users)
		assert.NoError(t, err)
	})
}
