package repository

import (
	productDto "shopping-service-be/modules/product/productDto"
	productEntity "shopping-service-be/modules/product/productEntity"
)

type (
	ProductRepositoryService interface {
		CreateOneProduct(createProduct *productDto.CreateProductReq) (*productEntity.Product, error)
		GetAllProducts(keyword string) ([]*productEntity.Product, error)
		DeleteProductById(id string) error
		UpdateProductById(id string, updateProductReq *productDto.UpdateProductReq) (*productEntity.Product, error)
	}
)
