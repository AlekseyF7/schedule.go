package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"main/internal/config"
	"main/internal/http-server/handlers"
	"main/internal/scheduler"
	"net/http"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//инициализация конфига
	cfg := config.GetConfig()
	fmt.Println(cfg)

	//инициализация scheduler
	sche := scheduler.NewSche()

	//инициализация хендлера
	newHandler := handlers.NewHandler(sche)

	//инициализация роутера
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // Разрешить запросы со всех источников
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	//эндпойнты
	router.Post("/create-schedule", newHandler.GetScheduleHandler)

	//конфиг сервера
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	//инициализация сервера
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("failed to start server")
	}

	fmt.Println("server stopped!!")
}
