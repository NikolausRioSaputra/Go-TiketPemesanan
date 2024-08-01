package repository

// import (
// 	"Go-TiketPemesanan/internal/domain"
// 	"sync"
// 	"time"
// )

// type OrderRepositoryinterface interface {
// 	CreateOrder
// 	ListOrder
// }

// type CreateOrder interface {
// 	CreateOrder(order domain.Order) (domain.Order, error)
// }

// type ListOrder interface {
// 	ListOrder() ([]domain.Order, error)
// }

// type OrderRepository struct {
// 	orders map[int]domain.Order
// 	mu     sync.Mutex
// }

// func NewOrderRepository() OrderRepositoryinterface {
// 	return &OrderRepository{
// 		orders: map[int]domain.Order{},
// 		mu: sync.Mutex{},
// 	}
// }

// // CreateOrder implements OrderRepositoryinterface.
// func (repo *OrderRepository) CreateOrder(order domain.Order) (domain.Order, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	order.ID = len(repo.orders) + 1
// 	order.Date = time.Now().Format("2006-01-02")
// 	repo.orders[order.ID] = order

// 	return order, nil
// }

// // ListOrder implements OrderRepositoryinterface.
// func (repo *OrderRepository) ListOrder() ([]domain.Order, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	orders := []domain.Order{}
// 	for _, order := range repo.orders {
// 		orders = append(orders, order)
// 	}
// 	return orders, nil
// }
