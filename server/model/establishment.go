package model

// Establishment represents a business establishment
type Establishment struct {
	ID            int64  `json:"id"`
	Number        string `json:"number"`
	Name          string `json:"name"`
	CorporateName string `json:"corporate_name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zip_code"`
	AddressNumber string `json:"address_number"`
}
