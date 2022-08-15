package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uptrace/bun"
)

type TransactionModel struct {
	bun.BaseModel `bun:"table:transaction"`
	ID            uuid.UUID        `bun:"id,notnull,pk,type:uuid,default:gen_random_uuid()"`
	Amount        int64            `bun:"amount,notnull"`
	Currency      string           `bun:"currency,notnull"`
	CreatedAt     pgtype.Timestamp `bun:"createdAt,notnull"`
}

type CreateTransactionRequest struct {
	Amount    int64            `json:"amount" validate:"notnull,required"`
	Currency  string           `json:"currency" validate:"notnull,required"`
	CreatedAt pgtype.Timestamp `json:"createdAt" validate:"notnull,required"`
}

type TransactionDetailsResponse struct {
	ID        string           `json:"id"`
	Amount    string           `json:"amount"`
	Currency  string           `json:"currency"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
}
