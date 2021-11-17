package product

import (
	"context"

	"github.com/agustinrabini/Gocker/internal/domain"
)

type Service interface {
	Get(ctx context.Context, id int) (domain.Product, error)
	GetAll(ctx context.Context) (domain.Products, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, id int, p domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) (domain.Products, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, p domain.Product) (domain.Product, error) {

	id, err := s.repository.Save(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}

	p.Id_product = id

	return p, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) Update(ctx context.Context, id int, p domain.Product) (domain.Product, error) {

	err := s.repository.Update(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}
