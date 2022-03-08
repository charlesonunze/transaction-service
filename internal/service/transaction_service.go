package service

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	repo "github.com/charlesonunze/transaction-service/internal/db/repo"
	walletpb "github.com/charlesonunze/wallet-service/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TransactionService - interface for the transaction service
type TransactionService interface {
	CreditUser(ctx context.Context, userID, amount int64) (int64, error)
	DebitUser(ctx context.Context, userID, amount int64) (int64, error)
}

type transactionService struct {
	repo *repo.Queries
}

// New - returns an instance of the TransactionService
func New(repo *repo.Queries) TransactionService {
	return &transactionService{
		repo,
	}
}

var WALLET_CLIENT_PORT = os.Getenv("WALLET_CLIENT_PORT")

func (s *transactionService) CreditUser(ctx context.Context, userID, amount int64) (int64, error) {
	var balance int64
	transaction, err := s.repo.CreateTransaction(ctx, repo.CreateTransactionParams{
		UserID: userID,
		Amount: amount,
		Type:   CreditTransaction,
		Status: PendingTransaction,
	})
	if err != nil {
		if strings.Contains(err.Error(), "user_id") {
			return balance, errors.New("user not found")
		}
		return balance, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	clientConn, err := grpc.Dial(WALLET_CLIENT_PORT, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer clientConn.Close()

	walletClient := walletpb.NewWalletServiceClient(clientConn)

	res, err := walletClient.CreditUser(ctx, &walletpb.CreditUserRequest{
		UserId: userID,
		Amount: amount,
	})

	if err != nil {
		_, updateErr := s.repo.UpdateTransaction(ctx, repo.UpdateTransactionParams{
			ID:     transaction.ID,
			Status: FailedTransaction,
		})
		if updateErr != nil {
			return balance, updateErr
		}

		return balance, err
	}

	_, err = s.repo.UpdateTransaction(ctx, repo.UpdateTransactionParams{
		ID:     transaction.ID,
		Status: SuccessfulTransaction,
	})
	if err != nil {
		return balance, err
	}

	return res.Balance, nil
}

func (s *transactionService) DebitUser(ctx context.Context, userID, amount int64) (int64, error) {
	var balance int64
	transaction, err := s.repo.CreateTransaction(ctx, repo.CreateTransactionParams{
		UserID: userID,
		Amount: amount,
		Type:   CreditTransaction,
		Status: PendingTransaction,
	})
	if err != nil {
		if strings.Contains(err.Error(), "user_id") {
			return balance, errors.New("user not found")
		}
		return balance, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	clientConn, err := grpc.Dial(WALLET_CLIENT_PORT, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer clientConn.Close()
	walletClient := walletpb.NewWalletServiceClient(clientConn)

	res, err := walletClient.DebitUser(ctx, &walletpb.DebitUserRequest{
		UserId: userID,
		Amount: amount,
	})
	if err != nil {
		_, updateErr := s.repo.UpdateTransaction(ctx, repo.UpdateTransactionParams{
			ID:     transaction.ID,
			Status: FailedTransaction,
		})
		if updateErr != nil {
			return balance, updateErr
		}

		return balance, err
	}

	_, err = s.repo.UpdateTransaction(ctx, repo.UpdateTransactionParams{
		ID:     transaction.ID,
		Status: SuccessfulTransaction,
	})
	if err != nil {
		return balance, err
	}

	return res.Balance, nil
}
