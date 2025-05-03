package database

import (
	"database/sql"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/google/uuid"
)

func CreateUserTable(db *sql.DB) error {
  query := `CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    name varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    createdAt timestamp DEFAULT NOW(),
    updatedAt timestamp,
    deletedAt timestamp
  )`

  _, err := db.Exec(query)

  if err != nil {
    return err
  }

  return nil
}

func SaveUser(db *sql.DB, dto domain.UserDomain) (uuid.UUID, error) {
  query := `INSERT INTO users (
    id,
    name,
    password,
    email
  ) VALUES (
    $1, $2, $3, $4
  ) RETURNING id`

  var pk uuid.UUID

  err := db.QueryRow(query, dto.Id, dto.Name, dto.Password, dto.Email).Scan(&pk)

  if err != nil {
    return uuid.New(), err
  }

  return pk, nil
}

func ListUser(db *sql.DB, id uuid.UUID) (domain.UserDomain, error) {

  uDomain := domain.UserDomain{}

  query := `SELECT id, name, password, email FROM users WHERE id = $1`

  err := db.QueryRow(query, id).Scan(&uDomain.Id, &uDomain.Name, &uDomain.Password, &uDomain.Email)

  if err != nil {
    return domain.UserDomain{}, err
  }

  return uDomain, nil
}

func ListAllUser(db *sql.DB) ([]domain.UserDomain, error) {

  var users []domain.UserDomain

  query := `SELECT id, name, password, email FROM users`

  rows, err := db.Query(query)

  if err != nil {
    return users, err
  }

  defer rows.Close()

  for rows.Next() {
    user := domain.UserDomain{}
    err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)

    if err != nil {
      return []domain.UserDomain{}, err
    }

    users = append(users, user)
  }

  return users, nil
}

func UpdateUser(db *sql.DB, id uuid.UUID, dto domain.UserDomain) (uuid.UUID, error) {

  user, err := ListUser(db, id)

  newUserData := domain.UserDomain{}

  if dto.Name == "" {
    newUserData.Name = user.Name
  } else {
    newUserData.Name = dto.Name
  }

  if dto.Email == "" {
    newUserData.Email = user.Email
  } else {
    newUserData.Email = dto.Email
  }

  if err != nil {
    return uuid.New(), err
  }

  query := `UPDATE users SET name = $2, email = $3 WHERE id = $1 RETURNING id`

  var pk uuid.UUID

  err = db.QueryRow(query, id, newUserData.Name, newUserData.Email).Scan(&pk)

  if err != nil {
    return uuid.New(), err
  }

  return pk, nil
}

func DeleteUser(db *sql.DB, id uuid.UUID) (uuid.UUID, error) {

  var pk uuid.UUID

  query := `DELETE FROM users WHERE id = $1 RETURNING id`

  err := db.QueryRow(query, id).Scan(&pk)

  if err != nil {
    return uuid.New(), err
  }

  return pk, nil
}

func UpdateUserPassword(db *sql.DB, id uuid.UUID, password string) (uuid.UUID, error) {

  query := `UPDATE users SET password = $2 WHERE id = $1 RETURNING id`

  var pk uuid.UUID

  err := db.QueryRow(query, id, password).Scan(&pk)

  if err != nil {
    return uuid.New(), err
  }

  return pk, nil
}
