package repository

import (
	"Go-TiketPemesanan/internal/domain"
	"errors"
)

type UserRepositoryInterface interface {
	UserSaver
	UserFindById
	UserUpdater
	UserDeleter
	UserLister
}

type UserSaver interface {
	UserSaver(user *domain.User) (domain.User, error)
}

type UserFindById interface {
	UserFindById(id int) (domain.User, error)
}

type UserUpdater interface {
	UpdateUser(user *domain.User) (domain.User, error)
}

type UserDeleter interface {
	DeleteUser(id int) (domain.User, error)
}

type UserLister interface {
	GetAllUser() ([]domain.User, error)
}

type UserRepository struct {
	users map[int]domain.User
}


func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{
		users: map[int]domain.User{},
	}
}

// UserFindById implements UserRepositoryInterface.
func (repo *UserRepository) UserFindById(id int) (domain.User, error) {
	usr, exist := repo.users[id]
	if !exist {
		return domain.User{}, errors.New("user not found")
	}
	return usr, nil
}

// DeleteUser implements UserRepositoryInterface.
func (repo *UserRepository) DeleteUser(id int) (domain.User, error) {
	deletedUser, exist := repo.users[id]
	if !exist {
		return domain.User{}, errors.New("user not found")
	}

	delete(repo.users, id)
	return deletedUser, nil
}

// GetAllUser implements UserRepositoryInterface.
func (repo *UserRepository) GetAllUser() ([]domain.User, error) {
	users := []domain.User{}
	for _, user := range repo.users {
		users = append(users, user)
	}
	return users, nil
}

// SaveUser implements UserRepositoryInterface.
func (repo *UserRepository) UserSaver(user *domain.User) (domain.User, error) {
	if _, exist := repo.users[user.ID]; exist {
		return *user, errors.New("user already exist")
	}

	if repo.users == nil || len(repo.users) == 0 {
		user.ID = 1
	} else {
		user.ID = repo.users[len(repo.users)].ID + 1
	}

	repo.users[user.ID] = *user
	return *user, nil
}

// UpdateUser implements UserRepositoryInterface.
func (repo *UserRepository) UpdateUser(user *domain.User) (domain.User, error) {
	if _, exist := repo.users[user.ID]; exist {
		repo.users[user.ID] = *user
		return *user, nil
	}
	return *user, errors.New("user not found")
}
