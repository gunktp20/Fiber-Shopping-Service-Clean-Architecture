package usecase

import (
	appGroupDto "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupDto"
	appGroupRepository "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupRepository"
)

type (
	AppGroupUsecaseService interface {
		CreateAppGroup(createAppGroupReq *appGroupDto.CreateAppGroupReq) (*appGroupDto.CreateAppGroupRes, error)
	}

	appGroupUsecase struct {
		appGroupRepo appGroupRepository.AppGroupRepositoryService
	}
)

func NewAppGroupUsecase(appGroupRepo appGroupRepository.AppGroupRepositoryService) AppGroupUsecaseService {
	return &appGroupUsecase{appGroupRepo: appGroupRepo}
}
