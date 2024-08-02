package handler

// import (
// 	"Go-TiketPemesanan/internal/domain"
// 	"Go-TiketPemesanan/internal/usecase"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/rs/zerolog/log"
// )

// type EventHandlerInterface interface {
// 	ListEvent(w http.ResponseWriter, r *http.Request)
// 	CreateEvent(w http.ResponseWriter, r *http.Request)
// 	GetEventById(w http.ResponseWriter, r *http.Request)

// }

// type EventHandler struct {
// 	EventUsecase usecase.EventUsecaseInterface
// }

// func NewEventHandler(eventUsecase usecase.EventUsecaseInterface) EventHandlerInterface {
// 	return &EventHandler{
// 		EventUsecase: eventUsecase,
// 	}
// }

// func (h *EventHandler) CreateEvent (w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	start := time.Now()
// 	if r.Method != "POST" {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		log.Info().
// 			Int("http.status.code", http.StatusMethodNotAllowed).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg("invalid method")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var event domain.Event
// 	err := json.NewDecoder(r.Body).Decode(&event)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		log.Error().
// 			Int("http.status.code", http.StatusBadRequest).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg(err.Error())
// 		return
// 	}

// 	events, err := h.EventUsecase.CreateEvent(ctx ,event)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Error().
// 			Int("http.status.code", http.StatusInternalServerError).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg(err.Error())
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	err = json.NewEncoder(w).Encode(ResponseMasage{
// 		Message: "Success create event",
// 		Data:    events,
// 	})
// 	log.Info().
// 		Int("http.status.code", http.StatusCreated).
// 		TimeDiff("waktu process", time.Now(), start).
// 		Msg("Create Event API-Complated")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Error().
// 			Int("http.status.code", http.StatusInternalServerError).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg(err.Error())
// 		return
// 	}
// }

// // GetEventById implements EventHandlerInterface.
// func (h *EventHandler) GetEventById(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	start := time.Now()
// 	if r.Method != "GET" {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		log.Info().
// 			Int("http.status.code", http.StatusMethodNotAllowed).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg("invalid method")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	// Get the event ID from the request parameters
// 	eventId := r.URL.Query().Get("id")
// 	// Validate the event ID
// 	if eventId == "" {
// 		http.Error(w, "Event ID is required", http.StatusBadRequest)
// 		log.Error().
// 			Int("http.status.code", http.StatusBadRequest).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg("event id is required")
// 		return
// 	}

// 	id, err := strconv.Atoi(eventId)
// 	// Validate the event ID
// 	if err != nil {
// 		http.Error(w, "Invalid event ID", http.StatusBadRequest)
// 		log.Error().
// 			Int("http.status.code", http.StatusBadRequest).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg("invalid event id")
// 		return
// 	}
// 	// Call the use case to get the event by ID
// 	event, err := h.EventUsecase.GetEventById(ctx, id)
// 	// Handle any errors
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Error().
// 			Int("http.status.code", http.StatusInternalServerError).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg(err.Error())
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(ResponseMasage{
// 		Message: "Success get event by id",
// 		Data:    event,
// 	})
// 	log.Info().
// 		Int("http.status.code", http.StatusOK).
// 		TimeDiff("waktu process", time.Now(), start).
// 		Msg("Get Event By ID API-Complated")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Error().
// 			Int("http.status.code", http.StatusInternalServerError).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg(err.Error())
// 		return
// 	}

// }

// func (h *EventHandler) ListEvent(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	start := time.Now()
// 	if r.Method != "GET" {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		log.Info().
// 			Int("http.status.code", http.StatusMethodNotAllowed).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg("invalid method")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	events, err := h.EventUsecase.ListEvent(ctx)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Error().
// 			Int("http.status.code", http.StatusInternalServerError).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg(err.Error())
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(ResponseMasage{
// 		Message: "Success get all events",
// 		Data:    events,
// 	})
// 	log.Info().
// 		Int("http.status.code", http.StatusOK).
// 		TimeDiff("waktu process", time.Now(), start).
// 		Msg("Get All Event API-Completed")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Error().
// 			Int("http.status.code", http.StatusInternalServerError).
// 			TimeDiff("waktu process", time.Now(), start).
// 			Msg(err.Error())
// 		return
// 	}
// }
