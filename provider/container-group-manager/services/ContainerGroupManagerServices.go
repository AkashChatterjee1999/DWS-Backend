package services

import (
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/container-group-manager/repository"
	"github.com/labstack/echo"
)

type ContainerManagerServices struct {
	repo   repository.ContainerGroupManagerRepoUsecases
	logger helpers.LogConfig
}

func NewContainerGroupManagerServices(repoInstance repository.ContainerGroupManagerRepoUsecases, logger helpers.LogConfig) ContainerGroupManagerServiceUsecases {
	return &ContainerManagerServices{repo: repoInstance, logger: logger}
}

func (c ContainerManagerServices) CreateContainerGroup(ctx echo.Context) {
	panic("implement me")
}

func (c ContainerManagerServices) UpdateContainerGroupConfig(ctx echo.Context) {
	panic("implement me")
}

func (c ContainerManagerServices) DeleteContainerGroup(ctx echo.Context) {
	panic("implement me")
}
