package limits

import "time"

type Limit struct {
	ID         uint      `json:"id"`
	CustomerId int       `json:"customer_id"`
	Tenor      int       `json:"tenor"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
