package api

import (
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/labstack/echo"
	"net/http"
)

func GatewayManagerRoutes(e *echo.Echo) {
	e.POST("/gateway/", createGateway)                   // Create a gateway to add in a project
	e.PATCH("/gateway/:gatewayID", updateGatewayConfigs) // Update a gateway configs
	e.DELETE("/gateway/:gatewayID", deleteGateway)       // Deletes a gateway attached to a project
}

func createGateway(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func updateGatewayConfigs(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func deleteGateway(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}
