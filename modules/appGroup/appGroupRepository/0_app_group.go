package repository

import appGroupDto "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupDto"

type (
	AppGroupRepositoryService interface {
		CreateAppGroup(createAppGroupReq *appGroupDto.CreateAppGroupReq) (*appGroupDto.CreateAppGroupRes, error)
	}
)
