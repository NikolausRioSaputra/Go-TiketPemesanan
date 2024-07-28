package main

import (
	"Go-TiketPemesanan/internal/handler"
	"Go-TiketPemesanan/internal/repository"
	"Go-TiketPemesanan/internal/usecase"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

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
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	
	<-stop
	fmt.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	
	wg.Wait()
}
