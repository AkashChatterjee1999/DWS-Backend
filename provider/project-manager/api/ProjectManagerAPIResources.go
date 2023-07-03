package api

import (
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/labstack/echo"
	"net/http"
)

func ProjectManagerRoutes(e *echo.Echo) {
	e.GET("/project/", getAllProjects)                // Gets all projects
	e.POST("/project/", createProject)                // Create a project
	e.PATCH("/project/:projectID", updateProject)     // Update a project
	e.DELETE("/project/:projectID", deleteProject)    // Deletes a project
	e.POST("/project/:projectID/stop", stopProject)   // Stops a project
	e.POST("/project/:projectID/start", startProject) // Starts a project
}

func getAllProjects(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func createProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func updateProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func deleteProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func stopProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}

func startProject(c echo.Context) error {
	helpers.Logger.Info(c, c.Request().Method+":/"+c.Request().RequestURI)
	return c.JSON(http.StatusOK, "Success")
}
