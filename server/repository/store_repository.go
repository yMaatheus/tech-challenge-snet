package repository

import (
	"context"
	"database/sql"

	"github.com/yMaatheus/tech-challenge-snet/model"
)

type StoreRepository interface {
	Create(ctx context.Context, store *model.Store) error
	FindAll(ctx context.Context) ([]model.Store, error)
	FindByID(ctx context.Context, id int64) (*model.Store, error)
	Update(ctx context.Context, store *model.Store) error
	Delete(ctx context.Context, id int64) error
}

type storeRepository struct {
	db *sql.DB
}

func NewStoreRepository(db *sql.DB) StoreRepository {
	return &storeRepository{db}
}

func (r *storeRepository) Create(ctx context.Context, s *model.Store) error {
	query := `INSERT INTO stores (number, name, corporate_name, address, city, state, zip_code, address_number, establishment_id)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	return r.db.QueryRowContext(ctx, query, s.Number, s.Name, s.CorporateName, s.Address, s.City, s.State, s.ZipCode, s.AddressNumber, s.EstablishmentID).Scan(&s.ID)
}

func (r *storeRepository) FindAll(ctx context.Context) ([]model.Store, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, number, name, corporate_name, address, city, state, zip_code, address_number, establishment_id FROM stores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stores []model.Store
	for rows.Next() {
		var s model.Store
		if err := rows.Scan(&s.ID, &s.Number, &s.Name, &s.CorporateName, &s.Address, &s.City, &s.State, &s.ZipCode, &s.AddressNumber, &s.EstablishmentID); err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}
	return stores, nil
}

func (r *storeRepository) FindByID(ctx context.Context, id int64) (*model.Store, error) {
	var s model.Store
	err := r.db.QueryRowContext(ctx, "SELECT id, number, name, corporate_name, address, city, state, zip_code, address_number, establishment_id FROM stores WHERE id=$1", id).
		Scan(&s.ID, &s.Number, &s.Name, &s.CorporateName, &s.Address, &s.City, &s.State, &s.ZipCode, &s.AddressNumber, &s.EstablishmentID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *storeRepository) Update(ctx context.Context, s *model.Store) error {
	_, err := r.db.ExecContext(ctx, `UPDATE stores SET number=$1, name=$2, corporate_name=$3, address=$4, city=$5, state=$6, zip_code=$7, address_number=$8, establishment_id=$9 WHERE id=$10`,
		s.Number, s.Name, s.CorporateName, s.Address, s.City, s.State, s.ZipCode, s.AddressNumber, s.EstablishmentID, s.ID)
	return err
}

func (r *storeRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM stores WHERE id=$1", id)
	return err
}
