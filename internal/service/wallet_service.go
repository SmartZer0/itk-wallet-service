package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"itk/internal/models"
	"itk/internal/repository"
)

var (
	ErrInvalidAmount     = errors.New("amount must be greater than 0")
	ErrInsufficientFunds = errors.New("not enough funds in wallet")
	ErrUnknownOperation  = errors.New("unknown operation type")
)

type WalletService struct {
	repo repository.WalletRepo
	db   *sql.DB
}

func NewWalletService(db *sql.DB, repo repository.WalletRepo) *WalletService {
	return &WalletService{
		repo: repo,
		db:   db,
	}
}

func (s *WalletService) ProcessOperation(ctx context.Context, walletID string, opType models.OperationType, amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("cannot begin tx: %w", err)
	}
	defer func() { _ = tx.Rollback() }() // на случай паники или ошибки

	// Получаем баланс с блокировкой (FOR UPDATE)
	currentBalance, err := s.repo.GetBalanceForUpdate(ctx, tx, walletID)
	if err != nil {
		return fmt.Errorf("cannot get balance: %w", err)
	}

	var newBalance int64

	switch opType {
	case models.Deposit:
		newBalance = currentBalance + amount
	case models.Withdraw:
		if currentBalance < amount {
			return ErrInsufficientFunds
		}
		newBalance = currentBalance - amount
	default:
		return ErrUnknownOperation
	}

	if err := s.repo.UpdateBalance(ctx, tx, walletID, newBalance); err != nil {
		return fmt.Errorf("cannot update balance: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot commit tx: %w", err)
	}

	return nil
}
