package repository

import (
	"context"
	"database/sql"
	"fmt"

	"itk/internal/models"

	"github.com/google/uuid"
)

type WalletRepo interface {
	GetBalanceForUpdate(ctx context.Context, tx *sql.Tx, walletID string) (int64, error)
	UpdateBalance(ctx context.Context, tx *sql.Tx, walletID string, newBalance int64) error
}

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetWalletByID(ctx context.Context, id uuid.UUID) (*models.Wallet, error) {
	var w models.Wallet
	row := r.db.QueryRowContext(ctx, "SELECT id, balance FROM wallets WHERE id = $1", id)
	err := row.Scan(&w.ID, &w.Balance)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *WalletRepository) CreateWallet(ctx context.Context, id uuid.UUID, initialAmount int64) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO wallets (id, balance) VALUES ($1, $2)
		ON CONFLICT (id) DO NOTHING`,
		id, initialAmount,
	)
	return err
}

func (r *WalletRepository) GetBalanceForUpdate(ctx context.Context, tx *sql.Tx, walletID string) (int64, error) {
	var balance int64
	err := tx.QueryRowContext(ctx,
		"SELECT balance FROM wallets WHERE id = $1 FOR UPDATE", walletID).Scan(&balance)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("wallet not found")
	}
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, tx *sql.Tx, walletID string, newBalance int64) error {
	_, err := tx.ExecContext(ctx,
		"UPDATE wallets SET balance = $1 WHERE id = $2", newBalance, walletID)
	return err
}

func (r *WalletRepository) GetBalance(ctx context.Context, id uuid.UUID) (int64, error) {
	var balance int64
	err := r.db.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id = $1", id).Scan(&balance)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("wallet not found")
	}
	if err != nil {
		return 0, err
	}
	return balance, nil
}
