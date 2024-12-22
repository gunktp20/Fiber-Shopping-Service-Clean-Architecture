package usecase

import (
	productDto "shopping-service-be/modules/product/productDto"
	productEntity "shopping-service-be/modules/product/productEntity"
)

func (r *productUsecase) UpdateProductById(id string, updateProduct *productDto.UpdateProductReq) (*productEntity.Product, error) {
	return r.productRepo.UpdateProductById(id, updateProduct)
}
