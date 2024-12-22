package usecase

import (
	"fmt"
	"os"
	productDto "shopping-service-be/modules/product/productDto"
	"shopping-service-be/pkg/helpers"
	"time"
)

func (r *productUsecase) UploadProductImage(fileBytes []byte) (*productDto.UploadProductImageRes, error) {

	// Get file extension from fileBytes
	fileExtension, err := helpers.GetImageFileExtension(fileBytes)
	if err != nil {
		return &productDto.UploadProductImageRes{}, err
	}

	// Generate a new file name
	newFileName := fmt.Sprintf("product_%d%s", time.Now().UnixNano(), fileExtension)
	filePath := fmt.Sprintf("./uploads/vendors/%s", newFileName)

	// Save the file to the server
	if err := os.WriteFile(filePath, fileBytes, 0644); err != nil {
		return &productDto.UploadProductImageRes{}, err
	}
	// Get file size
	fileSize := len(fileBytes)

	return &productDto.UploadProductImageRes{
		FileName: newFileName,
		FilePath: filePath,
		FileSize: helpers.HumanFileSize(int64(fileSize), true, 2),
	}, nil
}
