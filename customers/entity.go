package customers

import "time"

type Customer struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	KTP          uint64    `gorm:"type:bigint" json:"ktp"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	Gaji         float32   `json:"gaji"`
	FotoKTP      string    `json:"foto_ktp"`
	FotoSelfie   string    `json:"foto_selfie"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"Updated_at"`
}
