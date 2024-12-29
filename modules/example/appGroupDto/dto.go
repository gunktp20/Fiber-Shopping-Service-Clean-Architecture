package dto

type (
	CreateAppGroupReq struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		IconURL     string `json:"icon_url" validate:"required"`
	}

	CreateAppGroupRes struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		IconURL     string `json:"icon_url"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
)
