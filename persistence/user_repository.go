package persistence

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
	"todo-app--go-gin/domain"
)

type IUserRepository interface {
	GetAllUsers() ([]domain.User, error)
	AddUser(user domain.User) (domain.User, error)
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

func (userRepository UserRepository) AddUser(user domain.User) (domain.User, error) {
	ctx := context.Background()
	insertSql := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := userRepository.dbPool.Exec(ctx, insertSql, user.Username, user.Email, user.Password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
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
