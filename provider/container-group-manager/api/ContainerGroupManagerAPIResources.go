package api

import (
	"fmt"
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/container-group-manager/services"
	"github.com/labstack/echo"
	"net/http"
)

type ContainerGroupManagerHandler struct {
	ContainerGroupManagerUsecase services.ContainerGroupManagerServiceUsecases
	logger                       helpers.LogConfig
}

func NewContainerGroupManagerRoutes(services services.ContainerGroupManagerServiceUsecases, logger helpers.LogConfig, e *echo.Echo) {
	p := ContainerGroupManagerHandler{services, logger}
	e.POST("/api/v1/containerGroups/", p.createContainerGroup)                // Create a container group to add in a project
	e.PATCH("/api/v1/containerGroup/:groupID", p.updateContainerGroupConfig)  // Update a container group configs
	e.DELETE("/api/v1/containerGroup/:groupID", p.deleteContainerGroupConfig) // Deletes a container group attached to a project
}

func (p *ContainerGroupManagerHandler) createContainerGroup(c echo.Context) error {
	p.logger.Info(c, "Executing createContainerGroup route")
	p.logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)

	response := "Success"

	p.logger.Info(c, fmt.Sprintf("Route returned Code: 200 Response: %s \n", response))
	return c.JSON(http.StatusOK, response)
}

func (p *ContainerGroupManagerHandler) updateContainerGroupConfig(c echo.Context) error {
	p.logger.Info(c, "Executing updateContainerGroupConfig route")
	p.logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)

	response := "Success"

	p.logger.Info(c, fmt.Sprintf("Route returned Code: 200 Response: %s \n", response))
	return c.JSON(http.StatusOK, response)
}

func (p *ContainerGroupManagerHandler) deleteContainerGroupConfig(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}
