package main

import (
	"fmt"
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/data/pgsql"
	"github.com/labstack/echo"
)

func main() {
	helpers.InitLogger("info")
	helpers.InitEnvConfigs()
	_, err := pgsql.PgDBConnect()
	if err != nil {
		helpers.Logger.Error(nil, "Error while connecting to pgsql data source", err)
		panic(err)
	}
	e := echo.New()
	e.Use(helpers.SetTracingDetails)
	DWSRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", helpers.EnvConfigs.LocalServerPort)))
}
