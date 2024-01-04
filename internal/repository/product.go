package repository

import (
	"clean-arch-go-grpc/internal/entity"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

type IProductRepository interface {
	Gets(ctx context.Context) ([]*entity.Product, error)
	GetByID(ctx context.Context, id string) (*entity.Product, error)
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
}

func NewProductRepository(
	DB *gorm.DB,
	Log *logrus.Logger,
) IProductRepository {
	return &productRepository{
		DB,
		Log,
	}
}

func (repo *productRepository) Gets(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product

	if err := repo.DB.
		Find(&products).Error; err != nil {
		repo.Log.Errorf("error %v", err)
		return nil, err
	}

	return products, nil
}

func (repo *productRepository) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	var product *entity.Product

	if err := repo.DB.
		Where("id = ?", id).
		First(&product).Error; err != nil {
		repo.Log.Errorf("error %v", err)
		return nil, err
	}

	return nil, nil
}

func (repo *productRepository) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	if err := repo.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
