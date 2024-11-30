package handler

import (
	"context"

	"github.com/fedeizzo/gonetworth/internal/app"
	"github.com/fedeizzo/gonetworth/internal/web/api"
)

type Accounts struct {
	app *app.App
}

func (a *Accounts) AccountsList(ctx context.Context, request api.AccountsListRequestObject) (api.AccountsListResponseObject, error) {

	return api.AccountsList200JSONResponse{}, nil
}
