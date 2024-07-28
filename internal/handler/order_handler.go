package handler

import (
	// "Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/usecase"
	"encoding/json"
	"net/http"
	"time"

	// "github.com/benebobaa/valo"
	"github.com/rs/zerolog/log"
)

type OrderHandlerInterface interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	ListOrders(w http.ResponseWriter, r *http.Request)
}

type OrderHandler struct {
	OrderUsecase usecase.OrderUsecaseInterface
}

func NewOrderHandler(orderUsecase usecase.OrderUsecaseInterface) OrderHandlerInterface {
	return &OrderHandler{
		OrderUsecase: orderUsecase,
	}
}

type ResponseMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	start := time.Now()
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Info().
			Int("http.status.code", http.StatusMethodNotAllowed).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	var req struct {
		UserID    int    `json:"user_id"`
		EventID   int    `json:"event_id"`
		TiketType string `json:"tiket_type"`
		Quantity  int    `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid request body")
		return
	}

	select {
		case <-time.After(5 * time.Second): 
	order, err := h.OrderUsecase.CreateOrder(req.UserID, req.EventID, req.TiketType, req.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ResponseMessage{
		Message: "Success create order",
		Data:    order,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	log.Info().
		Int("http.status.code", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("Create Order API-Completed")
	case <-ctx.Done():
		// Operasi dibatalkan
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		log.Info().
			Int("http.status.code", http.StatusRequestTimeout).
			TimeDiff("waktu process", time.Now(), start).
			Msg("request canceled")
		return
	}
}
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Info().
			Int("http.status.code", http.StatusMethodNotAllowed).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	orders, err := h.OrderUsecase.ListOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ResponseMessage{
		Message: "Success get all orders",
		Data:    orders,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	log.Info().
		Int("http.status.code", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("List Orders API-Completed")
}