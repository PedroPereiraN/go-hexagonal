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

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := service.NewUserService(repository)

	t.Run("user_phone_already_registered", func(t *testing.T) {
		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "Test name",
			Email: "test@email.com",
			Phone: "00000000000",
			Password: "password@123",
		}

		repository.EXPECT().FindUserByPhone(uDomain.Phone).Return(uDomain, nil)
		id, err := service.Create(uDomain)

		assert.EqualValues(t, uuid.Nil, id)

		assert.EqualError(t, err, "Phone is already registered")
	})

	t.Run("user_email_already_registered", func(t *testing.T) {
		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "Test name",
			Email: "test@email.com",
			Phone: "00000000000",
			Password: "password@123",
		}

		repository.EXPECT().FindUserByPhone(uDomain.Phone).Return(domain.UserDomain{}, errors.New("sql: no rows in result set"))

		repository.EXPECT().FindUserByEmail(uDomain.Email).Return(uDomain, nil)

		id, err := service.Create(uDomain)

		assert.EqualValues(t, uuid.Nil, id)
		assert.EqualError(t, err, "Email is already registered")
	})

	t.Run("repository_error", func(t *testing.T) {
		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "Test name",
			Email: "test@email.com",
			Phone: "00000000000",
			Password: "password@123",
		}

		repository.EXPECT().FindUserByPhone(uDomain.Phone).Return(domain.UserDomain{}, errors.New("sql: no rows in result set"))

		repository.EXPECT().FindUserByEmail(uDomain.Email).Return(domain.UserDomain{}, errors.New("sql: no rows in result set"))

		repository.EXPECT().Create(gomock.Any()).Return(uuid.Nil, errors.New("repository error"))

		id, err := service.Create(uDomain)

		assert.EqualValues(t, uuid.Nil, id)
		assert.EqualError(t, err, "repository error")
	})

	t.Run("create_user_success", func(t *testing.T) {
		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "Test name",
			Email: "test@email.com",
			Phone: "00000000000",
			Password: "password@123",
		}

		repository.EXPECT().FindUserByPhone(uDomain.Phone).Return(domain.UserDomain{}, errors.New("sql: no rows in result set"))

		repository.EXPECT().FindUserByEmail(uDomain.Email).Return(domain.UserDomain{}, errors.New("sql: no rows in result set"))

		repository.EXPECT().Create(gomock.Any()).Return(uDomain.Id, nil)

		id, err := service.Create(uDomain)

		assert.EqualValues(t, uDomain.Id, id)
		assert.NoError(t, err)
	})
}
