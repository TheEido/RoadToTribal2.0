package models

type Transaction struct {
	ID        int     `json:"ID"`
	Amount    float64 `json:"Amount"`
	Currency  string  `json:"Currency"`
	CreatedAt string  `json:"CreatedAt"`
}

type Transactions []Transaction

var transactions = Transactions{
	Transaction{
		ID:        1,
		Amount:    50.47,
		Currency:  "USD",
		CreatedAt: "2022-07-19T13:02:01.440618Z",
	},
	Transaction{
		ID:        2,
		Amount:    8000.7865,
		Currency:  "MXN",
		CreatedAt: "2022-07-19T13:02:01.440618Z",
	},
	Transaction{
		ID:        3,
		Amount:    60789.674,
		Currency:  "EGP",
		CreatedAt: "2022-07-19T13:02:01.440618Z",
	},
}
