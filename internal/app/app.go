package app

import (
	"github.com/fedeizzo/gonetworth/internal/db"
	"github.com/fedeizzo/gonetworth/internal/service/account"
)

type App struct {
	Querier db.Querier

	accountService account.AccountService
}

func (a App) GetAccountService() account.AccountService {
	if a.accountService == nil {
		a.accountService = account.New(a.Querier)
	}

	return a.accountService
}
