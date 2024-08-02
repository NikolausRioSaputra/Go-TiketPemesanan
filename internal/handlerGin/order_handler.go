package handlergin

import (
	// "Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/usecase"
	"net/http"
	"time"

	// "github.com/benebobaa/valo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type OrderHandlerInterface interface {
	CreateOrder(c *gin.Context)
	ListOrders(c *gin.Context)
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

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	var req struct {
		UserID    int    `json:"user_id"`
		EventID   int    `json:"event_id"`
		TiketType string `json:"tiket_type"`
		Quantity  int    `json:"quantity"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid request body")
		return
	}

	order, err := h.OrderUsecase.CreateOrder(ctx, req.UserID, req.EventID, req.TiketType, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Message: "Failed to create order",
			Data:    err.Error(),
		})
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}
	c.JSON(http.StatusOK, ResponseMessage{
		Message: "Success create order",
		Data:    order,
	})
	log.Info().
		Int("http.status.code", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("Create Order Tiket API-Completed")

}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	orders, err := h.OrderUsecase.ListOrder(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Message: "Failed to list orders",
			Data:    err.Error(),
		})
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseMessage{
		Message: "Success get all orders",
		Data:    orders,
	})
	log.Info().
		Int("http.status.code", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("List Orders Tiket API-Completed")
}
