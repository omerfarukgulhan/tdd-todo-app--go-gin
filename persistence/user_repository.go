package persistence

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"todo-app--go-gin/domain"
)

type IUserRepository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(userId int) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	AddUser(user domain.User) (domain.User, error)
	UpdateUser(userId int, user domain.User) (domain.User, error)
	DeleteUser(userId int) error
}

type UserRepository struct {
	dbPool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) IUserRepository {
	return &UserRepository{dbPool: dbPool}
}

func (userRepository UserRepository) GetAllUsers() ([]domain.User, error) {
	ctx := context.Background()
	queryRow, err := userRepository.dbPool.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return []domain.User{}, err
	}

	return extractUsersFromRows(queryRow), nil
}

func (userRepository UserRepository) GetUserById(userId int) (domain.User, error) {
	ctx := context.Background()
	var user domain.User
	getByIdSql := `SELECT * FROM users WHERE id = $1`
	queryRow := userRepository.dbPool.QueryRow(ctx, getByIdSql, userId)
	scanErr := queryRow.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return domain.User{}, errors.New(fmt.Sprintf("User with id %d not found", userId))
		}
		return domain.User{}, errors.New(fmt.Sprintf("Error while getting user with id %d: %v", userId, scanErr))
	}

	return user, nil
}

func (userRepository UserRepository) GetUserByEmail(email string) (domain.User, error) {
	ctx := context.Background()
	var user domain.User
	getByEmailSql := `SELECT * FROM users WHERE email = $1`
	queryRow := userRepository.dbPool.QueryRow(ctx, getByEmailSql, email)
	scanErr := queryRow.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return domain.User{}, errors.New(fmt.Sprintf("User with email %s not found", email))
		}
		return domain.User{}, errors.New(fmt.Sprintf("Error while getting user with email %s: %v", email, scanErr))
	}

	return user, nil
}

func (userRepository UserRepository) AddUser(user domain.User) (domain.User, error) {
	ctx := context.Background()
	insertSql := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id int
	queryRow := userRepository.dbPool.QueryRow(ctx, insertSql, user.Username, user.Email, user.Password)
	scanErr := queryRow.Scan(&id)
	if scanErr != nil {
		return domain.User{}, scanErr
	}

	user.Id = id

	return user, nil
}

func (userRepository UserRepository) UpdateUser(userId int, user domain.User) (domain.User, error) {
	ctx := context.Background()
	updateUserSql := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4 RETURNING id, username, email, password;`
	queryRow := userRepository.dbPool.QueryRow(ctx, updateUserSql, user.Username, user.Email, user.Password, userId)
	scanErr := queryRow.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return domain.User{}, errors.New(fmt.Sprintf("User with id %d not found", userId))
		}
		return domain.User{}, errors.New(fmt.Sprintf("Failed to update user: %v", scanErr))
	}

	return user, nil
}

func (userRepository UserRepository) DeleteUser(userId int) error {
	ctx := context.Background()
	_, getErr := userRepository.GetUserById(userId)
	if getErr != nil {
		return fmt.Errorf("user with id %d not found: %w", userId, getErr)
	}

	// Prepare the delete query
	deleteSql := `DELETE FROM users WHERE id = $1`
	_, err := userRepository.dbPool.Exec(ctx, deleteSql, userId)
	if err != nil {
		return fmt.Errorf("error while deleting user with id %d: %w", userId, err)
	}

	return nil
}

func extractUsersFromRows(queryRow pgx.Rows) []domain.User {
	var users = []domain.User{}
	for queryRow.Next() {
		var user domain.User
		err := queryRow.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	if err := queryRow.Err(); err != nil {
		return users
	}

	return users
}
