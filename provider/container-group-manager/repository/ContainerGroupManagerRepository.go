package repository

import (
	"fmt"
	"github.com/DWS-Backend/provider/common/exceptions"
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/container-group-manager/class"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"time"
)

type ContainerGroupManagerRepository struct {
	db     *gorm.DB
	logger helpers.LogConfig
}

func NewContainerGroupManagerRepository(conn *gorm.DB, logPtr helpers.LogConfig) ContainerGroupManagerRepoUsecases {
	return &ContainerGroupManagerRepository{db: conn, logger: logPtr}
}

func (cgm *ContainerGroupManagerRepository) executeQuery(ctx echo.Context, query string, resultType interface{}) (
	response interface{}, err error) {
	cgm.logger.Info(ctx, fmt.Sprintf("Executing Query: %s", query))
	hasError := cgm.db.Raw(query).Scan(&resultType).Error
	return resultType, hasError
}

func (cgm *ContainerGroupManagerRepository) CreateContainerGroup(ctx echo.Context) {
	panic("implement me")
}

func (cgm *ContainerGroupManagerRepository) UpdateContainerGroupConfig(ctx echo.Context) {
	panic("implement me")
}

func (cgm *ContainerGroupManagerRepository) DeleteContainerGroup(ctx echo.Context) {
	panic("implement me")
}
