package repositories

import (
	"RoadToTribal2.0/internal/models"
	"context"
	"github.com/dranikpg/dto-mapper"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type DatabaseRepository struct {
	log *zap.SugaredLogger
	db  *bun.DB
}

//NewDatabaseRepository Creates a new repository
func NewDatabaseRepository(logger *zap.SugaredLogger, conn *bun.DB) *DatabaseRepository {
	return &DatabaseRepository{
		log: logger,
		db:  conn,
	}
}

func (d *DatabaseRepository) FindAll(ctx context.Context) (error, []models.TransactionModel) {
	var model = &[]models.TransactionModel{}

	_, err := d.db.NewSelect().
		Model(model).
		Exec(ctx)

	if err != nil {
		d.log.Errorf("Error while retrieving transactions: %s", err)
		return err, nil
	}

	d.log.Infof("Retrieving all transactions")
	return nil, *model
}

func (d *DatabaseRepository) FindById(ctx context.Context, transactionID string) (models.TransactionModel, bool) {
	var model = &models.TransactionModel{}

	err := d.db.NewSelect().
		Model(model).
		Where("id = ?", transactionID).
		Scan(ctx)

	if err != nil {
		return *model, false
	}

	return *model, true

}

func (d *DatabaseRepository) Create(ctx context.Context, transaction *models.TransactionModel) (error, *models.TransactionModel) {
	result, err := d.db.NewInsert().Model(transaction).Exec(ctx)
	if err != nil {
		return err, nil
	}

	var transactionModel *models.TransactionModel
	dto.Map(transactionModel, result)

	return nil, transactionModel
}
