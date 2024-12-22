package usecase

import productEntity "shopping-service-be/modules/product/productEntity"

func (r *productUsecase) GetAllProducts(keyword string) ([]*productEntity.Product, error) {
	return r.productRepo.GetAllProducts(keyword)
}
