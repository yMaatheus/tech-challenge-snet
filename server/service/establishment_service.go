package service

import (
	"context"
	"errors"

	"github.com/yMaatheus/tech-challenge-snet/model"
	"github.com/yMaatheus/tech-challenge-snet/repository"
)

// EstablishmentService defines business logic for establishments.
type EstablishmentService interface {
	Create(ctx context.Context, e *model.Establishment) error
	FindAll(ctx context.Context) ([]model.Establishment, error)
	FindByID(ctx context.Context, id int64) (*model.Establishment, error)
	Update(ctx context.Context, e *model.Establishment) error
	Delete(ctx context.Context, id int64) error
}

// establishmentService implements EstablishmentService.
type establishmentService struct {
	repo repository.EstablishmentRepository
}

func NewEstablishmentService(r repository.EstablishmentRepository) EstablishmentService {
	return &establishmentService{repo: r}
}

func (s *establishmentService) Create(ctx context.Context, e *model.Establishment) error {
	return s.repo.Create(ctx, e)
}

func (s *establishmentService) FindAll(ctx context.Context) ([]model.Establishment, error) {
	return s.repo.FindAll(ctx)
}

func (s *establishmentService) FindByID(ctx context.Context, id int64) (*model.Establishment, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *establishmentService) Update(ctx context.Context, e *model.Establishment) error {
	return s.repo.Update(ctx, e)
}

func (s *establishmentService) Delete(ctx context.Context, id int64) error {
	hasStores, err := s.repo.HasStores(ctx, id)
	if err != nil {
		return err
	}
	if hasStores {
		return errors.New("cannot delete establishment: it has related stores")
	}
	return s.repo.Delete(ctx, id)
}
