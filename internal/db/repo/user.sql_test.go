package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	name := gofakeit.Name()
	user, err := testQueries.CreateUser(context.Background(), name)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
