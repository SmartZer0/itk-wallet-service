package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"itk/internal/handler"
	"itk/internal/repository"
	"itk/internal/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupHandler(t *testing.T) (*handler.WalletHandler, sqlmock.Sqlmock, string) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := repository.NewWalletRepository(db)
	svc := service.NewWalletService(db, repo)
	h := handler.NewWalletHandler(repo, svc)

	walletID := uuid.New().String()

	return h, mock, walletID
}

func TestWalletHandler_Deposit(t *testing.T) {
	h, mock, walletID := setupHandler(t)

	// Моки для обработки запроса
	mock.ExpectQuery(`SELECT id, balance FROM wallets WHERE id = \$1`).
		WithArgs(walletID).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectExec(`INSERT INTO wallets`).
		WithArgs(walletID, int64(0)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectBegin()

	mock.ExpectQuery(`SELECT balance FROM wallets WHERE id = \$1 FOR UPDATE`).
		WithArgs(walletID).
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(int64(0)))

	mock.ExpectExec(`UPDATE wallets SET balance = \$1 WHERE id = \$2`).
		WithArgs(int64(500), walletID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	body := map[string]interface{}{
		"walletId":      walletID,
		"operationType": "DEPOSIT",
		"amount":        500,
	}
	bodyJSON, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewReader(bodyJSON))
	rec := httptest.NewRecorder()

	h.HandleOperation(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		t.Fatalf("unexpected status: got %v, want %v, body: %s", resp.StatusCode, http.StatusOK, b)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestWalletHandler_GetBalance_InvalidUUID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewWalletRepository(db)
	svc := service.NewWalletService(db, repo)
	h := handler.NewWalletHandler(repo, svc)

	// Плохой UUID
	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/invalid-uuid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid-uuid"})
	rec := httptest.NewRecorder()

	h.HandleGetBalance(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Contains(t, string(body), "invalid UUID")

	// Проверяем, что не было никаких запросов в БД
	assert.NoError(t, mock.ExpectationsWereMet())
}
