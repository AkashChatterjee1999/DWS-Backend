package main

import (
	"fmt"
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/data/pgsql"
	projectMgrRoutes "github.com/DWS-Backend/provider/project-manager/api"
	projectMgrRepo "github.com/DWS-Backend/provider/project-manager/repository"
	projectMgrServices "github.com/DWS-Backend/provider/project-manager/services"
	"github.com/labstack/echo"
)

func main() {
	helpers.InitLogger("info")
	helpers.InitEnvConfigs()
	db, err := pgsql.PgDBConnect()
	if err != nil {
		helpers.Logger.Error(nil, "Error while connecting to pgsql data source", err)
	}
	e := echo.New()
	e.Use(helpers.SetTracingDetails)
	DWSRoutes(e)

	// Initializing modules
	projectMgrRepository := projectMgrRepo.NewProjectManagerRepository(db, helpers.Logger)
	projectMgrService := projectMgrServices.NewProjectManagerServices(projectMgrRepository, helpers.Logger)
	projectMgrRoutes.NewProjectManagerRoutes(projectMgrService, helpers.Logger, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", helpers.EnvConfigs.LocalServerPort)))
}
