package main

import (
	// "Go-TiketPemesanan/internal/handler"
	"Go-TiketPemesanan/internal/handlerGin"
	"Go-TiketPemesanan/internal/provider/db"
	"Go-TiketPemesanan/internal/repositorydb"
	"Go-TiketPemesanan/internal/routes"
	"Go-TiketPemesanan/internal/usecase"
	"fmt"
	"log"
	// "net/http"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
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
	userHandler := handlergin.NewUserHandler(userUsacase)

	eventRepo := repositorydb.NewEventRepository(db)
	eventUsecase := usecase.NewEventUsecase(eventRepo)
	eventHandler := handlergin.NewEventHandler(eventUsecase)

	orderRepo := repositorydb.NewOrderRepository(db)
	orderService := usecase.NewOrderUsecase(orderRepo, userRepo, eventRepo)
	orderHandler := handlergin.NewOrderHandler(orderService)

	router := gin.Default()
	routes.InitializeRoutes(router, userHandler, eventHandler, orderHandler)

	fmt.Printf("Server running on :8080")
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := router.Run(":8080")
		if err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()
	wg.Wait()
}
