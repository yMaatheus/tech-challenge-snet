package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yMaatheus/tech-challenge-snet/model"
	"github.com/yMaatheus/tech-challenge-snet/testutil"
)

func resetDB(t *testing.T, db *sql.DB) {
	testutil.ExecSQLFile(t, db, "../database/reset.sql")
	testutil.ExecSQLFile(t, db, "../database/migration.sql")
}

func TestEstablishmentRepository_CRUD(t *testing.T) {
	db := testutil.GetTestDB(t)
	resetDB(t, db)
	repo := NewEstablishmentRepository(db)
	ctx := context.Background()

	// Create
	est := &model.Establishment{
		Number:        "T123",
		Name:          "Test Establishment",
		CorporateName: "Test Corp",
		Address:       "123 Road",
		City:          "Testville",
		State:         "TS",
		ZipCode:       "99999-999",
		AddressNumber: "42",
	}
	err := repo.Create(ctx, est)
	assert.NoError(t, err)
	assert.NotZero(t, est.ID)

	// FindByID
	found, err := repo.FindByID(ctx, est.ID)
	assert.NoError(t, err)
	assert.Equal(t, est.Name, found.Name)

	// Update
	est.Name = "Updated Name"
	err = repo.Update(ctx, est)
	assert.NoError(t, err)
	updated, _ := repo.FindByID(ctx, est.ID)
	assert.Equal(t, "Updated Name", updated.Name)

	// FindAll
	list, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.True(t, len(list) > 0)

	// Delete
	err = repo.Delete(ctx, est.ID)
	assert.NoError(t, err)
	notFound, err := repo.FindByID(ctx, est.ID)
	assert.NoError(t, err)
	assert.Nil(t, notFound)
}
