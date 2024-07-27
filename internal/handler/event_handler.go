package handler

import (
	"Go-TiketPemesanan/internal/usecase"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type EventHandlerInterface interface {
	ListEvent(w http.ResponseWriter, r *http.Request)
}

type EventHandler struct {
	EventUsecase usecase.EventUsecaseInterface
}

func NewEventHandler(eventUsecase usecase.EventUsecaseInterface) EventHandlerInterface {
	return &EventHandler{
		EventUsecase: eventUsecase,
	}
}

func (h *EventHandler) ListEvent(w http.ResponseWriter, r *http.Request) {
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
	events, err := h.EventUsecase.ListEvent()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ResponseMasage{
		Message: "Success get all events",
		Data:    events,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}
}
