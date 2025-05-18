package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID uuid.UUID
	Name      string
	Brand     *string
	Category  *string
	Price     int64
	Rating    *float32
}

type Category struct {
	Name  string
	Label *string
}

type ProductCharacteristic struct {
	ProductID         uuid.UUID
	Description       *string
	Weight            *int32
	QuantityInPackage *int32
	ShelfLife         time.Duration
	StorageConditions *string
	Nutrition         *string
}
