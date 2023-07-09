package services

import (
	"github.com/DWS-Backend/provider/common/exceptions"
	"github.com/DWS-Backend/provider/project-manager/class"
	"github.com/labstack/echo"
)

type ProjectManagerServiceUsecases interface {
	GetAllProjects(ctx echo.Context) (response class.GetAllProjectsResponse, err *exceptions.DWSErrorResponse)
	CreateProject(ctx echo.Context, projectName string, projectDescription string) (
		response *class.GetAllProjectsResponse, err *exceptions.DWSErrorResponse)
	UpdateProject(ctx echo.Context, projectID int, updatedProjectDescription *string, isDeactivated bool, isStarted bool) (
		response *class.GetAllProjectsResponse, err *exceptions.DWSErrorResponse)
	DeleteProject(ctx echo.Context, projectID int) (response *class.GetAllProjectsResponse, err *exceptions.DWSErrorResponse)
	StopProject(ctx echo.Context, projectID int) (response *class.GetAllProjectsResponse, err *exceptions.DWSErrorResponse)
	StartProject(ctx echo.Context, projectID int) (response *class.GetAllProjectsResponse, err *exceptions.DWSErrorResponse)
}
