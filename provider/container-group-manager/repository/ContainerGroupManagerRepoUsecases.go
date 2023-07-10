package repository

import "github.com/labstack/echo"

type ContainerGroupManagerRepoUsecases interface {
	CreateContainerGroup(ctx echo.Context)
	UpdateContainerGroupConfig(ctx echo.Context)
	DeleteContainerGroup(ctx echo.Context)
}
