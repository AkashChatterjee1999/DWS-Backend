package api

import (
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/labstack/echo"
	"net/http"
)

func GatewayManagerRoutes(e *echo.Echo) {
	e.POST("/containerGroup/", createContainerGroup)                 // Create a gateway to add in a project
	e.PATCH("/containerGroup/:groupID", updateContainerGroupConfig)  // Update a gateway configs
	e.DELETE("/containerGroup/:groupID", deleteContainerGroupConfig) // Deletes a gateway attached to a project
}

func createContainerGroup(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func updateContainerGroupConfig(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func deleteContainerGroupConfig(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}
