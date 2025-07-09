package test

import (
	"errors"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PedroPereiraN/go-hexagonal/adapter/output/repository"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestUserRepository_List(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	repository := repository.NewUserRepository(db)

	t.Run("user_not_found", func(t *testing.T) {

		userId := uuid.New()

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userId).
		WillReturnError(errors.New("sql: no rows in result set"))

		uDomain, err := repository.List(userId)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("createdAt_parse_error", func(t *testing.T) {

		userId := uuid.New()

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userId).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        userId, "Test", "hashedPass", "invalid@email.com", "00000000000",
        "invalid-time-format",
        nil, nil,
    ))

		uDomain, err := repository.List(userId)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "parsing time \"invalid-time-format\" as \"2006-01-02T15:04:05\": cannot parse \"invalid-time-format\" as \"2006\"")
	})

	t.Run("updatedAt_parse_error", func(t *testing.T) {

		userId := uuid.New()

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userId).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        userId, "Test", "hashedPass", "invalid@email.com", "00000000000",
        nil,
        "invalid-time-format", nil,
    ))

		uDomain, err := repository.List(userId)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "parsing time \"invalid-time-format\" as \"2006-01-02T15:04:05\": cannot parse \"invalid-time-format\" as \"2006\"")
	})

	t.Run("deletedAt_parse_error", func(t *testing.T) {

		userId := uuid.New()

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userId).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        userId, "Test", "hashedPass", "invalid@email.com", "00000000000",
        nil,
        nil, "invalid-time-format",
    ))

		uDomain, err := repository.List(userId)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "parsing time \"invalid-time-format\" as \"2006-01-02T15:04:05\": cannot parse \"invalid-time-format\" as \"2006\"")
	})

	t.Run("list_user_success", func(t *testing.T) {

		userData := domain.UserDomain{
			Id: uuid.New(),
			Name: "test name",
			Email: "test@email.com",
			Phone: "00000000000",
			Password: "hashedPass",
		}

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userData.Id).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        userData.Id, userData.Name,	userData.Password, userData.Email, userData.Phone,
        nil,
        nil, nil,
    ))

		uDomain, err := repository.List(userData.Id)

		assert.EqualValues(t, userData, uDomain)
		assert.NoError(t, err)
	})
}
