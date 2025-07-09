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

func TestUserRepository_FindUserByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)

	t.Run("user_not_found", func(t *testing.T) {

		userPhone := "00000000000"

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userPhone).
		WillReturnError(errors.New("sql: no rows in result set"))

		uDomain, err := repository.FindUserByPhone(userPhone)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("createdAt_parse_error", func(t *testing.T) {

		userPhone := "00000000000"

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userPhone).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        uuid.New(), "Test", "hashedPass", "invalid@email.com", userPhone,
        "invalid-time-format",
        nil, nil,
    ))

		uDomain, err := repository.FindUserByPhone(userPhone)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "parsing time \"invalid-time-format\" as \"2006-01-02T15:04:05\": cannot parse \"invalid-time-format\" as \"2006\"")
	})

	t.Run("updatedAt_parse_error", func(t *testing.T) {

		userPhone := "00000000000"

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userPhone).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        uuid.New(), "Test", "hashedPass", "invalid@email.com", userPhone,
        nil,
        "invalid-time-format", nil,
    ))

		uDomain, err := repository.FindUserByPhone(userPhone)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "parsing time \"invalid-time-format\" as \"2006-01-02T15:04:05\": cannot parse \"invalid-time-format\" as \"2006\"")
	})

	t.Run("deletedAt_parse_error", func(t *testing.T) {

		userPhone := "00000000000"

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userPhone).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        uuid.New(), "Test", "hashedPass", "invalid@email.com", userPhone,
        nil,
        nil, "invalid-time-format",
    ))

		uDomain, err := repository.FindUserByPhone(userPhone)

		assert.EqualValues(t, domain.UserDomain{}, uDomain)
		assert.EqualError(t, err, "parsing time \"invalid-time-format\" as \"2006-01-02T15:04:05\": cannot parse \"invalid-time-format\" as \"2006\"")
	})

	t.Run("find_user_by_phone_success", func(t *testing.T) {

		userData := domain.UserDomain{
			Id: uuid.New(),
			Name: "test name",
			Email: "test@email.com",
			Phone: "00000000000",
			Password: "hashedPass",
		}

		mock.
		ExpectQuery("SELECT (.+) FROM users").
    WithArgs(userData.Phone).
    WillReturnRows(sqlmock.NewRows([]string{
        "id", "name", "password", "email", "phone", "createdAt", "updatedAt", "deletedAt",
    }).AddRow(
        userData.Id, userData.Name,	userData.Password, userData.Email, userData.Phone,
        nil,
        nil, nil,
    ))

		uDomain, err := repository.FindUserByPhone(userData.Phone)

		assert.EqualValues(t, userData, uDomain)
		assert.NoError(t, err)
	})
}
