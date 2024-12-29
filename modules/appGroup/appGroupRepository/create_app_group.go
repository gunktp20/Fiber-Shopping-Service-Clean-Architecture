package repository

import (
	"fmt"

	appGroupDto "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

func (r *appGroupGormRepository) CreateAppGroup(createAppGroupReq *appGroupDto.CreateAppGroupReq) (*appGroupDto.CreateAppGroupRes, error) {

	appGroup := models.AppGroup{
		Name:        createAppGroupReq.Name,
		Description: createAppGroupReq.Description,
		IconURL:     createAppGroupReq.IconURL,
	}

	if err := r.db.Create(&appGroup).Error; err != nil {
		return nil, fmt.Errorf("failed to create app group: %w", err)
	}

	return &appGroupDto.CreateAppGroupRes{
		ID:          appGroup.ID,
		Name:        appGroup.Name,
		Description: appGroup.Description,
		IconURL:     appGroup.IconURL,
		CreatedAt:   appGroup.CreatedAt.String(),
		UpdatedAt:   appGroup.UpdatedAt.String(),
	}, nil
}
