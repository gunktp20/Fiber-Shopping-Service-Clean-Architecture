package dto

import "time"

type (
	CreateProductReq struct {
		Name         string  `json:"name" validate:"required"`
		Description  string  `json:"description" validate:"required"`
		Price        float64 `json:"price" validate:"required"`
		ProductImage string  `json:"product_image"`
	}

	CreateProductRes struct {
		ID          int64     `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	GetAllProductsRes []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	UploadProductImageRes struct {
		FileName string `json:"file_name"`
		FilePath string `json:"file_path"`
		FileSize string `json:"file_size"`
	}

	UpdateProductReq struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}
)
