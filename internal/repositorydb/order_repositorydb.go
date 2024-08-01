package repositorydb

import (
	"Go-TiketPemesanan/internal/domain"
	"context"
	"database/sql"
	"sync"
)

type OrderRepositoryinterface interface {
	CreateOrder
	ListOrder
}

type CreateOrder interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}

type ListOrder interface {
	ListOrder(ctx context.Context) ([]domain.Order, error)
}

type OrderRepository struct {
	DB *sql.DB
	mu sync.Mutex
}

func NewOrderRepository(db *sql.DB) OrderRepositoryinterface {
	return &OrderRepository{
		DB: db,
		mu: sync.Mutex{},
	}
}

// CreateOrder implements OrderRepositoryinterface.
func (repo *OrderRepository) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return domain.Order{}, err
	}

	query := "INSERT INTO orders(date, status, user_id, event_id, total) VALUES (CURRENT_DATE, $1,$2,$3,$4) RETURNING id"
	err = tx.QueryRowContext(ctx, query, order.Status, order.User.ID, order.Event.ID, order.Total).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	for _, ticket := range order.Tiket {
		query = "INSERT INTO order_tikets(order_id, tiket_id, quantity) VALUES ($1, $2, $3)"
		_, err = tx.ExecContext(ctx, query, order.ID, ticket.ID, ticket.Stock)
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

// ListOrder implements OrderRepositoryinterface.
func (repo *OrderRepository) ListOrder(ctx context.Context) ([]domain.Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	query := `SELECT o.id, o.date, o.status, u.id, u.name, u.address, 
					e.id, e.name, e.date, e.location,
					t.id, t.stock, t.type, t.price, 
					o.total
			FROM orders o
			JOIN users u ON o.user_id = u.id
			JOIN events e ON o.event_id = e.id
			JOIN tikets t ON o.id = t.id`

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var user domain.User
		var event domain.Event
		var tiket domain.Tiket

		err := rows.Scan(&order.ID, &order.Date, &order.Status,
			&user.ID, &user.Name, &user.Address,
			&event.ID, &event.Name, &event.Date, &event.Location,
			&tiket.ID, &tiket.Stock, &tiket.Type, &tiket.Price,
			&order.Total)
		if err != nil {
			return nil, err
		}

		order.User = user
		order.Event = event
		order.Tiket =  append(order.Tiket, tiket)

		orders = append(orders, order)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return orders, nil
}
