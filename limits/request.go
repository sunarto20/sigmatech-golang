package limits

type LimitRequest struct {
	Tenor  int     `json:"tenor"`
	Amount float64 `json:"amount"`
}
