package usecase

import (
	"Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/repository"
)

type UserUsecaseInterface interface {
	UserSaver
	UpdateUser
	DeleteUser
	GetAllUser
}

type UserSaver interface {
	UserSaver(user domain.User) (domain.User, error)
}

type UpdateUser interface {
	UpdateUser(user domain.User) (domain.User, error)
}

type DeleteUser interface {
	DeleteUser(id int) (domain.User ,error)
}

type GetAllUser interface {
	GetAllUser() ([]domain.User, error)
}

type UserUsecase struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserUsecase(userRepo repository.UserRepositoryInterface) UserUsecase {
	return UserUsecase{
		UserRepo: userRepo,
	}
}

func (uc UserUsecase) UserSaver(user domain.User) (domain.User, error) {
	return uc.UserRepo.UserSaver(&user)
}

func (uc UserUsecase) UpdateUser(user domain.User) (userResponse domain.User, err error) {
	userResponse, err = uc.UserRepo.UpdateUser(&user)
	if err != nil {
		return userResponse, err
	}
	return userResponse, nil
}

func (uc UserUsecase) DeleteUser(id int) (user domain.User, err error) {
	user, err = uc.UserRepo.DeleteUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uc UserUsecase) GetAllUser() ([]domain.User, error) {
	users, err := uc.UserRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}
