package repository

import (
	"context"
	"database/sql"

	"github.com/yMaatheus/tech-challenge-snet/model"
)

// EstablishmentRepository defines methods for interacting with the establishments table.
type EstablishmentRepository interface {
	Create(ctx context.Context, e *model.Establishment) error
	FindAll(ctx context.Context) ([]model.Establishment, error)
	FindByID(ctx context.Context, id int64) (*model.Establishment, error)
	Update(ctx context.Context, e *model.Establishment) error
	Delete(ctx context.Context, id int64) error
	HasStores(ctx context.Context, id int64) (bool, error)
}

// establishmentRepository is a concrete implementation of EstablishmentRepository.
type establishmentRepository struct {
	db *sql.DB
}

func NewEstablishmentRepository(db *sql.DB) EstablishmentRepository {
	return &establishmentRepository{db}
}

func (r *establishmentRepository) Create(ctx context.Context, e *model.Establishment) error {
	query := `
        INSERT INTO establishments
            (number, name, corporate_name, address, city, state, zip_code, address_number)
        VALUES
            ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
    `
	return r.db.QueryRowContext(ctx, query,
		e.Number, e.Name, e.CorporateName, e.Address,
		e.City, e.State, e.ZipCode, e.AddressNumber,
	).Scan(&e.ID)
}

func (r *establishmentRepository) FindAll(ctx context.Context) ([]model.Establishment, error) {
	query := `SELECT id, number, name, corporate_name, address, city, state, zip_code, address_number FROM establishments`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var establishments []model.Establishment
	for rows.Next() {
		var e model.Establishment
		err := rows.Scan(&e.ID, &e.Number, &e.Name, &e.CorporateName, &e.Address, &e.City, &e.State, &e.ZipCode, &e.AddressNumber)
		if err != nil {
			return nil, err
		}
		establishments = append(establishments, e)
	}
	return establishments, nil
}

func (r *establishmentRepository) FindByID(ctx context.Context, id int64) (*model.Establishment, error) {
	query := `SELECT id, number, name, corporate_name, address, city, state, zip_code, address_number FROM establishments WHERE id = $1`
	var e model.Establishment
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&e.ID, &e.Number, &e.Name, &e.CorporateName, &e.Address, &e.City, &e.State, &e.ZipCode, &e.AddressNumber,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *establishmentRepository) Update(ctx context.Context, e *model.Establishment) error {
	query := `
        UPDATE establishments SET
            number = $1, name = $2, corporate_name = $3, address = $4,
            city = $5, state = $6, zip_code = $7, address_number = $8
        WHERE id = $9
    `
	_, err := r.db.ExecContext(ctx, query,
		e.Number, e.Name, e.CorporateName, e.Address, e.City,
		e.State, e.ZipCode, e.AddressNumber, e.ID,
	)
	return err
}

func (r *establishmentRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM establishments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *establishmentRepository) HasStores(ctx context.Context, id int64) (bool, error) {
	query := `SELECT COUNT(1) FROM stores WHERE establishment_id = $1`
	var count int
	err := r.db.QueryRowContext(ctx, query, id).Scan(&count)
	return count > 0, err
}
