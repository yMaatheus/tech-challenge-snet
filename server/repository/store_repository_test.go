package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yMaatheus/tech-challenge-snet/model"
	"github.com/yMaatheus/tech-challenge-snet/testutil"
)

func TestStoreRepository_CRUD(t *testing.T) {
	db := testutil.GetTestDB(t)

	// Reset DB
	testutil.ExecSQLFile(t, db, "../database/reset.sql")
	testutil.ExecSQLFile(t, db, "../database/migration.sql")

	// Cria Establishment dummy para FK
	_, err := db.Exec(`
		INSERT INTO establishments (number, name, corporate_name, address, city, state, zip_code, address_number)
		VALUES ('E001', 'Est1', 'Corp', 'Rua', 'Cidade', 'ST', '12345678', '10')
	`)
	assert.NoError(t, err)

	repo := NewStoreRepository(db)
	ctx := context.Background()

	// Create
	store := &model.Store{
		Number:          "S001",
		Name:            "Loja Teste",
		CorporateName:   "Corp Store",
		Address:         "Rua A",
		City:            "City",
		State:           "ST",
		ZipCode:         "000",
		AddressNumber:   "10",
		EstablishmentID: 1, // agora existe
	}
	err = repo.Create(ctx, store)
	assert.NoError(t, err)
	assert.True(t, store.ID > 0)

	// FindAll
	stores, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.True(t, len(stores) >= 1)

	// FindByID
	got, err := repo.FindByID(ctx, store.ID)
	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, "Loja Teste", got.Name)

	// Update
	store.Name = "Atualizada"
	err = repo.Update(ctx, store)
	assert.NoError(t, err)
	got, _ = repo.FindByID(ctx, store.ID)
	assert.Equal(t, "Atualizada", got.Name)

	// Delete
	err = repo.Delete(ctx, store.ID)
	assert.NoError(t, err)
	got, _ = repo.FindByID(ctx, store.ID)
	assert.Nil(t, got)
}
