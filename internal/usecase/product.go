package usecase

import (
	product_grpc "clean-arch-go-grpc/internal/delivery/grpc/proto"
	"clean-arch-go-grpc/internal/entity"
	"clean-arch-go-grpc/internal/repository"
	"context"

	"github.com/sirupsen/logrus"
)

type productUsecase struct {
	log               *logrus.Logger
	productRepository repository.IProductRepository
}

type IProductUsecase interface {
	Gets(ctx context.Context) ([]*entity.Product, error)
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	GetByID(ctx context.Context, id string) (*product_grpc.Product, error)
}

func NewProductUsecase(
	log *logrus.Logger,
	productRepository repository.IProductRepository,
) IProductUsecase {
	return &productUsecase{
		log,
		productRepository,
	}
}

func (usecase *productUsecase) Gets(ctx context.Context) ([]*entity.Product, error) {
	products, err := usecase.productRepository.Gets(ctx)
	if err != nil {
		usecase.log.Errorf("error %v", err)
		return nil, err
	}

	return products, nil
}

func (usecase *productUsecase) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	product, err := usecase.productRepository.Create(ctx, product)
	if err != nil {
		usecase.log.Errorf("error %v", err)
		return nil, err
	}

	return product, nil
}

func (usecase *productUsecase) GetByID(ctx context.Context, id string) (*product_grpc.Product, error) {
	product, err := usecase.GetByID(ctx, id)
	if err != nil {
		usecase.log.Errorf("error %v", err)
		return nil, err
	}

	return product, err
}
