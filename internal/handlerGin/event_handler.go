package handlergin

import (
	"Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/benebobaa/valo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type EventHandlerInterface interface {
	ListEvent(c *gin.Context)
	CreateEvent(c *gin.Context)
	GetEventById(c *gin.Context)
}

type EventHandler struct {
	EventUsecase usecase.EventUsecaseInterface
}

func NewEventHandler(eventUsecase usecase.EventUsecaseInterface) EventHandlerInterface {
	return &EventHandler{
		EventUsecase: eventUsecase,
	}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	var event domain.Event
	if err := c.ShouldBindBodyWithJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "Invalid request body",
			Errors:  err.Error(),
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("invalid request body")
		return
	}
	if err := valo.Validate(event); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "validation error",
			Errors:  err,
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("validation error")
		return
	}
	events, err := h.EventUsecase.CreateEvent(ctx, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMasage{
			Message: "Failed create event",
			Errors:  err.Error(),
		})
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseMasage{
		Message: "Success create event",
		Data:    events,
	})
	log.Info().
		Int("http.status.code", http.StatusCreated).
		TimeDiff("waktu process", time.Now(), start).
		Msg("Create Event API-Complated")
}

// GetEventById implements EventHandlerInterface.
func (h *EventHandler) GetEventById(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	eventId := c.Query("id")
	if eventId == "" {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "Event ID is required",
		})
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("event id is required")
		return
	}

	id, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "Invalid event ID",
		})
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid event id")
		return
	}

	events, err := h.EventUsecase.GetEventById(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMasage{
			Message: "Failed get event by ID",
			Errors:  err.Error(),
		})
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}
	c.JSON(http.StatusOK, ResponseMasage{
		Message: "Success get event by ID",
		Data:    events,
	})
	log.Info().
		Int("http.status.code", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("Get Event By ID API-Completed")
}

func (h *EventHandler) ListEvent(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	events, err := h.EventUsecase.ListEvent(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMasage{
			Message: "Failed get all events",
			Errors:  err.Error(),
		})
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}
	c.JSON(http.StatusOK, ResponseMasage{
		Message: "Success get all events",
		Data:    events,
	})
	log.Info().
		Int("http.status.code", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("Get All Event API-Completed")
}
