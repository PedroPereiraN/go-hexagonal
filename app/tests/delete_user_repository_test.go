package test

import (
	"errors"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PedroPereiraN/go-hexagonal/adapter/output/repository"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	//"time"
)

func TestUserRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	repository := repository.NewUserRepository(db)

	t.Run("user_not_found", func(t *testing.T) {

		userId := uuid.New()

		mock.
		ExpectQuery("UPDATE users SET (.+) WHERE id = (.+)").
    WithArgs(userId, sqlmock.AnyArg()).
		WillReturnError(errors.New("sql: no rows in result set"))

		id, err := repository.Delete(userId)

		assert.EqualValues(t, uuid.Nil, id)
		assert.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("delete_user_success", func(t *testing.T) {

		userId := uuid.New()

		mock.
		ExpectQuery("UPDATE users SET (.+) WHERE id = (.+)").
    WithArgs(userId, sqlmock.AnyArg()).
		WillReturnRows(
        sqlmock.NewRows([]string{"id"}).AddRow(userId),
    )

		id, err := repository.Delete(userId)

		assert.EqualValues(t, userId, id)
		assert.NoError(t, err)
	})
}
