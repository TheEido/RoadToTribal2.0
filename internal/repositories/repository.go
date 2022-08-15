package repositories

import (
	"RoadToTribal2.0/internal/models"
	"context"
)

type IRepository interface {
	FindAll(ctx context.Context) (error, []models.TransactionModel)
	FindById(ctx context.Context, transactionID string) (models.TransactionModel, bool)
	Create(ctx context.Context, transaction *models.TransactionModel) (error, *models.TransactionModel)
}
