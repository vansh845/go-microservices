package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId     uint64      `json:"order_id"`
	CustomerId  uuid.UUID   `json:"customer_id"`
	LineItems   []LineItems `json:"line_items"`
	CreatedAt   time.Time   `json:"created_at"`
	ShippedAt   time.Time   `json:"shipped_at"`
	CompletedAt time.Time   `json:"completed_at"`
}

type LineItems struct {
	ItemId   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
