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
	GetAll(ctx context.Context) ([]db.Account, error)
	Update(ctx context.Context, account db.Account) error
}

type accountService struct {
	querier db.Querier
}

func New(querier db.Querier) *accountService {
	return &accountService{
		querier: querier,
	}
}

func (a *accountService) GetAll(ctx context.Context) ([]db.Account, error) {
	return a.querier.GetAccounts(ctx)
}

func (a *accountService) Update(ctx context.Context, account db.Account) error {
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
