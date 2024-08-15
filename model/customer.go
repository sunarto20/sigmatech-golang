package model

import (
	"math/big"
	"time"
)

type Customer struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	KTP          big.Int   `json:"ktp"`
	FullName     string    `json:"fullName"`
	LegalName    string    `json:"legalName"`
	TempatLahir  string    `json:"tempatLahir"`
	TanggalLahir time.Time `json:"tanggalLahir"`
	Gaji         float32   `json:"gaji"`
	FotoKTP      string    `json:"fotoKTP"`
	FotoSelfie   string    `json:"fotoSelfie"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdateAt     time.Time `json:"updateDate"`
}
