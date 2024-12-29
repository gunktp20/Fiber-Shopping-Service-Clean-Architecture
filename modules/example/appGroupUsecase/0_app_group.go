package usecase

import appGroupRepository "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupRepository"

type (
	AppGroupUsecaseService interface {
	}

	appGroupUsecase struct {
		appGroupRepo appGroupRepository.AppGroupRepositoryService
	}
)

func NewAppGroupUsecase(appGroupRepo appGroupRepository.AppGroupRepositoryService) AppGroupUsecaseService {
	return &appGroupUsecase{appGroupRepo: appGroupRepo}
}
