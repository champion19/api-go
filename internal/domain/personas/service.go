package personas

import (
	"context"
	"ejemplo/internal/domain/model"
	"ejemplo/internal/repository/persona"
)

type Service interface {
	GetAll(ctx context.Context) ([]model.Persona, error)
	GetById(ctx context.Context, id int64) (model.Persona, error)
}

type service struct {
	repository persona.Repository
}

func NewService(repository persona.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(ctx context.Context) ([]model.Persona, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetById(ctx context.Context, id int64) (model.Persona, error) {
	return s.repository.GetById(ctx, id)
}