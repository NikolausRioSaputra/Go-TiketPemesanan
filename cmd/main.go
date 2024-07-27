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

	routes := http.NewServeMux()
	routes.HandleFunc("/users", userHandler.StoreNewUser)
	routes.HandleFunc("/users/all", userHandler.GetAllUser)
	routes.HandleFunc("/users/update", userHandler.UpdateUser)
	routes.HandleFunc("/users/delete",userHandler.DeleteUser)
	routes.HandleFunc("/events", eventHandler.ListEvent)

	server := http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	fmt.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal("Server failed to start: ", err)
	}
}
