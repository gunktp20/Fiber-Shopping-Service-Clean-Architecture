package usecase

import (
	productDto "shopping-service-be/modules/product/productDto"
	productEntity "shopping-service-be/modules/product/productEntity"
)

func (r *productUsecase) CreateProductJson(createProductReq *productDto.CreateProductReq) (*productEntity.Product, error) {
	return r.productRepo.CreateOneProduct(createProductReq)
}
