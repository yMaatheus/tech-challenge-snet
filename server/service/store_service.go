package service

import (
	"context"

	"github.com/yMaatheus/tech-challenge-snet/model"
	"github.com/yMaatheus/tech-challenge-snet/repository"
)

type StoreService interface {
	Create(ctx context.Context, store *model.Store) error
	FindAll(ctx context.Context) ([]model.Store, error)
	FindByID(ctx context.Context, id int64) (*model.Store, error)
	Update(ctx context.Context, store *model.Store) error
	Delete(ctx context.Context, id int64) error
	FindByEstablishmentID(ctx context.Context, establishmentID int64) ([]model.Store, error)
}

type storeService struct {
	repo repository.StoreRepository
}

func NewStoreService(repo repository.StoreRepository) StoreService {
	return &storeService{repo}
}

func (s *storeService) Create(ctx context.Context, store *model.Store) error {
	return s.repo.Create(ctx, store)
}

func (s *storeService) FindAll(ctx context.Context) ([]model.Store, error) {
	return s.repo.FindAll(ctx)
}

func (s *storeService) FindByID(ctx context.Context, id int64) (*model.Store, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *storeService) Update(ctx context.Context, store *model.Store) error {
	return s.repo.Update(ctx, store)
}

func (s *storeService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *storeService) FindByEstablishmentID(ctx context.Context, establishmentID int64) ([]model.Store, error) {
	return s.repo.FindByEstablishmentID(ctx, establishmentID)
}
