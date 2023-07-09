package repository

import (
	"github.com/DWS-Backend/provider/common/exceptions"
	"github.com/DWS-Backend/provider/project-manager/class"
	"github.com/labstack/echo"
)

type ProjectManagerRepoUsecases interface {
	GetAllProjects(ctx echo.Context) (response []class.GetAllProjectsDAO, err error)
	CreateProject(ctx echo.Context, projectName string, projectDescription string) (err exceptions.DWSErrorResponse)
	UpdateProjectByID(ctx echo.Context, projectID int, updatedProjectDescription *string, isDeactivated bool, isStarted bool) (
		response class.UpdateProjectResponse, err exceptions.DWSErrorResponse)
	DeleteProjectByID(ctx echo.Context, projectID int) (response class.DeleteProjectResponse, err exceptions.DWSErrorResponse)
}
