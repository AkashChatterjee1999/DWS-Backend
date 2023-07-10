package services

import (
	"github.com/DWS-Backend/provider/common/exceptions"
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/project-manager/class"
	"github.com/DWS-Backend/provider/project-manager/repository"
	"github.com/labstack/echo"
)

type ProjectManagerServices struct {
	repo   repository.ProjectManagerRepoUsecases
	logger helpers.LogConfig
}

func NewProjectManagerServices(repoInstance repository.ProjectManagerRepoUsecases, logger helpers.LogConfig) ProjectManagerServiceUsecases {
	return &ProjectManagerServices{repo: repoInstance, logger: logger}
}

func (p *ProjectManagerServices) GetAllProjects(ctx echo.Context) (response class.GetAllProjectsResponse, err *exceptions.DWSErrorResponse) {
	dbResult, hasError := p.repo.GetAllProjects(ctx)
	if hasError != nil {
		p.logger.Error(ctx, "There is some error occurred in repoLayer", hasError)
		errResponse := exceptions.UnexpectedError
		errResponse.UpdateRefID(ctx)
		return response, &errResponse
	}
	response.AllProjects = make([]class.ProjectResponse, 0)
	for i := range dbResult {
		isActive := true
		if dbResult[i].DeactivatedAt != nil {
			isActive = false
		}
		response.AllProjects = append(response.AllProjects, class.ProjectResponse{
			ProjectID:   dbResult[i].Id,
			Name:        dbResult[i].Name,
			Description: dbResult[i].Description,
			IsActive:    isActive,
			StartedAt:   dbResult[i].StartedAt,
			CreatedAt:   dbResult[i].CreatedAt,
			UpdatedAt:   dbResult[i].UpdatedAt,
		})
	}
	return
}

func (p *ProjectManagerServices) CreateProject(ctx echo.Context, projectName string, projectDescription string) (response class.CreateProjectResponse, err *exceptions.DWSErrorResponse) {
	hasError := p.repo.CreateProject(ctx, projectName, projectDescription)
	if hasError != nil {
		p.logger.Error(ctx, "There is some error occurred in repoLayer", hasError)
		errResponse := exceptions.UnexpectedError
		errResponse.UpdateRefID(ctx)
		return response, &errResponse
	}
	response.Status = "success"
	return response, nil
}

func (p *ProjectManagerServices) UpdateProject(ctx echo.Context, projectID int, updatedProjectDescription *string, isDeactivated bool, isStarted bool) (response class.UpdateProjectResponse, err *exceptions.DWSErrorResponse) {
	panic("implement me")
}

func (p *ProjectManagerServices) DeleteProject(ctx echo.Context, projectID int) (response class.DeleteProjectResponse, err *exceptions.DWSErrorResponse) {
	panic("implement me")
}

func (p *ProjectManagerServices) StopProject(ctx echo.Context, projectID int) (response class.StopProjectResponse, err *exceptions.DWSErrorResponse) {
	panic("implement me")
}

func (p *ProjectManagerServices) StartProject(ctx echo.Context, projectID int) (response class.StartProjectResponse, err *exceptions.DWSErrorResponse) {
	panic("implement me")
}
