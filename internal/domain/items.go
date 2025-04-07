package domain

import "github.com/jackc/pgx/v5/pgtype"

type Item struct {
	ID            pgtype.UUID
	Name          string
	Description   string
	Type          string
	WeightInGrams int32
	Amount        pgtype.Int4
}
