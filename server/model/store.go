package model

type Store struct {
	ID              int64  `json:"id"`
	Number          string `json:"number" validate:"required"`
	Name            string `json:"name" validate:"required"`
	CorporateName   string `json:"corporate_name"`
	Address         string `json:"address" validate:"required"`
	City            string `json:"city" validate:"required"`
	State           string `json:"state" validate:"required,len=2"`
	ZipCode         string `json:"zip_code" validate:"required"`
	AddressNumber   string `json:"address_number" validate:"required"`
	EstablishmentID int64  `json:"establishment_id" validate:"required"`
}
