package usecase

import (
	"fmt"
	"os"
	productDto "shopping-service-be/modules/product/productDto"
	productEntity "shopping-service-be/modules/product/productEntity"
	"shopping-service-be/pkg/helpers"
	"time"
)

func (r *productUsecase) CreateProductFormData(createProductReq *productDto.CreateProductReq, fileBytes []byte) (*productEntity.Product, error) {
	// Get file extension from fileBytes
	fileExtension, err := helpers.GetImageFileExtension(fileBytes)
	if err != nil {
		return &productEntity.Product{}, err
	}

	// Generate a new file name
	newFileName := fmt.Sprintf("product_%d%s", time.Now().UnixNano(), fileExtension)
	filePath := fmt.Sprintf("./uploads/vendors/%s", newFileName)

	// Save the file to the server
	if err := os.WriteFile(filePath, fileBytes, 0644); err != nil {
		return &productEntity.Product{}, err
	}

	createProductReq.ProductImage = filePath

	return r.productRepo.CreateOneProduct(createProductReq)
}
