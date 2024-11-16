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

func TestAccount_GetAccounts(t *testing.T) {
	t.Run("without errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		expectedAccounts := []db.Account{{ID: 1}, {ID: 2}}
		querier.On("GetAccounts", mock.Anything).Return(expectedAccounts, nil).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetAccounts(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedAccounts, accounts)
	})

	t.Run("with errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		querier.On("GetAccounts", mock.Anything).Return(nil, errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetAccounts(context.Background())

		require.Error(t, err)
		require.Nil(t, accounts)
	})
}

func TestAccount_GetExpenseAccounts(t *testing.T) {
	t.Run("without errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		expectedAccounts := []db.ExpenseAccount{{ID: 1}, {ID: 2}}
		querier.On("GetExpenseAccounts", mock.Anything).Return(expectedAccounts, nil).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetExpenseAccounts(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedAccounts, accounts)
	})

	t.Run("with errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		querier.On("GetExpenseAccounts", mock.Anything).Return(nil, errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetExpenseAccounts(context.Background())

		require.Error(t, err)
		require.Nil(t, accounts)
	})
}

func TestAccount_GetIncomeAccounts(t *testing.T) {
	t.Run("without errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		expectedAccounts := []db.IncomeAccount{{ID: 1}, {ID: 2}}
		querier.On("GetIncomeAccounts", mock.Anything).Return(expectedAccounts, nil).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetIncomeAccounts(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedAccounts, accounts)
	})

	t.Run("with errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		querier.On("GetIncomeAccounts", mock.Anything).Return(nil, errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		accounts, err := accountService.GetIncomeAccounts(context.Background())

		require.Error(t, err)
		require.Nil(t, accounts)
	})
}

func TestAccount_UpdateAccount(t *testing.T) {
	t.Run("without erros", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.Account{ID: 1, Name: "new name"}
		querier.On("UpdateAccount", mock.Anything, db.UpdateAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(nil).Once()
		defer querier.AssertExpectations(t)

		err := accountService.UpdateAccount(context.Background(), toUpdate)

		require.NoError(t, err)
	})

	t.Run("with validation errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.Account{ID: 1, Name: ""}
		defer querier.AssertExpectations(t)

		err := accountService.UpdateAccount(context.Background(), toUpdate)

		require.ErrorIs(t, err, account.ErrInvalidName)
	})

	t.Run("with database errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.Account{ID: 1, Name: "new name"}
		querier.On("UpdateAccount", mock.Anything, db.UpdateAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		err := accountService.UpdateAccount(context.Background(), toUpdate)

		require.Error(t, err)
	})
}

func TestIncomeAccount_UpdateIncomeAccount(t *testing.T) {
	t.Run("without erros", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.IncomeAccount{ID: 1, Name: "new name"}
		querier.On("UpdateIncomeAccount", mock.Anything, db.UpdateIncomeAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(nil).Once()
		defer querier.AssertExpectations(t)

		err := accountService.UpdateIncomeAccount(context.Background(), toUpdate)

		require.NoError(t, err)
	})

	t.Run("with validation errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.IncomeAccount{ID: 1, Name: ""}
		defer querier.AssertExpectations(t)

		err := accountService.UpdateIncomeAccount(context.Background(), toUpdate)

		require.ErrorIs(t, err, account.ErrInvalidName)
	})

	t.Run("with database errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.IncomeAccount{ID: 1, Name: "new name"}
		querier.On("UpdateIncomeAccount", mock.Anything, db.UpdateIncomeAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		err := accountService.UpdateIncomeAccount(context.Background(), toUpdate)

		require.Error(t, err)
	})
}

func TestExpenseAccount_UpdateExpenseAccount(t *testing.T) {
	t.Run("without erros", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.ExpenseAccount{ID: 1, Name: "new name"}
		querier.On("UpdateExpenseAccount", mock.Anything, db.UpdateExpenseAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(nil).Once()
		defer querier.AssertExpectations(t)

		err := accountService.UpdateExpenseAccount(context.Background(), toUpdate)

		require.NoError(t, err)
	})

	t.Run("with validation errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.ExpenseAccount{ID: 1, Name: ""}
		defer querier.AssertExpectations(t)

		err := accountService.UpdateExpenseAccount(context.Background(), toUpdate)

		require.ErrorIs(t, err, account.ErrInvalidName)
	})

	t.Run("with database errors", func(t *testing.T) {
		t.Parallel()

		querier := db.NewMockQuerier(t)
		accountService := account.New(querier)

		toUpdate := db.ExpenseAccount{ID: 1, Name: "new name"}
		querier.On("UpdateExpenseAccount", mock.Anything, db.UpdateExpenseAccountParams{Name: toUpdate.Name, ID: toUpdate.ID}).Return(errors.New("kaboom")).Once()
		defer querier.AssertExpectations(t)

		err := accountService.UpdateExpenseAccount(context.Background(), toUpdate)

		require.Error(t, err)
	})
}
