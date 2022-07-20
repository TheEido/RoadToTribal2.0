package Transaction

import (
	"errors"
	"strconv"
)

func GetAllTransaction() (error, Transactions) {
	return nil, transactions
}

func GetTransactionById(ID string) (error, *Transaction) {
	var transId = false
	id, _ := strconv.Atoi(ID)
	var trans *Transaction
	for _, value := range transactions {
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

func Addtransction(t *Transaction) (error, Transactions) {
	transactions := append(Transactions{}, *t)
	return nil, transactions
}
