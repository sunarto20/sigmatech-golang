package model

import "time"

type Tenor struct {
	ID         int       `json:"id"`
	CustomerId int       `json:"customer_id"`
	tenor      int       `json:"tenor"`
	amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
