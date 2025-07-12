package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"itk/internal/models"
	"itk/internal/repository"
	"itk/internal/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type WalletHandler struct {
	Repo    *repository.WalletRepository
	Service *service.WalletService
}

func NewWalletHandler(repo *repository.WalletRepository, service *service.WalletService) *WalletHandler {
	return &WalletHandler{
		Repo:    repo,
		Service: service,
	}
}

type OperationRequest struct {
	WalletID      uuid.UUID            `json:"walletId"`
	OperationType models.OperationType `json:"operationType"`
	Amount        int64                `json:"amount"`
}

// POST /api/v1/wallet
func (h *WalletHandler) HandleOperation(w http.ResponseWriter, r *http.Request) {
	var req OperationRequest

	body, _ := io.ReadAll(r.Body)
	log.Println("BODY:", string(body))
	r.Body = io.NopCloser(bytes.NewReader(body))

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("DECODE ERROR:", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	wallet, err := h.Repo.GetWalletByID(ctx, req.WalletID)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	if wallet == nil {
		err = h.Repo.CreateWallet(ctx, req.WalletID, 0)
		if err != nil {
			http.Error(w, "create error", http.StatusInternalServerError)
			return
		}
	}

	err = h.Service.ProcessOperation(ctx, req.WalletID.String(), req.OperationType, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GET /api/v1/wallets/{id}
func (h *WalletHandler) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	walletID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	wallet, err := h.Repo.GetWalletByID(ctx, walletID)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	if wallet == nil {
		http.Error(w, "wallet not found", http.StatusNotFound)
		return
	}

	resp := map[string]interface{}{
		"walletId": wallet.ID,
		"balance":  wallet.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
