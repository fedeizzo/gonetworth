package account_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/fedeizzo/gonetworth/internal/db"
	"github.com/fedeizzo/gonetworth/internal/service/account"
)

func TestAccount_GetAll(t *testing.T) {
	t.Run("without errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		expectedAccounts := []db.Account{{ID: 1}, {ID: 2}}
		querier.On("GetAccounts", mock.Anything).Return(expectedAccounts, nil).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetAll(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedAccounts, accounts)
	})

	t.Run("with errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		querier.On("GetAccounts", mock.Anything).Return(nil, errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetAll(context.Background())

		require.Error(t, err)
		require.Nil(t, accounts)
	})
}

func TestAccount_Update(t *testing.T) {
	t.Run("without erros", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.Account{ID: 1, Name: "new name"}
		querier.On("UpdateAccount", mock.Anything, db.UpdateAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(nil).Once()
		defer querier.AssertExpectations(t)

		err := accountService.Update(context.Background(), toUpdate)

		require.NoError(t, err)
	})

	t.Run("with validation errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.Account{ID: 1, Name: ""}
		defer querier.AssertExpectations(t)

		err := accountService.Update(context.Background(), toUpdate)

		require.ErrorIs(t, err, account.ErrInvalidName)
	})

	t.Run("with database errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.Account{ID: 1, Name: "new name"}
		querier.On("UpdateAccount", mock.Anything, db.UpdateAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		err := accountService.Update(context.Background(), toUpdate)

		require.Error(t, err)
	})
}
