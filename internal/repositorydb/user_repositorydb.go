package repositorydb

import (
	"Go-TiketPemesanan/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"sync"
)

type UserRepositoryInterface interface {
	UserSaver
	UserLister
	UserFindById
	UserUpdater
	UserDeleter
	UpdateBalance
}

type UserSaver interface {
	UserSaver(ctx context.Context, user *domain.User) (domain.User, error)
}

type UserLister interface {
	GetAllUser(ctx context.Context) ([]domain.User, error)
}

type UserFindById interface {
	UserFindById(ctx context.Context, id int) (domain.User, error)
}

type UserUpdater interface {
	UserUpdater(ctx context.Context, user *domain.User) (domain.User, error)
}

type UserDeleter interface {
	UserDeleter(ctx context.Context, id int) error
}

type UpdateBalance interface {
	UpdateBalance(ctx context.Context, id int, newBalance float64) (domain.User, error)
}

type UserRepository struct {
	DB *sql.DB
	mu sync.Mutex
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
		mu: sync.Mutex{},
	}
}

// GetAllUser implements UserRepositoryInterface.
func (repo *UserRepository) GetAllUser(ctx context.Context) ([]domain.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var users []domain.User
	query := "SELECT id, name, address, balance FROM users"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Balance)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UserSaver implements UserRepositoryInterface.
func (repo *UserRepository) UserSaver(ctx context.Context, user *domain.User) (domain.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	query := `INSERT INTO users( name, address, balance) VALUES($1, $2, $3) RETURNING id`
	err := repo.DB.QueryRowContext(ctx, query, user.Name, user.Address, user.Balance).Scan(&user.ID)
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}

func (repo *UserRepository) UserFindById(ctx context.Context, id int) (domain.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var user domain.User
	query := `SELECT id, name, address, balance FROM users WHERE id = $1`
	err := repo.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Address, &user.Balance)
	if err != nil {
		return domain.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (repo *UserRepository) UserUpdater(ctx context.Context, user *domain.User) (domain.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	query := `UPDATE users SET name = $1, address = $2, balance = $3 WHERE id = $4 RETURNING id`
	err := repo.DB.QueryRowContext(ctx, query, user.Name, user.Address, user.Balance, user.ID).Scan(&user.ID)
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}

func (repo *UserRepository) UserDeleter(ctx context.Context, id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	query := `DELETE FROM users WHERE id = $1`
	_, err := repo.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) UpdateBalance(ctx context.Context, id int, newBalance float64) (domain.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var user domain.User
	query := "UPDATE users SET balance = $1 where id = $2 RETURNING id, name, address, balance"
	err := repo.DB.QueryRowContext(ctx, query, newBalance, id).Scan(&user.ID, &user.Name, &user.Address, &user.Balance)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}
