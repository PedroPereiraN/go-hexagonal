package repository

import (
	"database/sql"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
	"time"
	"fmt"
	"strings"
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type UserRepository interface {
	CreateTable() error
  Create(domain.UserDomain) (uuid.UUID, error)
	FindUserByPhone(string) (domain.UserDomain, error)
	FindUserByEmail(string) (domain.UserDomain, error)
	List(uuid.UUID) (domain.UserDomain, error)
	ListAll() ([]domain.UserDomain, error)
	Delete(uuid.UUID) (uuid.UUID, error)
	Update(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
	UpdatePassword(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
}

type userRepository struct {
	db *sql.DB
}

func (repository *userRepository) CreateTable() error {

query := `CREATE TABLE IF NOT EXISTS users (
    id uuid UNIQUE PRIMARY KEY,
    name varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
		phone varchar(11) NOT NULL,
    createdAt timestamp DEFAULT NOW(),
    updatedAt timestamp,
    deletedAt timestamp
  )`

  _, err := repository.db.Exec(query)

  if err != nil {
    return err
  }

  return nil
}

func (repository *userRepository) Create(dto domain.UserDomain) (uuid.UUID, error) {
	uDomain, err := domain.CreateUser(
		dto.Id,
		dto.Name,
		dto.Email,
		dto.Phone,
		dto.Password,
		dto.CreatedAt,
		dto.UpdatedAt,
		dto.DeletedAt,
	)

	if err != nil {
    return uuid.Nil, err
  }

	query := `INSERT INTO users (
    id,
    name,
    password,
    email,
		phone,
		createdAt
  ) VALUES (
    $1, $2, $3, $4, $5, $6
  ) RETURNING id`

  var pk uuid.UUID

  err = repository.db.QueryRow(query, uDomain.Id, uDomain.Name, uDomain.Password, uDomain.Email, uDomain.Phone, uDomain.CreatedAt).Scan(&pk)

  if err != nil {
    return uuid.Nil, err
  }

  return pk, nil
}

func (repository *userRepository) FindUserByPhone(phone string) (domain.UserDomain,error) {
	uDomain := domain.UserDomain{}
	var createdAt sql.NullString
	var updatedAt sql.NullString
	var deletedAt sql.NullString

  query := `SELECT id, name, password, email, phone, createdAt, updatedAt, deletedAt FROM users WHERE phone = $1 AND deletedAt IS NULL`

  err := repository.db.QueryRow(query, phone).Scan(&uDomain.Id, &uDomain.Name, &uDomain.Password, &uDomain.Email, &uDomain.Phone, &createdAt, &updatedAt, &deletedAt)

	if err != nil {
    return domain.UserDomain{}, err
  }

	// to use time.Parse is necessary to pass the layout first and the value second
	// the layout is exclusively made with this date
	// this date is used because every field has a exclusive value

	timeParserLayout := "2006-01-02T15:04:05"

	if createdAt.Valid {
		parsedCreatedAt, err := time.Parse(timeParserLayout, createdAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.CreatedAt = parsedCreatedAt
	}

	if updatedAt.Valid {
		parsedUpdatedAt, err := time.Parse(timeParserLayout, updatedAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.UpdatedAt = parsedUpdatedAt
	}

	if deletedAt.Valid {
		parsedDeletedAt, err := time.Parse(timeParserLayout, deletedAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.DeletedAt = parsedDeletedAt
	}

  return uDomain, nil
}

func (repository *userRepository) FindUserByEmail(email string) (domain.UserDomain, error) {
	uDomain := domain.UserDomain{}
	var createdAt sql.NullString
	var updatedAt sql.NullString
	var deletedAt sql.NullString

  query := `SELECT id, name, password, email, phone, createdAt, updatedAt, deletedAt FROM users WHERE email = $1 AND deletedAt IS NULL`

  err := repository.db.QueryRow(query, email).Scan(&uDomain.Id, &uDomain.Name, &uDomain.Password, &uDomain.Email, &uDomain.Phone, &createdAt, &updatedAt, &deletedAt)

	if err != nil {
    return domain.UserDomain{}, err
  }

	// to use time.Parse is necessary to pass the layout first and the value second
	// the layout is exclusively made with this date
	// this date is used because every field has a exclusive value

	timeParserLayout := "2006-01-02T15:04:05"

	if createdAt.Valid {
		parsedCreatedAt, err := time.Parse(timeParserLayout, createdAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.CreatedAt = parsedCreatedAt
	}

	if updatedAt.Valid {
		parsedUpdatedAt, err := time.Parse(timeParserLayout, updatedAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.UpdatedAt = parsedUpdatedAt
	}

	if deletedAt.Valid {
		parsedDeletedAt, err := time.Parse(timeParserLayout, deletedAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.DeletedAt = parsedDeletedAt
	}

  return uDomain, nil
}

func (repository *userRepository) List(id uuid.UUID) (domain.UserDomain, error) {
	uDomain := domain.UserDomain{}
	var createdAt sql.NullString
	var updatedAt sql.NullString
	var deletedAt sql.NullString

  query := `SELECT id, name, password, email, phone, createdAt, updatedAt, deletedAt FROM users WHERE id = $1 AND deletedAt IS NULL`

  err := repository.db.QueryRow(query, id).Scan(&uDomain.Id, &uDomain.Name, &uDomain.Password, &uDomain.Email, &uDomain.Phone, &createdAt, &updatedAt, &deletedAt)

	if err != nil {
    return domain.UserDomain{}, err
  }

	// to use time.Parse is necessary to pass the layout first and the value second
	// the layout is exclusively made with this date
	// this date is used because every field has a exclusive value

	timeParserLayout := "2006-01-02T15:04:05"

	if createdAt.Valid {
		parsedCreatedAt, err := time.Parse(timeParserLayout, createdAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.CreatedAt = parsedCreatedAt
	}

	if updatedAt.Valid {
		parsedUpdatedAt, err := time.Parse(timeParserLayout, updatedAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.UpdatedAt = parsedUpdatedAt
	}

	if deletedAt.Valid {
		parsedDeletedAt, err := time.Parse(timeParserLayout, deletedAt.String[0:19])

		if err != nil {
			return domain.UserDomain{}, err
		}

		uDomain.DeletedAt = parsedDeletedAt
	}

  return uDomain, nil
}

func (repository *userRepository) ListAll() ([]domain.UserDomain, error) {
	var users []domain.UserDomain

  query := `SELECT id, name, password, email, phone, createdAt, updatedAt, deletedAt FROM users WHERE deletedAt IS NULL`

  rows, err := repository.db.Query(query)

	if err != nil {
    return []domain.UserDomain{}, err
  }

	defer rows.Close()

	for rows.Next() {
		uDomain := domain.UserDomain{}
		var createdAt sql.NullString
		var updatedAt sql.NullString
		var deletedAt sql.NullString

		rows.Scan(&uDomain.Id, &uDomain.Name, &uDomain.Password, &uDomain.Email, &uDomain.Phone, &createdAt, &updatedAt, &deletedAt)

		// to use time.Parse is necessary to pass the layout first and the value second
		// the layout is exclusively made with this date
		// this date is used because every field has a exclusive value
		timeParserLayout := "2006-01-02T15:04:05"

		if createdAt.Valid {
			parsedCreatedAt, err := time.Parse(timeParserLayout, createdAt.String[0:19])

			if err != nil {
				return []domain.UserDomain{}, err
			}

			uDomain.CreatedAt = parsedCreatedAt
		}

		if updatedAt.Valid {
			parsedUpdatedAt, err := time.Parse(timeParserLayout, updatedAt.String[0:19])

			if err != nil {
				return []domain.UserDomain{}, err
			}

			uDomain.UpdatedAt = parsedUpdatedAt
		}

		if deletedAt.Valid {
			parsedDeletedAt, err := time.Parse(timeParserLayout, deletedAt.String[0:19])

			if err != nil {
				return []domain.UserDomain{}, err
			}

			uDomain.DeletedAt = parsedDeletedAt
		}

		users = append(users, uDomain)
	}


  return users, nil
}


func (repository *userRepository) Delete(id uuid.UUID) (uuid.UUID, error) {
	var pk uuid.UUID

  //query := `DELETE FROM users WHERE id = $1 RETURNING id`
	query := `UPDATE users SET deletedAt = $2 WHERE id = $1 RETURNING id`

  err := repository.db.QueryRow(query, id, time.Now()).Scan(&pk)

  if err != nil {
    return uuid.Nil, err
  }

  return pk, nil
}

func (repository *userRepository) Update(id uuid.UUID, dto domain.UserDomain) (uuid.UUID, error) {
	uDomain, err := domain.CreateUser(
		dto.Id,
		dto.Name,
		dto.Email,
		dto.Phone,
		dto.Password,
		dto.CreatedAt,
		dto.UpdatedAt,
		dto.DeletedAt,
	)

	if err != nil {
		return uuid.Nil, err
	}

	setClauses := []string{}
	args := []any{id}
	argIndex := 2

	if uDomain.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, uDomain.Name)
		argIndex++
	}

	if uDomain.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIndex)) // 3
		args = append(args, uDomain.Email)
		argIndex++
	}

	if uDomain.Phone != "" {
		setClauses = append(setClauses, fmt.Sprintf("phone = $%d", argIndex))
		args = append(args, uDomain.Phone)
		argIndex++
	}

	setClauses = append(setClauses, fmt.Sprintf("updatedAt = $%d", argIndex))
	args = append(args, time.Now())

	query := fmt.Sprintf(`UPDATE users SET %s WHERE id = $1 RETURNING id`, strings.Join(setClauses, ", "))

	var pk uuid.UUID

	err = repository.db.QueryRow(query, args...).Scan(&pk)

	if err != nil {
		return uuid.Nil, err
	}

	return pk, nil
}

func (repository *userRepository) UpdatePassword(id uuid.UUID, dto domain.UserDomain) (uuid.UUID, error) {
	uDomain, err := domain.CreateUser(
		dto.Id,
		dto.Name,
		dto.Email,
		dto.Phone,
		dto.Password,
		dto.CreatedAt,
		dto.UpdatedAt,
		dto.DeletedAt,
	)

	if err != nil {
		return uuid.Nil, err
	}

	query := `UPDATE users SET password = $2 WHERE id = $1 RETURNING id`

	var pk uuid.UUID

	err = repository.db.QueryRow(query, id, uDomain.Password).Scan(&pk)

	if err != nil {
		return uuid.Nil, err
	}

	return pk, nil
}
