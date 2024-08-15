package model

import "time"

type Transaction struct {
	ID            uint      `json:"id"`
	CustomerId    uint      `json:"customer_id"`
	NomorKontrak  string    `json:"nomor_kontrak"`
	OTR           string    `json:"otr"`
	AdminFee      float32   `json:"admin_fee"`
	JumlahCicilan float32   `json:"jumlah_cicilan"`
	JumlahBunga   float32   `json:"jumlah_bunga"`
	NamaAsset     string    `json:"nama_asset"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
