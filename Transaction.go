package Transaction

import "time"

type Transaction struct {
	ID        int       `json:"ID"`
	Amount    float64   `json:"Amount"`
	Currency  string    `json:"Currency"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type Transactions []Transaction

var transactions = Transactions{
	Transaction{
		ID:        1,
		Amount:    50.47,
		Currency:  "USD",
		CreatedAt: time.Now().UTC(),
	},
	Transaction{
		ID:        2,
		Amount:    8000.7865,
		Currency:  "MXN",
		CreatedAt: time.Now().UTC(),
	},
	Transaction{
		ID:        3,
		Amount:    60789.674,
		Currency:  "EGP",
		CreatedAt: time.Now().UTC(),
	},
}
