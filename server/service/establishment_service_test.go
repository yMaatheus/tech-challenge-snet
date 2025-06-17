package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yMaatheus/tech-challenge-snet/model"
)

// Mock para o EstablishmentRepository
type mockRepo struct {
	hasStoresResult bool
	hasStoresErr    error
	deleteCalled    bool
}

func (m *mockRepo) Create(ctx context.Context, e *model.Establishment) error   { return nil }
func (m *mockRepo) FindAll(ctx context.Context) ([]model.Establishment, error) { return nil, nil }
func (m *mockRepo) FindByID(ctx context.Context, id int64) (*model.Establishment, error) {
	return nil, nil
}
func (m *mockRepo) Update(ctx context.Context, e *model.Establishment) error { return nil }
func (m *mockRepo) Delete(ctx context.Context, id int64) error {
	m.deleteCalled = true
	return nil
}
func (m *mockRepo) HasStores(ctx context.Context, id int64) (bool, error) {
	return m.hasStoresResult, m.hasStoresErr
}

func TestDeleteEstablishment_WhenNoStores_ShouldDelete(t *testing.T) {
	repo := &mockRepo{hasStoresResult: false}
	service := NewEstablishmentService(repo)

	err := service.Delete(context.Background(), 1)
	assert.NoError(t, err)
	assert.True(t, repo.deleteCalled, "Delete should be called")
}

func TestDeleteEstablishment_WhenHasStores_ShouldReturnError(t *testing.T) {
	repo := &mockRepo{hasStoresResult: true}
	service := NewEstablishmentService(repo)

	err := service.Delete(context.Background(), 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "cannot delete establishment: it has related stores")
}

func TestDeleteEstablishment_WhenHasStoresReturnsError(t *testing.T) {
	repo := &mockRepo{hasStoresErr: errors.New("db error")}
	service := NewEstablishmentService(repo)

	err := service.Delete(context.Background(), 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
}
