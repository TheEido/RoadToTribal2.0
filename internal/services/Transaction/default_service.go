package Transaction

import (
	"RoadToTribal2.0/internal/models"
	"RoadToTribal2.0/internal/repositories"
	"context"
	"github.com/dranikpg/dto-mapper"
	"go.uber.org/zap"
)

type DefaultTransactionService struct {
	log *zap.SugaredLogger
	db  repositories.IRepository
}

func NewDefaultTransactionService(logger *zap.SugaredLogger, transactionRepo repositories.IRepository) *DefaultTransactionService {
	return &DefaultTransactionService{
		log: logger,
		db:  transactionRepo,
	}
}

func (ds *DefaultTransactionService) FindAllTransactions(ctx context.Context) (*[]models.TransactionDetailsResponse, bool) {

	err, Result := ds.db.FindAll(ctx)
	if err != nil {
		ds.log.Errorf("Error while retrieving transactions: %s", err)
		return nil, false
	}

	var data = &[]models.TransactionDetailsResponse{}
	/*for _, value := range Result {
		data = append(data, &models.TransactionDetailsResponse{
			ID:                   value.ID,
			CustomerID:           value.CustomerID,
			Provider:             value.Provider,
			ProviderReference:    value.ProviderReference,
			AdditionalReferences: value.AdditionalReferences,
			CreatedAt:            value.CreatedAt,
		})
	}*/

	dto.Map(data, Result)
	ds.log.Infof("Retrieving all transactions")

	return data, true
}

func (ds *DefaultTransactionService) FindTransactionDetails(ctx context.Context, id string) (*models.TransactionDetailsResponse, bool) {

	result, ok := ds.db.FindById(ctx, id)
	if !ok {
		ds.log.Errorf("This transaction is not exist")
		return nil, false
	}
	ds.log.Infof("Finding details for transaction --> %s", id)
	var response *models.TransactionDetailsResponse
	dto.Map(response, result)

	return response, true
}

func (ds *DefaultTransactionService) AddTransaction(ctx context.Context, request *models.CreateTransactionRequest) (*models.TransactionModel, bool) {
	var trans = &models.TransactionModel{}
	dto.Map(trans, request)
	result, err := ds.db.Create(ctx, trans)

	if err != nil {
		ds.log.Errorf("Error with creating transaction")
		return nil, false
	}

	dto.Map(trans, result)
	ds.log.Infof("Transaction created!")
	return trans, true
}
