package test

import (
	"errors"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PedroPereiraN/go-hexagonal/adapter/output/repository"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)

	t.Run("invalid_fields", func(t *testing.T) {

		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "INVALID_NAME",
			Email: "INVALID_EMAIL",
			Phone: "INVALID_PHONE000",
			Password: "INVALID_PASSWORD",
		}

		mock.ExpectQuery("INSERT INTO users (.+)").
    WithArgs(
			uDomain.Id,
    	uDomain.Name,
    	sqlmock.AnyArg(),
    	uDomain.Email,
			uDomain.Phone,
			sqlmock.AnyArg(),
		).
		WillReturnError(errors.New("database insert failed"))

		id, err := repository.Create(uDomain)

		assert.EqualValues(t, uuid.Nil, id)
		assert.EqualError(t, err, "database insert failed")
	})

	t.Run("create_user_success", func(t *testing.T) {
		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "Test name",
			Email: "test@email.com",
			Phone: "00000000000",
			Password: "password@123",
		}

		mock.ExpectQuery("INSERT INTO users (.+)").
    WithArgs(
			uDomain.Id,
    	uDomain.Name,
    	sqlmock.AnyArg(),
    	uDomain.Email,
			uDomain.Phone,
			sqlmock.AnyArg(),
		).
    WillReturnRows(
        sqlmock.NewRows([]string{"id"}).AddRow(uDomain.Id),
    )

		id, err := repository.Create(uDomain)

		assert.EqualValues(t, uDomain.Id, id)
		assert.NoError(t, err)
	})
}
