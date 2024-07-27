package usecase

import (
	"Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/repository"
	"errors"
)

type OrderUsecaseInterface interface {
	CreateOrder
	ListOrder
}

type CreateOrder interface {
	CreateOrder(userId, eventId int, tiketType string, quantity int) (domain.Order, error)
}

type ListOrder interface {
	ListOrder() ([]domain.Order, error)
}

type OrderUsecase struct {
	OrderRepo repository.OrderRepositoryinterface
	UserRepo  repository.UserRepositoryInterface
	EventRepo repository.EventRepositoryInterface
}

// CreateOrder implements OrderUsecaseInterface.
func (uc OrderUsecase) CreateOrder(userId int, eventId int, tiketType string, quantity int) (domain.Order, error) {
	user, err := uc.UserRepo.UserFindById(userId)
	if err != nil {
		return domain.Order{}, err
	}

	event, err := uc.EventRepo.GetEventById(eventId)
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

	event.Tiket[tiketIndex].Stock -= quantity
	err = uc.EventRepo.UpdateEvent(event)
	if err != nil {
		return domain.Order{}, err
	}

	total := float64(quantity) * event.Tiket[tiketIndex].Price
	order := domain.Order{
		Status: "Confirmed",
		User:   user,
		Event:  event,
		Tiket:  []domain.Tiket{event.Tiket[tiketIndex]},
		Total:  total,
	}

	return uc.OrderRepo.CreateOrder(order)
}

// ListOrder implements OrderUsecaseInterface.
func (uc OrderUsecase) ListOrder() ([]domain.Order, error) {
	return uc.OrderRepo.ListOrder()
}

func NewOrderUsecase(orderRepo repository.OrderRepositoryinterface, userRepo repository.UserRepositoryInterface, eventRepo repository.EventRepositoryInterface) OrderUsecase {
	return OrderUsecase{
		OrderRepo: orderRepo,
		UserRepo:  userRepo,
		EventRepo: eventRepo,
	}
}
