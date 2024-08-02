package routes

import (
	handlergin "Go-TiketPemesanan/internal/handlerGin"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, userHandler handlergin.UserHandlerInterface, eventHandler handlergin.EventHandlerInterface, orderHandler handlergin.OrderHandlerInterface) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userHandler.StoreNewUser)
		userRoutes.GET("/:id", userHandler.UserFindById)
		userRoutes.GET("/", userHandler.GetAllUser)
		userRoutes.PUT("/:id", userHandler.UserUpdater)
		userRoutes.DELETE("/:id", userHandler.UserDeleter)
	}

	eventRoutes := router.Group("/events")
	{
		eventRoutes.GET("/", eventHandler.ListEvent)
		eventRoutes.GET("/:id", eventHandler.GetEventById)
		eventRoutes.POST("/", eventHandler.CreateEvent)
	}

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("/", orderHandler.ListOrders)
		orderRoutes.POST("/", orderHandler.CreateOrder)

	}
}
