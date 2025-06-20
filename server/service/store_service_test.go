package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yMaatheus/tech-challenge-snet/model"
)

type mockStoreRepo struct {
	CreateFn   func(ctx context.Context, store *model.Store) error
	FindAllFn  func(ctx context.Context) ([]model.Store, error)
	FindByIDFn func(ctx context.Context, id int64) (*model.Store, error)
	UpdateFn   func(ctx context.Context, store *model.Store) error
	DeleteFn   func(ctx context.Context, id int64) error
}

func (m *mockStoreRepo) Create(ctx context.Context, s *model.Store) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, s)
	}
	return nil
}
func (m *mockStoreRepo) FindAll(ctx context.Context) ([]model.Store, error) {
	if m.FindAllFn != nil {
		return m.FindAllFn(ctx)
	}
	return []model.Store{}, nil
}
func (m *mockStoreRepo) FindByID(ctx context.Context, id int64) (*model.Store, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	return nil, nil
}
func (m *mockStoreRepo) Update(ctx context.Context, s *model.Store) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, s)
	}
	return nil
}
func (m *mockStoreRepo) Delete(ctx context.Context, id int64) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

func TestStoreService_Create(t *testing.T) {
	repo := &mockStoreRepo{
		CreateFn: func(ctx context.Context, s *model.Store) error {
			s.ID = 99
			return nil
		},
	}
	service := NewStoreService(repo)
	store := &model.Store{Name: "Loja"}
	err := service.Create(context.Background(), store)
	assert.NoError(t, err)
	assert.Equal(t, int64(99), store.ID)
}

func TestStoreService_Create_Erro(t *testing.T) {
	repo := &mockStoreRepo{
		CreateFn: func(ctx context.Context, s *model.Store) error {
			return errors.New("erro ao criar")
		},
	}
	service := NewStoreService(repo)
	store := &model.Store{Name: "Loja"}
	err := service.Create(context.Background(), store)
	assert.Error(t, err)
}

func TestStoreService_FindAll(t *testing.T) {
	repo := &mockStoreRepo{
		FindAllFn: func(ctx context.Context) ([]model.Store, error) {
			return []model.Store{{ID: 1, Name: "Loja"}}, nil
		},
	}
	service := NewStoreService(repo)
	stores, err := service.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, stores, 1)
}

func TestStoreService_FindAll_Erro(t *testing.T) {
	repo := &mockStoreRepo{
		FindAllFn: func(ctx context.Context) ([]model.Store, error) {
			return nil, errors.New("erro find all")
		},
	}
	service := NewStoreService(repo)
	stores, err := service.FindAll(context.Background())
	assert.Error(t, err)
	assert.Nil(t, stores)
}

func TestStoreService_FindAll_Default(t *testing.T) {
	repo := &mockStoreRepo{}
	service := NewStoreService(repo)
	stores, err := service.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, stores, 0)
}

func TestStoreService_FindByID(t *testing.T) {
	repo := &mockStoreRepo{
		FindByIDFn: func(ctx context.Context, id int64) (*model.Store, error) {
			if id == 1 {
				return &model.Store{ID: 1, Name: "Loja"}, nil
			}
			return nil, nil
		},
	}
	service := NewStoreService(repo)
	store, err := service.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, store)
}

func TestStoreService_FindByID_NaoEncontrado(t *testing.T) {
	repo := &mockStoreRepo{
		FindByIDFn: func(ctx context.Context, id int64) (*model.Store, error) {
			return nil, nil
		},
	}
	service := NewStoreService(repo)
	store, err := service.FindByID(context.Background(), 2)
	assert.NoError(t, err)
	assert.Nil(t, store)
}

func TestStoreService_FindByID_Erro(t *testing.T) {
	repo := &mockStoreRepo{
		FindByIDFn: func(ctx context.Context, id int64) (*model.Store, error) {
			return nil, errors.New("erro ao buscar")
		},
	}
	service := NewStoreService(repo)
	store, err := service.FindByID(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, store)
}

func TestStoreService_Update(t *testing.T) {
	repo := &mockStoreRepo{
		UpdateFn: func(ctx context.Context, s *model.Store) error {
			return nil
		},
	}
	service := NewStoreService(repo)
	err := service.Update(context.Background(), &model.Store{})
	assert.NoError(t, err)
}

func TestStoreService_Update_Erro(t *testing.T) {
	repo := &mockStoreRepo{
		UpdateFn: func(ctx context.Context, s *model.Store) error {
			return errors.New("erro update")
		},
	}
	service := NewStoreService(repo)
	err := service.Update(context.Background(), &model.Store{})
	assert.Error(t, err)
}

func TestStoreService_Delete(t *testing.T) {
	repo := &mockStoreRepo{
		DeleteFn: func(ctx context.Context, id int64) error {
			return nil
		},
	}
	service := NewStoreService(repo)
	err := service.Delete(context.Background(), 1)
	assert.NoError(t, err)
}

func TestStoreService_Delete_Erro(t *testing.T) {
	repo := &mockStoreRepo{
		DeleteFn: func(ctx context.Context, id int64) error {
			return errors.New("erro delete")
		},
	}
	service := NewStoreService(repo)
	err := service.Delete(context.Background(), 1)
	assert.Error(t, err)
}
