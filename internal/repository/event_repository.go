package repository

// import (
// 	"Go-TiketPemesanan/internal/domain"
// 	"errors"
// 	"sync"
// )

// type EventRepositoryInterface interface {
// 	ListEvent
// 	GetEventById
// 	UpdateEven
// }

// type ListEvent interface {
// 	ListEvent() ([]domain.Event, error)
// }

// type GetEventById interface {
// 	GetEventById(id int) (domain.Event, error)
// }

// type UpdateEven interface {
// 	UpdateEvent(event domain.Event) error
// }

// type EventRepository struct {
// 	events map[int]domain.Event
// 	mu sync.Mutex
// }

// func NewEventRepository() EventRepositoryInterface {
// 	return &EventRepository {
// 		events: map[int]domain.Event{
// 			1: {ID: 1, Name: "Sulap Beneran", Date: "2024-08-01", Location: "Samarinda", Tiket: []domain.Tiket{
// 				{ID: 1, Type: "VIP", Price: 5000, Stock: 10},
// 				{ID: 2, Type: "CAT1", Price: 250, Stock: 100},
// 			}},
// 			2: {ID: 2, Name: "Dangdut Pro", Date: "2024-09-01", Location: "Jakarta", Tiket: []domain.Tiket{
// 				{ID: 1, Type: "VIP", Price: 5000, Stock: 10},
// 				{ID: 2, Type: "CAT1", Price: 250, Stock: 100},
// 			}},
// 		},

// 		mu: sync.Mutex{},
// 	}
// }

// func (repo *EventRepository) ListEvent() ([]domain.Event, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	events := []domain.Event{}
// 	for _, event := range repo.events{
// 		events = append(events, event)
// 	}
// 	return events, nil
// }

// func (repo *EventRepository) GetEventById(id int) (domain.Event, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	event, ok := repo.events[id]
// 	if !ok {
// 		return domain.Event{}, errors.New("event not found")
// 	}
// 	return event, nil
// }

// func (repo *EventRepository) UpdateEvent(event domain.Event) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()
	
// 	repo.events[event.ID] = event
// 	return nil
// }