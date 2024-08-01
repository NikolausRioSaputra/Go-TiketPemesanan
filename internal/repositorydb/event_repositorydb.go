package repositorydb

import (
	"Go-TiketPemesanan/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"sync"
	// "github.com/rs/zerolog/log"
)

type EventRepositoryInterface interface {
	CreateEvent
	GetEventById
	ListEvent
	UpdateEvent
}

type CreateEvent interface {
	CreateEvent(ctx context.Context, event *domain.Event) (domain.Event, error)
}

type GetEventById interface {
	GetEventById(ctx context.Context, id int) (domain.Event, error)
}

type ListEvent interface {
	ListEvent(ctx context.Context) ([]domain.Event, error)
}

type UpdateEvent interface {
	UpdateEvent(ctx context.Context, event *domain.Event) error
}

type EventRepository struct {
	DB *sql.DB
	mu sync.Mutex
}

func NewEventRepository(db *sql.DB) EventRepositoryInterface {
	return &EventRepository{
		DB: db,
		mu: sync.Mutex{},
	}
}

// CreateEvent implements EventRepositoryInterface.
func (repo *EventRepository) CreateEvent(ctx context.Context, event *domain.Event) (domain.Event, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return domain.Event{}, err
	}
	query := "INSERT INTO events(name, date, location) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRowContext(ctx, query, event.Name, event.Date, event.Location).Scan(&event.ID)
	if err != nil {
		tx.Rollback()
		return domain.Event{}, err
	}

	var tempTiket []domain.Tiket
	for _, ticket := range event.Tiket {

		query = "INSERT INTO tikets(event_id, stock, type, price) VALUES ($1, $2, $3, $4) RETURNING id"
		err = tx.QueryRowContext(ctx, query, event.ID, ticket.Stock, ticket.Type, ticket.Price).Scan(&ticket.ID)

		if err != nil {
			tx.Rollback()
			return domain.Event{}, err
		}

		tempTiket = append(tempTiket, ticket)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return domain.Event{}, err
	}

	event.Tiket = tempTiket

	return *event, nil
}

// GetEventById implements EventRepositoryInterface.
func (repo *EventRepository) GetEventById(ctx context.Context, id int) (domain.Event, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var event domain.Event
	query := "SELECT id, name, date, location FROM events WHERE id = $1"
	err := repo.DB.QueryRowContext(ctx, query, id).Scan(&event.ID, &event.Name, &event.Date, &event.Location)
	if err != nil {
		return domain.Event{}, fmt.Errorf("event not found")
	}

	rows, err := repo.DB.QueryContext(ctx, "SELECT id, stock, type, price FROM tikets where event_id = $1", event.ID)
	if err != nil {
		return domain.Event{}, err
	}

	for rows.Next() {
		var tiket domain.Tiket
		if err := rows.Scan(&tiket.ID, &tiket.Stock, &tiket.Type, &tiket.Price); err != nil {
			return domain.Event{}, err
		}
		event.Tiket = append(event.Tiket, tiket)
	}

	return event, nil
}

// ListEvent implements EventRepositoryInterface.
func (repo *EventRepository) ListEvent(ctx context.Context) ([]domain.Event, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var events []domain.Event
	query := "SELECT id, name, date, location FROM events"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}

		tiketsRows, err := repo.DB.QueryContext(ctx, "SELECT id, stock, type, price FROM tikets WHERE event_id = $1", event.ID)
		if err != nil {
			return nil, err
		}

		defer tiketsRows.Close()
		for tiketsRows.Next() {
			var tiket domain.Tiket
			err := tiketsRows.Scan(&tiket.ID, &tiket.Stock, &tiket.Type, &tiket.Price)
			if err != nil {
				return nil, err
			}
			event.Tiket = append(event.Tiket, tiket)
		}

		if err := tiketsRows.Err(); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

// UpdateEvent implements EventRepositoryInterface.
func (repo *EventRepository) UpdateEvent(ctx context.Context, event *domain.Event) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, tiket := range event.Tiket {
		query := "UPDATE tikets SET stock = $1 WHERE id = $2"
		_, err = tx.ExecContext(ctx, query, tiket.Stock, tiket.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
