package handlers

import (
	"encoding/json"
	"net/http"
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/response"
)

func (h *Handler) TransferMoney(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var req models.TransferRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		resp = response.BadRequest
		return
	}

	err = h.svc.TransferMoney(req.SenderPhone, req.RecipientPhone, req.Amount)

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
