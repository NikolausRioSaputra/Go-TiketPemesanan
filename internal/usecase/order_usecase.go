package usecase

import (
	"Go-TiketPemesanan/internal/domain"
	// "Go-TiketPemesanan/internal/repository"
	"Go-TiketPemesanan/internal/repositorydb"
	"context"
	"errors"
)

type OrderUsecaseInterface interface {
	CreateOrder
	ListOrder
}

type CreateOrder interface {
	CreateOrder(ctx context.Context, userId, eventId int, tiketType string, quantity int) (domain.Order, error)
}

type ListOrder interface {
	ListOrder(ctx context.Context) ([]domain.Order, error)
}

type OrderUsecase struct {
	OrderRepo repositorydb.OrderRepositoryinterface
	UserRepo  repositorydb.UserRepositoryInterface
	EventRepo repositorydb.EventRepositoryInterface
}

func NewOrderUsecase(orderRepo repositorydb.OrderRepositoryinterface, userRepo repositorydb.UserRepositoryInterface, eventRepo repositorydb.EventRepositoryInterface) OrderUsecase {
	return OrderUsecase{
		OrderRepo: orderRepo,
		UserRepo:  userRepo,
		EventRepo: eventRepo,
	}
}

// CreateOrder implements OrderUsecaseInterface.
func (uc OrderUsecase) CreateOrder(ctx context.Context, userId int, eventId int, tiketType string, quantity int) (domain.Order, error) {
	user, err := uc.UserRepo.UserFindById(ctx, userId)
	if err != nil {
		return domain.Order{}, err
	}

	event, err := uc.EventRepo.GetEventById(ctx, eventId)
	if err != nil {
		return domain.Order{}, err
	}

	tiketIndex := -1
	for i, tiket := range event.Tiket {
		if tiket.Type == tiketType {
			tiketIndex = i
			break
		}
	}

	if tiketIndex == -1 {
		return domain.Order{}, errors.New("ticket type not found")
	}

	if event.Tiket[tiketIndex].Stock < quantity {
		return domain.Order{}, errors.New("not enough tickets")
	}

	total := float64(quantity) * event.Tiket[tiketIndex].Price
	if user.Balance < total {
		return domain.Order{}, errors.New("insufficient balance")
	}

	event.Tiket[tiketIndex].Stock -= quantity
	err = uc.EventRepo.UpdateEvent(ctx, &event)
	if err != nil {
		return domain.Order{}, err
	}

	// Mengurangi balance pengguna
	newBalance := user.Balance - total
	_, err = uc.UserRepo.UpdateBalance(ctx, userId, newBalance)
	if err != nil {
		return domain.Order{}, err
	}

	total = float64(quantity) * event.Tiket[tiketIndex].Price
	order := domain.Order{
		Status: "Confirmed",
		User:   user,
		Event:  event,
		Tiket:  []domain.Tiket{event.Tiket[tiketIndex]},
		Total:  total,
	}

	return uc.OrderRepo.CreateOrder(ctx, order)
}

// ListOrder implements OrderUsecaseInterface.
func (uc OrderUsecase) ListOrder(ctx context.Context) ([]domain.Order, error) {
	return uc.OrderRepo.ListOrder(ctx)
	
}
