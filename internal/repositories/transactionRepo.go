package repositories

import (
	"RoadToTribal2.0/internal/models"
	"errors"
	"strconv"
)

func GetAllTransaction() (error, models.Transactions) {
	return nil, models.Transactions{}
}

func GetTransactionById(ID string) (error, *models.Transaction) {
	var transId = false
	id, _ := strconv.Atoi(ID)
	var trans *models.Transaction
	ts := models.Transactions{}
	for _, value := range ts {
		if value.ID == id {
			transId = true
			trans = &value
		}
	}
	if transId == false {
		return errors.New("this transaction is not exist"), nil
	} else {
		return nil, trans
	}

}

func Addtransction(t *models.Transaction) (error, models.Transactions) {
	transactions := append(models.Transactions{}, *t)
	return nil, transactions
}
