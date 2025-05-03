package repository

import (
	"database/sql"
	"github.com/PedroPereiraN/go-hexagonal/adapter/output/repository"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)
var connStr = "postgres://postgres:postgres@postgres:5432/db?sslmode=disable"

func NewUserRepository() UserRepository {
  return &userRepository{}
}

type UserRepository interface {
  Create(domain.UserDomain) (uuid.UUID, error)
  List(uuid.UUID) (domain.UserDomain, error)
  ListAll() ([]domain.UserDomain, error)
  Update(uuid.UUID, domain.UserDomain) (uuid.UUID, error)
  Delete(uuid.UUID) (uuid.UUID, error)
  UpdatePassword(uuid.UUID, string) (uuid.UUID, error)
}

type userRepository struct {}

func (repository *userRepository) Create(dto domain.UserDomain) (uuid.UUID, error) {

  db, err := sql.Open("postgres", connStr)

  if err != nil {
    return uuid.New(), err
  }

  if err = db.Ping(); err != nil {
    return uuid.New(), err
  }

  err = database.CreateUserTable(db)

  if err != nil {
    return uuid.New(), err
  }

  userId, err := database.SaveUser(db, dto)

  if err != nil {
    return uuid.New(), err
  }

  defer db.Close()

  return userId, nil
}

func (repository *userRepository) List(id uuid.UUID) (domain.UserDomain, error) {

  db, err := sql.Open("postgres", connStr)

  if err != nil {
    return domain.UserDomain{}, err
  }

  if err = db.Ping(); err != nil {
    return domain.UserDomain{}, err
  }

  user, err := database.ListUser(db, id)

  if err != nil {
    return domain.UserDomain{}, err
  }

  defer db.Close()

  return user, nil
}

func (repository *userRepository) ListAll() ([]domain.UserDomain, error) {

  db, err := sql.Open("postgres", connStr)

  if err != nil {
    return []domain.UserDomain{}, err
  }

  if err = db.Ping(); err != nil {
    return []domain.UserDomain{}, err
  }

  users, err := database.ListAllUser(db)

  if err != nil {
    return []domain.UserDomain{}, err
  }

  defer db.Close()

  return users, nil
}

func (repository *userRepository) Update(id uuid.UUID, dto domain.UserDomain) (uuid.UUID, error) {
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    return uuid.New(), err
  }

  if err = db.Ping(); err != nil {
    return uuid.New(), err
  }

  userId, err := database.UpdateUser(db, id, dto)

  if err != nil {
    return uuid.New(), err
  }

  defer db.Close()

  return userId, nil
}

func (repository *userRepository) Delete(id uuid.UUID) (uuid.UUID, error) {

  db, err := sql.Open("postgres", connStr)

  if err != nil {
    return uuid.New(), err
  }

  if err = db.Ping(); err != nil {
    return uuid.New(), err
  }

  userId, err := database.DeleteUser(db, id)

  if err != nil {
    return uuid.New(), err
  }

  defer db.Close()

  return userId, nil
}

func (repository *userRepository) UpdatePassword(id uuid.UUID, password string) (uuid.UUID, error) {
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    return uuid.New(), err
  }

  if err = db.Ping(); err != nil {
    return uuid.New(), err
  }

  userId, err := database.UpdateUserPassword(db, id, password)

  if err != nil {
    return uuid.New(), err
  }

  defer db.Close()

  return userId, nil
}
