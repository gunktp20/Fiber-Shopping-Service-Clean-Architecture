package repository

import (
	"errors"
	productDto "shopping-service-be/modules/product/productDto"
	productEntity "shopping-service-be/modules/product/productEntity"

	"gorm.io/gorm"
)

type productGormRepository struct {
	db *gorm.DB
}

func NewProductGormRepository(db *gorm.DB) ProductRepositoryService {
	return &productGormRepository{db}
}

func (r *productGormRepository) CreateOneProduct(createProductReq *productDto.CreateProductReq) (*productEntity.Product, error) {
	return &productEntity.Product{}, errors.New("")

}

func (r *productGormRepository) GetAllProducts(keyword string) ([]*productEntity.Product, error) {
	return []*productEntity.Product{}, nil
}

func (r *productGormRepository) DeleteProductById(id string) error {
	return nil
}

func (r *productGormRepository) UpdateProductById(id string, updateProductReq *productDto.UpdateProductReq) (*productEntity.Product, error) {
	return &productEntity.Product{}, nil
}
