package usecase

func (r *productUsecase) DeleteProductById(productId string) error {
	return r.productRepo.DeleteProductById(productId)
}
