package services

import "github.com/labstack/echo"

type ContainerGroupManagerServiceUsecases interface {
	CreateContainerGroup(ctx echo.Context)
	UpdateContainerGroupConfig(ctx echo.Context)
	DeleteContainerGroup(ctx echo.Context)
}
