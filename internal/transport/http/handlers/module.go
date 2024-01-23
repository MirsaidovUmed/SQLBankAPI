package handlers

import "sqlBankCLI/internal/service"

type Handler struct {
	svc service.ServiceInterface
}

func NewHandler(svc service.ServiceInterface) *Handler {
	return &Handler{
		svc: svc,
	}
}
