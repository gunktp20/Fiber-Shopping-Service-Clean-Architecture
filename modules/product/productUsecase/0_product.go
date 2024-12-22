package usecase

import (
	productDto "shopping-service-be/modules/product/productDto"
	productEntity "shopping-service-be/modules/product/productEntity"
	repository "shopping-service-be/modules/product/productRepository"
)

type (
	ProductUsecaseService interface {
		CreateProduct(createProduct *productDto.CreateProductReq) (*productEntity.Product, error)
		GetAllProducts(keyword string) ([]*productEntity.Product, error)
		UpdateProductById(id string, updateProduct *productDto.UpdateProductReq) (*productEntity.Product, error)
		DeleteProductById(productId string) error

		CreateProductUrlencoded(createProduct *productDto.CreateProductReq) (*productEntity.Product, error)
		CreateProductJson(createProduct *productDto.CreateProductReq) (*productEntity.Product, error)
		CreateProductFormData(createProduct *productDto.CreateProductReq, fileBytes []byte) (*productEntity.Product, error)

		UploadProductImage(fileBytes []byte) (*productDto.UploadProductImageRes, error)
	}

	productUsecase struct {
		productRepo repository.ProductRepositoryService
	}
)

func NewProductUsecase(productRepo repository.ProductRepositoryService) ProductUsecaseService {
	return &productUsecase{productRepo: productRepo}
}
