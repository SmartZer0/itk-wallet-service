package tests

import (
	"context"
	"testing"

	"itk/internal/models"
	"itk/internal/repository"
	"itk/internal/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestProcessOperation_Deposit(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1 FOR UPDATE").
		WithArgs("test-wallet").
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(int64(1000)))

	mock.ExpectExec("UPDATE wallets SET balance = \\$1 WHERE id = \\$2").
		WithArgs(int64(1500), "test-wallet").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	repo := repository.NewWalletRepository(db)
	svc := service.NewWalletService(db, repo)

	err = svc.ProcessOperation(context.Background(), "test-wallet", models.Deposit, 500)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProcessOperation_Withdraw(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1 FOR UPDATE").
		WithArgs("wallet-123").
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(1500))

	mock.ExpectExec("UPDATE wallets SET balance = \\$1 WHERE id = \\$2").
		WithArgs(500, "wallet-123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	repo := repository.NewWalletRepository(db)
	service := service.NewWalletService(db, repo)

	err = service.ProcessOperation(context.Background(), "wallet-123", models.Withdraw, 1000)
	assert.NoError(t, err)
}

func TestProcessOperation_InsufficientFunds(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1 FOR UPDATE").
		WithArgs("wallet-123").
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(300))

	mock.ExpectRollback()

	repo := repository.NewWalletRepository(db)
	svc := service.NewWalletService(db, repo)

	err = svc.ProcessOperation(context.Background(), "wallet-123", models.Withdraw, 1000)
	assert.ErrorIs(t, err, service.ErrInsufficientFunds)
	assert.NoError(t, mock.ExpectationsWereMet())
}
