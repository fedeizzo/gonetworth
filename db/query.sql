-- name: GetAccounts :many
SELECT * FROM accounts;

-- name: GetExpenseAccounts :many
SELECT * FROM expense_accounts;

-- name: GetIncomeAccounts :many
SELECT * FROM income_accounts;

-- name: UpdateAccount :exec
UPDATE accounts
SET
        name = $1,
        notes = $2,
        align = $3
WHERE id = $4;

-- name: InsertAccount :exec
INSERT INTO accounts(name, notes)
VALUES ($1, $2);

-- name: UpdateExpenseAccount :exec
UPDATE expense_accounts
SET
        name = $1,
        notes = $2
WHERE id = $3;

-- name: InsertExpenseAccount :exec
INSERT INTO expense_accounts(name, notes)
VALUES ($1, $2);

-- name: UpdateIncomeAccount :exec
UPDATE income_accounts
SET
        name = $1,
        notes = $2
WHERE id = $3;

-- name: InsertIncomeAccount :exec
INSERT INTO income_accounts(name, notes)
VALUES ($1, $2);

-- name: GetExpenseCategories :many
SELECT * FROM expense_categories;

-- name: GetIncomeCategories :many
SELECT * FROM income_categories;

-- name: GetTransferCategories :many
SELECT * FROM transfer_categories;

-- name: UpdateIncomeCategory :exec
UPDATE income_categories
SET
        name = $1,
        notes = $2
WHERE id = $3;

-- name: InsertIncomeCategory :exec
INSERT INTO income_categories(name, notes)
VALUES ($1, $2);

-- name: UpdateExpenseCategory :exec
UPDATE expense_categories
SET
        name = $1,
        notes = $2
WHERE id = $3;

-- name: InsertExpenseCategory :exec
INSERT INTO expense_categories(name, notes)
VALUES ($1, $2);

-- name: UpdateTransferCategory :exec
UPDATE transfer_categories
SET
        name = $1,
        notes = $2
WHERE id = $3;

-- name: InsertTransferCategory :exec
INSERT INTO transfer_categories(name, notes)
VALUES ($1, $2);

-- name: GetExpenses :exec
SELECT * FROM expenses;

-- name: GetIncomes :exec
SELECT * FROM incomes;

-- name: GetTransfers :exec
SELECT * FROM transfers;
