package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yMaatheus/tech-challenge-snet/model"
)

type mockRepo struct {
	hasStoresResult              bool
	hasStoresErr                 error
	deleteCalled                 bool
	findByIDResult               *model.Establishment
	findByIDErr                  error
	findAllWithStoresTotalResult []model.EstablishmentWithStoresTotal
	findAllWithStoresTotalErr    error
	findAllWithStoresTotalCalled bool
	findStoresResult             []model.Store
	findStoresErr                error
}

func (m *mockRepo) Create(ctx context.Context, e *model.Establishment) error {
	if e.Name == "erro" {
		return errors.New("erro ao criar")
	}
	e.ID = 9
	return nil
}
func (m *mockRepo) FindAll(ctx context.Context) ([]model.Establishment, error) { return nil, nil }
func (m *mockRepo) FindAllWithStoresTotal(ctx context.Context) ([]model.EstablishmentWithStoresTotal, error) {
	m.findAllWithStoresTotalCalled = true
	if m.findAllWithStoresTotalErr != nil {
		return nil, m.findAllWithStoresTotalErr
	}
	if m.findAllWithStoresTotalResult != nil {
		return m.findAllWithStoresTotalResult, nil
	}
	return []model.EstablishmentWithStoresTotal{
		{
			ID:            1,
			Number:        "E001",
			Name:          "Test",
			CorporateName: "Corp Test",
			Address:       "Rua Teste",
			AddressNumber: "10",
			City:          "Cidade Teste",
			State:         "ST",
			ZipCode:       "12345678",
			StoresTotal:   2,
		},
	}, nil
}
func (m *mockRepo) FindByID(ctx context.Context, id int64) (*model.Establishment, error) {
	return m.findByIDResult, m.findByIDErr
}
func (m *mockRepo) Update(ctx context.Context, e *model.Establishment) error {
	if e.Name == "erro" {
		return errors.New("erro ao atualizar")
	}
	return nil
}
func (m *mockRepo) Delete(ctx context.Context, id int64) error {
	m.deleteCalled = true
	return nil
}
func (m *mockRepo) HasStores(ctx context.Context, id int64) (bool, error) {
	return m.hasStoresResult, m.hasStoresErr
}
func (m *mockRepo) FindStoresByEstablishmentID(ctx context.Context, establishmentID int64) ([]model.Store, error) {
	return m.findStoresResult, m.findStoresErr
}

func TestEstablishmentService_Create(t *testing.T) {
	repo := &mockRepo{}
	service := NewEstablishmentService(repo)

	est := &model.Establishment{Name: "Loja"}
	err := service.Create(context.Background(), est)
	assert.NoError(t, err)
	assert.Equal(t, int64(9), est.ID)

	est2 := &model.Establishment{Name: "erro"}
	err2 := service.Create(context.Background(), est2)
	assert.Error(t, err2)
}

func TestEstablishmentService_Update(t *testing.T) {
	repo := &mockRepo{}
	service := NewEstablishmentService(repo)

	est := &model.Establishment{Name: "Loja"}
	err := service.Update(context.Background(), est)
	assert.NoError(t, err)

	est2 := &model.Establishment{Name: "erro"}
	err2 := service.Update(context.Background(), est2)
	assert.Error(t, err2)
}

func TestEstablishmentService_FindAll(t *testing.T) {
	repo := &mockRepo{}
	service := NewEstablishmentService(repo)

	list, err := service.FindAll(context.Background())
	assert.NoError(t, err)
	assert.True(t, repo.findAllWithStoresTotalCalled)
	assert.Len(t, list, 1)
	assert.Equal(t, int64(1), list[0].ID)
}

func TestEstablishmentService_FindAll_ErroNoRepo(t *testing.T) {
	repo := &mockRepo{findAllWithStoresTotalErr: errors.New("erro repo")}
	service := NewEstablishmentService(repo)

	list, err := service.FindAll(context.Background())
	assert.Error(t, err)
	assert.Nil(t, list)
}

func TestEstablishmentService_FindByID_Sucesso(t *testing.T) {
	repo := &mockRepo{
		findByIDResult: &model.Establishment{
			ID:            2,
			Number:        "X",
			Name:          "teste",
			CorporateName: "Corp",
			Address:       "Rua",
			City:          "C",
			State:         "SP",
			ZipCode:       "123",
			AddressNumber: "99",
		},
		findStoresResult: []model.Store{
			{ID: 1, Name: "Loja A"},
		},
	}
	service := NewEstablishmentService(repo)
	est, err := service.FindByID(context.Background(), 2)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), est.ID)
	assert.Len(t, est.Stores, 1)
}

func TestEstablishmentService_FindByID_NotFound(t *testing.T) {
	repo := &mockRepo{
		findByIDResult: nil,
		findByIDErr:    nil,
	}
	service := NewEstablishmentService(repo)
	est, err := service.FindByID(context.Background(), 123)
	assert.Error(t, err)
	assert.Nil(t, est)
}

func TestEstablishmentService_FindByID_ErroNoRepo(t *testing.T) {
	repo := &mockRepo{
		findByIDErr: errors.New("falha repo"),
	}
	service := NewEstablishmentService(repo)
	est, err := service.FindByID(context.Background(), 3)
	assert.Error(t, err)
	assert.Nil(t, est)
}

func TestEstablishmentService_FindByID_ErroStores(t *testing.T) {
	repo := &mockRepo{
		findByIDResult: &model.Establishment{ID: 1},
		findStoresErr:  errors.New("erro stores"),
	}
	service := NewEstablishmentService(repo)
	est, err := service.FindByID(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, est)
}

func TestEstablishmentService_Delete_QuandoNaoTemStores_DeveDeletar(t *testing.T) {
	repo := &mockRepo{hasStoresResult: false}
	service := NewEstablishmentService(repo)

	err := service.Delete(context.Background(), 1)
	assert.NoError(t, err)
	assert.True(t, repo.deleteCalled, "Delete deve ser chamado")
}

func TestEstablishmentService_Delete_QuandoTemStores_DeveRetornarErro(t *testing.T) {
	repo := &mockRepo{hasStoresResult: true}
	service := NewEstablishmentService(repo)

	err := service.Delete(context.Background(), 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "cannot delete establishment: it has related stores")
}

func TestEstablishmentService_Delete_HasStoresRetornaErro(t *testing.T) {
	repo := &mockRepo{hasStoresErr: errors.New("db error")}
	service := NewEstablishmentService(repo)

	err := service.Delete(context.Background(), 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
}
