package Transaction

import (
	"RoadToTribal2.0/internal/models"
	"context"
)

type ITransactionService interface {
	FindAllTransactions(ctx context.Context) ([]*models.TransactionDetailsResponse, bool)
	FindTransactionDetails(ctx context.Context, id string) (*models.TransactionDetailsResponse, bool)
	AddTransaction(ctx context.Context, request *models.CreateTransactionRequest) (*models.TransactionModel, bool)
}
