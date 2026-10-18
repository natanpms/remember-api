package repository

import (
	"database/sql"
	"fmt"

	"github.com/natanpms/remember-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) CreateUser(user model.User) (model.User, error) {
	const query = `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	err := ur.connection.QueryRow(query, user.Name, user.Email, user.PasswordHash).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (model.User, error) {
	const query = `
		SELECT id, name, email, password_hash, created_at
		FROM users
		WHERE email = $1`

	var user model.User

	err := ur.connection.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}

	return user, nil
}
