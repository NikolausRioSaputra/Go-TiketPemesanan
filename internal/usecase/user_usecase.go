package usecase

import (
	"Go-TiketPemesanan/internal/domain"
	// "Go-TiketPemesanan/internal/repository"
	"Go-TiketPemesanan/internal/repositorydb"
	"context"
)

type UserUsecaseInterface interface {
	UserSaver
	UserFindById
	UserUpdater
	UserDeleter
	GetAllUser
}

type UserSaver interface {
	UserSaver(ctx context.Context, user domain.User) (domain.User, error)
}

type UserFindById interface {
	UserFindById(ctx context.Context, id int) (domain.User, error)
}

type UserUpdater interface {
	UserUpdater(ctx context.Context, user domain.User) (domain.User, error)
}

type UserDeleter interface {
	UserDeleter(ctx context.Context, id int) error
}

type GetAllUser interface {
	GetAllUser(ctx context.Context) ([]domain.User, error)
}

type UserUsecase struct {
	UserRepo repositorydb.UserRepositoryInterface
}

func NewUserUsecase(userRepo repositorydb.UserRepositoryInterface) UserUsecase {
	return UserUsecase{
		UserRepo: userRepo,
	}
}

func (uc UserUsecase) UserSaver(ctx context.Context, user domain.User) (domain.User, error) {
	return uc.UserRepo.UserSaver(ctx, &user)
}

func (uc UserUsecase) UserFindById(ctx context.Context, id int) (domain.User, error) {
	return uc.UserRepo.UserFindById(ctx, id)
}

func (uc UserUsecase) UserUpdater(ctx context.Context, user domain.User) (userResponse domain.User, err error) {
	return uc.UserRepo.UserUpdater(ctx, &user)
}

func (uc UserUsecase) UserDeleter(ctx context.Context, id int) error {
	_, err := uc.UserFindById(ctx, id)
	if err != nil {
		return err
	}
	return uc.UserRepo.UserDeleter(ctx, id)
}

func (uc UserUsecase) GetAllUser(ctx context.Context) ([]domain.User, error) {
	return uc.UserRepo.GetAllUser(ctx)
}
