package main

import (
	"Go-TiketPemesanan/internal/handler"
	"Go-TiketPemesanan/internal/repository"
	"Go-TiketPemesanan/internal/usecase"
	"fmt"
	"log"
	"net/http"
)

func main() {
	userRepo := repository.NewUserRepository()
	userUsacase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsacase)

	eventRepo := repository.NewEventRepository()
	eventUsecase := usecase.NewEventUsecase(eventRepo)
	eventHandler := handler.NewEventHandler(eventUsecase)

	orderRepo := repository.NewOrderRepository()
	orderService := usecase.NewOrderUsecase(orderRepo, userRepo, eventRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	routes := http.NewServeMux()
	routes.HandleFunc("/users", userHandler.StoreNewUser)
	routes.HandleFunc("/users/findbyid", userHandler.UserFindById)
	routes.HandleFunc("/users/all", userHandler.GetAllUser)
	routes.HandleFunc("/users/update", userHandler.UpdateUser)
	routes.HandleFunc("/users/delete", userHandler.DeleteUser)

	routes.HandleFunc("/events", eventHandler.ListEvent)
	routes.HandleFunc("/events/findbyid", eventHandler.GetEventById)

	routes.HandleFunc("/book", orderHandler.CreateOrder)
	routes.HandleFunc("/orders", orderHandler.ListOrders)

	server := http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	fmt.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
