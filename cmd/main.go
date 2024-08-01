package main

import (
	"Go-TiketPemesanan/internal/handler"
	"Go-TiketPemesanan/internal/provider/db"
	"Go-TiketPemesanan/internal/repositorydb"
	"Go-TiketPemesanan/internal/usecase"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	db, err := db.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var wg sync.WaitGroup

	userRepo := repositorydb.NewUserRepository(db)
	userUsacase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsacase)

	eventRepo := repositorydb.NewEventRepository(db)
	eventUsecase := usecase.NewEventUsecase(eventRepo)
	eventHandler := handler.NewEventHandler(eventUsecase)

	orderRepo := repositorydb.NewOrderRepository(db)
	orderService := usecase.NewOrderUsecase(orderRepo, userRepo, eventRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	routes := http.NewServeMux()
	routes.HandleFunc("/users", userHandler.StoreNewUser)
	routes.HandleFunc("/users/findbyid", userHandler.UserFindById)
	routes.HandleFunc("/users/all", userHandler.GetAllUser)
	routes.HandleFunc("/users/update", userHandler.UserUpdater)
	routes.HandleFunc("/users/delete", userHandler.UserDeleter)

	routes.HandleFunc("/events", eventHandler.ListEvent)
	routes.HandleFunc("/events/findbyid", eventHandler.GetEventById)
	routes.HandleFunc("/events/create", eventHandler.CreateEvent)

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
	wg.Wait()
}
