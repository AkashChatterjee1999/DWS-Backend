package api

import (
	"fmt"
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/project-manager/services"
	"github.com/labstack/echo"
	"net/http"
)

type ProjectManagerHandler struct {
	ProjectManagerUsecase services.ProjectManagerServiceUsecases
	logger                helpers.LogConfig
}

func NewProjectManagerRoutes(services services.ProjectManagerServiceUsecases, logger helpers.LogConfig, e *echo.Echo) {
	p := ProjectManagerHandler{services, logger}
	e.GET("/api/v1/project/", p.getAllProjects)                // Gets all projects
	e.POST("/api/v1/project/", p.createProject)                // Create a project
	e.PATCH("/api/v1/project/:projectID", p.updateProject)     // Update a project
	e.DELETE("/api/v1/project/:projectID", p.deleteProject)    // Deletes a project
	e.POST("/api/v1/project/:projectID/stop", p.stopProject)   // Stops a project
	e.POST("/api/v1/project/:projectID/start", p.startProject) // Starts a project
}

func (p *ProjectManagerHandler) getAllProjects(c echo.Context) error {
	p.logger.Info(c, "Executing getAllProjects route")
	p.logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response, err := p.ProjectManagerUsecase.GetAllProjects(c)

	if err != nil {
		p.logger.Info(c, fmt.Sprintf("Route returned Code: 400 Response: %s", err))
		return c.JSON(http.StatusBadRequest, err)
	}
	p.logger.Info(c, fmt.Sprintf("Route returned Code: 200 Response: %s \n", response))
	return c.JSON(http.StatusOK, response)
}

func (p *ProjectManagerHandler) createProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func (p *ProjectManagerHandler) updateProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func (p *ProjectManagerHandler) deleteProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func (p *ProjectManagerHandler) stopProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func (p *ProjectManagerHandler) startProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}
