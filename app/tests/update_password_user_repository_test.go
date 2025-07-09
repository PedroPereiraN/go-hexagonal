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

func TestUserRepository_UpdatePassword(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	repository := repository.NewUserRepository(db)

	t.Run("user_not_found", func(t *testing.T) {

		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "INVALID_NAME",
			Email: "INVALID_EMAIL",
			Phone: "INVALID_PHONE000",
		}

		mock.ExpectQuery("UPDATE users SET (.+) WHERE id = (.+)").
    WithArgs(
			uDomain.Id,
			sqlmock.AnyArg(),
		).
		WillReturnError(errors.New("sql: no rows in result set"))

		id, err := repository.UpdatePassword(uDomain.Id, uDomain)

		assert.EqualValues(t, uuid.Nil, id)
		assert.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("invalid_fields", func(t *testing.T) {

		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "INVALID_NAME",
			Email: "INVALID_EMAIL",
			Phone: "INVALID_PHONE000",
		}

		mock.ExpectQuery("UPDATE users SET (.+) WHERE id = (.+)").
    WithArgs(
			uDomain.Id,
			sqlmock.AnyArg(),
		).
		WillReturnError(errors.New("database update failed"))

		id, err := repository.UpdatePassword(uDomain.Id, uDomain)

		assert.EqualValues(t, uuid.Nil, id)
		assert.EqualError(t, err, "database update failed")
	})

	t.Run("update_user_password_success", func(t *testing.T) {
		uDomain := domain.UserDomain{
			Id:    uuid.New(),
			Name:  "Test name",
			Email: "test@email.com",
			Phone: "00000000000",
		}

		mock.ExpectQuery("UPDATE users SET (.+) WHERE id = (.+)").
    WithArgs(
			uDomain.Id,
			sqlmock.AnyArg(),
		).
    WillReturnRows(
        sqlmock.NewRows([]string{"id"}).AddRow(uDomain.Id),
    )

		id, err := repository.UpdatePassword(uDomain.Id, uDomain)

		assert.EqualValues(t, uDomain.Id, id)
		assert.NoError(t, err)
	})
}
