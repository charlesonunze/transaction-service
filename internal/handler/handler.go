package handler

import (
	"context"

	repo "github.com/charlesonunze/transaction-service/internal/db/repo"
	"github.com/charlesonunze/transaction-service/internal/service"
	transactionpb "github.com/charlesonunze/transaction-service/pb/v1"
)

type server struct {
	repo *repo.Queries
}

// New - returns an instance of the TransferServiceServer
func New(repo *repo.Queries) transactionpb.TransferServiceServer {
	return &server{
		repo,
	}
}

func (s *server) GetService() service.TransactionService {
	return service.New(s.repo)
}

func (s *server) CreditUser(ctx context.Context, req *transactionpb.CreditUserRequest) (*transactionpb.CreditUserResponse, error) {
	var res transactionpb.CreditUserResponse
	err := req.Validate()
	if err != nil {
		return &res, err
	}

	svc := s.GetService()
	balance, err := svc.CreditUser(ctx, req.UserId, req.Amount)
	if err != nil {
		return &res, err
	}

	res.Balance = balance

	return &res, nil
}

func (s *server) DebitUser(ctx context.Context, req *transactionpb.DebitUserRequest) (*transactionpb.DebitUserResponse, error) {
	var res transactionpb.DebitUserResponse
	err := req.Validate()
	if err != nil {
		return &res, err
	}

	svc := s.GetService()
	balance, err := svc.DebitUser(ctx, req.UserId, req.Amount)
	if err != nil {
		return &res, err
	}

	res.Balance = balance

	return &res, nil
}
