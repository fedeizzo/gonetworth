package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/fedeizzo/gonetworth/internal/db"
)

var (
	ErrInvalidName = errors.New("invalid account name")
)

type AccountService interface {
	GetAccounts(ctx context.Context) ([]db.Account, error)
	GetExpenseAccounts(ctx context.Context) ([]db.ExpenseAccount, error)
	GetIncomeAccounts(ctx context.Context) ([]db.IncomeAccount, error)
	UpdateAccount(ctx context.Context, account db.Account) error
	UpdateExpenseAccount(ctx context.Context, account db.ExpenseAccount) error
	UpdateIncomeAccount(ctx context.Context, account db.IncomeAccount) error
	CreateAccount(ctx context.Context, name, notes string) error
	CreateExpenseAccount(ctx context.Context, name, notes string) error
	CreateIncomeAccount(ctx context.Context, name, notes string) error
}

type accountService struct {
	querier db.Querier
}

func New(querier db.Querier) *accountService {
	return &accountService{
		querier: querier,
	}
}

func (a *accountService) GetAccounts(ctx context.Context) ([]db.Account, error) {
	return a.querier.GetAccounts(ctx)
}

func (a *accountService) GetExpenseAccounts(ctx context.Context) ([]db.ExpenseAccount, error) {
	return a.querier.GetExpenseAccounts(ctx)
}

func (a *accountService) GetIncomeAccounts(ctx context.Context) ([]db.IncomeAccount, error) {
	return a.querier.GetIncomeAccounts(ctx)
}

func (a *accountService) UpdateAccount(ctx context.Context, account db.Account) error {
	if account.Name == "" {
		return fmt.Errorf("%w: %s", ErrInvalidName, account.Name)
	}

	return a.querier.UpdateAccount(ctx, db.UpdateAccountParams{
		Name:  account.Name,
		Notes: account.Notes,
		Align: account.Align,
		ID:    account.ID,
	})
}

func (a *accountService) UpdateExpenseAccount(ctx context.Context, account db.ExpenseAccount) error {
	if account.Name == "" {
		return fmt.Errorf("%w: %s", ErrInvalidName, account.Name)
	}

	return a.querier.UpdateExpenseAccount(ctx, db.UpdateExpenseAccountParams{
		Name:  account.Name,
		Notes: account.Notes,
		ID:    account.ID,
	})
}

func (a *accountService) UpdateIncomeAccount(ctx context.Context, account db.IncomeAccount) error {
	if account.Name == "" {
		return fmt.Errorf("%w: %s", ErrInvalidName, account.Name)
	}

	return a.querier.UpdateIncomeAccount(ctx, db.UpdateIncomeAccountParams{
		Name:  account.Name,
		Notes: account.Notes,
		ID:    account.ID,
	})
}

func (a *accountService) CreateAccount(ctx context.Context, name, notes string) error {
	if name == "" {
		return fmt.Errorf("%w: %s", ErrInvalidName, name)
	}

	return a.querier.InsertAccount(ctx, db.InsertAccountParams{
		Name:  name,
		Notes: &notes,
	})
}

func (a *accountService) CreateExpenseAccount(ctx context.Context, name, notes string) error {
	if name == "" {
		return fmt.Errorf("%w: %s", ErrInvalidName, name)
	}

	return a.querier.InsertExpenseAccount(ctx, db.InsertExpenseAccountParams{
		Name:  name,
		Notes: &notes,
	})
}

func (a *accountService) CreateIncomeAccount(ctx context.Context, name, notes string) error {
	if name == "" {
		return fmt.Errorf("%w: %s", ErrInvalidName, name)
	}

	return a.querier.InsertIncomeAccount(ctx, db.InsertIncomeAccountParams{
		Name:  name,
		Notes: &notes,
	})
}
