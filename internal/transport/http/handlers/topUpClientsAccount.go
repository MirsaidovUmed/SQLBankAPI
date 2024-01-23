package handlers

import (
	"encoding/json"
	"net/http"
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/response"
)

func (h *Handler) TopUpClientsAccount(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var req models.ChangeBalanceRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		resp = response.BadRequest
		return
	}

	err = h.svc.TopUpClientsAccount(req.Account, req.Amount)
	if err != nil {
		resp = response.InternalServer
		return
	}
	resp = response.Success
}
