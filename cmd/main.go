package main

import (
	"sqlBankCLI/internal/repository"
	"sqlBankCLI/internal/service"
	"sqlBankCLI/internal/transport/http/handlers"
	"sqlBankCLI/internal/transport/http/router"
	"sqlBankCLI/pkg/config"
	"sqlBankCLI/pkg/database"
	"sqlBankCLI/pkg/http"
)

func main() {
	conf := config.NewConfig()
	db := database.NewDatabase(conf)
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)

	handlers := handlers.NewHandler(svc)

	router := router.InitRouter(handlers)

	server := http.NewServer(conf.ServerAddress, conf.ServerPort, router)

	server.Run()
}
