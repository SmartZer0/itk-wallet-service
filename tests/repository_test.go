package tests

import (
	"context"
	"testing"

	"itk/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetWallet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewWalletRepository(db)
	walletID := uuid.New()
	initialBalance := int64(1000)

	mock.ExpectExec(`INSERT INTO wallets`).WithArgs(walletID, initialBalance).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateWallet(context.Background(), walletID, initialBalance)
	assert.NoError(t, err)

	mock.ExpectQuery(`SELECT id, balance FROM wallets WHERE id = \$1`).
		WithArgs(walletID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(walletID, initialBalance))

	wallet, err := repo.GetWalletByID(context.Background(), walletID)
	assert.NoError(t, err)
	assert.Equal(t, walletID, wallet.ID)
	assert.Equal(t, initialBalance, wallet.Balance)
}

func TestGetBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewWalletRepository(db)
	walletID := uuid.New()
	balance := int64(1500)

	mock.ExpectQuery(`SELECT balance FROM wallets WHERE id = \$1`).
		WithArgs(walletID).
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(balance))

	got, err := repo.GetBalance(context.Background(), walletID)
	assert.NoError(t, err)
	assert.Equal(t, balance, got)
}
