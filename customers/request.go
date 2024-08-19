package customers

type CustomerRequest struct {
	Name         string  `json:"name"`
	KTP          uint64  `json:"'ktp'"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	TempatLahir  string  `json:"tempat_lahir"`
	TanggalLahir string  `json:"tanggal_lahir"`
	Gaji         float32 `json:"gaji"`
	FotoKTP      string  `json:"foto_ktp"`
	FotoSelfie   string  `json:"foto_selfie"`
}
