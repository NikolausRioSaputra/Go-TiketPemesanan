package usecase

import (
	"Go-TiketPemesanan/internal/domain"
	// "Go-TiketPemesanan/internal/repository"
	"Go-TiketPemesanan/internal/repositorydb"
	"context"
)

type EventUsecaseInterface interface {
	ListEvent
	CreateEvent
	GetEventById
}

type ListEvent interface {
	ListEvent(ctx context.Context) ([]domain.Event, error)
}

type CreateEvent interface {
	CreateEvent(ctx context.Context, event domain.Event) (domain.Event, error)
}

type GetEventById interface {
	GetEventById(ctx context.Context, id int) (domain.Event, error)
}

type EventUsecase struct {
	EventRepo repositorydb.EventRepositoryInterface
}


func NewEventUsecase(eventRepo repositorydb.EventRepositoryInterface) EventUsecase {
	return EventUsecase{
		EventRepo: eventRepo,
	}
}

func (uc EventUsecase) ListEvent(ctx context.Context) ([]domain.Event, error) {
	return uc.EventRepo.ListEvent(ctx)
}

func (uc EventUsecase) CreateEvent(ctx context.Context, event domain.Event) (domain.Event, error) {
	return uc.EventRepo.CreateEvent(ctx, &event)

}

func (uc EventUsecase) GetEventById(ctx context.Context, id int) (domain.Event, error) {
	return uc.EventRepo.GetEventById(ctx, id)
}