package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T, userID int64) Transaction {
	arg := CreateTransactionParams{
		UserID: userID,
		Amount: int64(gofakeit.Number(1000, 10000)),
		Type:   TransactionType(gofakeit.RandomString([]string{"CREDIT", "DEBIT"})),
		Status: TransactionStatus(gofakeit.RandomString([]string{"PENDING", "FAILED", "SUCCESSFUL"})),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, transaction.UserID, arg.UserID)
	require.Equal(t, transaction.Amount, arg.Amount)
	require.Equal(t, transaction.Type, arg.Type)
	require.Equal(t, transaction.Status, arg.Status)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	user := createRandomUser(t)
	createRandomTransaction(t, user.ID)
}

func TestGetTransaction(t *testing.T) {
	user := createRandomUser(t)
	transaction1 := createRandomTransaction(t, user.ID)

	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.UserID, transaction2.UserID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.Type, transaction2.Type)
	require.Equal(t, transaction1.Status, transaction2.Status)
	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)
}

func TestUpdateTransaction(t *testing.T) {
	user := createRandomUser(t)
	transaction1 := createRandomTransaction(t, user.ID)
	transaction2 := createRandomTransaction(t, user.ID)

	arg := UpdateTransactionParams{
		ID:     transaction1.ID,
		Status: TransactionStatus(gofakeit.RandomString([]string{"PENDING", "FAILED", "SUCCESSFUL"})),
	}

	updatedTransaction1, err := testQueries.UpdateTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedTransaction1)

	require.Equal(t, transaction1.UserID, updatedTransaction1.UserID)
	require.Equal(t, transaction1.Amount, updatedTransaction1.Amount)
	require.Equal(t, transaction1.Type, updatedTransaction1.Type)
	require.Equal(t, arg.Status, updatedTransaction1.Status)
	require.WithinDuration(t, transaction1.CreatedAt, updatedTransaction1.CreatedAt, time.Second)

	arg2 := UpdateTransactionParams{
		ID:     transaction2.ID,
		Status: "Random string",
	}

	updatedTransaction2, err := testQueries.UpdateTransaction(context.Background(), arg2)

	require.Error(t, err)
	require.Empty(t, updatedTransaction2)
}
