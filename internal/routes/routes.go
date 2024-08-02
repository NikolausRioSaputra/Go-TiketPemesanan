package routes

import (
	handlergin "Go-TiketPemesanan/internal/handlerGin"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, userHandler handlergin.UserHandlerInterface, eventHandler handlergin.EventHandlerInterface, orderHandler handlergin.OrderHandlerInterface) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/create", userHandler.StoreNewUser)
		userRoutes.GET("/findbyid", userHandler.UserFindById)
		userRoutes.GET("/all", userHandler.GetAllUser)
		userRoutes.PUT("/edit", userHandler.UserUpdater)
		userRoutes.DELETE("/delete", userHandler.UserDeleter)
	}

	eventRoutes := router.Group("/events")
	{
		eventRoutes.GET("/all", eventHandler.ListEvent)
		eventRoutes.GET("/findbyid", eventHandler.GetEventById)
		eventRoutes.POST("/create", eventHandler.CreateEvent)
	}

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("/all", orderHandler.ListOrders)
		orderRoutes.POST("/create", orderHandler.CreateOrder)

	}
}
