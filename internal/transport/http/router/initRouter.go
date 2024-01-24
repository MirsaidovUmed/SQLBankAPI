package router

import (
	"sqlBankCLI/internal/transport/http/handlers"
	"sqlBankCLI/pkg/http"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler) *mux.Router {
	router := http.NewRouter()

	router.HandleFunc("/api/createBankAccount", handlers.CreateBankAccount).Methods("POST")

	router.HandleFunc("/api/topUpClientsAccount", handlers.TopUpClientsAccount).Methods("POST")

	router.HandleFunc("/api/withdrawClientAccount", handlers.WithdrawClientAccount).Methods("POST")

	router.HandleFunc("/api/transferMoney", handlers.TransferMoney).Methods("POST")

	router.HandleFunc("/api/showAccountBalance", handlers.ShowAccountBalance).Methods("GET")

	return router
}
