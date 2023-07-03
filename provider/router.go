package main

import (
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/labstack/echo"
	"net/http"
)

func DWSRoutes(e *echo.Echo) {
	e.GET("/healthCheck", checkHealth)

}

func checkHealth(c echo.Context) error {

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	helpers.Logger.Info(c, "Health Check route called")
	return c.JSON(http.StatusOK, &GenericResponse{
		Status: "success",
	})
}
