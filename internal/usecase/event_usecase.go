package usecase

import (
	"Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/repository"
)

type EventUsecaseInterface interface {
	ListEvent
	GetEventById
}

type ListEvent interface {
	ListEvent() ([]domain.Event, error)
}

type GetEventById interface {
	GetEventById(id int) (domain.Event, error)
}

type EventUsecase struct {
	EventRepo repository.EventRepositoryInterface
}


func NewEventUsecase(eventRepo repository.EventRepositoryInterface) EventUsecase {
	return EventUsecase{
		EventRepo: eventRepo,
	}
}

func (uc EventUsecase) ListEvent() ([]domain.Event, error) {
	return uc.EventRepo.ListEvent()
}

func (uc EventUsecase) GetEventById(id int) (domain.Event, error) {
	return uc.EventRepo.GetEventById(id)
}
