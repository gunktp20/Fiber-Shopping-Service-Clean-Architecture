package usecase

import (
	appGroupDto "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupDto"
)

func (u *appGroupUsecase) CreateAppGroup(createAppGroupReq *appGroupDto.CreateAppGroupReq) (*appGroupDto.CreateAppGroupRes, error) {
	return u.appGroupRepo.CreateAppGroup(createAppGroupReq)
}
