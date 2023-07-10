package repository

import (
	"fmt"
	"github.com/DWS-Backend/provider/common/exceptions"
	"github.com/DWS-Backend/provider/common/helpers"
	"github.com/DWS-Backend/provider/project-manager/class"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"time"
)

type ProjectManagerRepository struct {
	db     *gorm.DB
	logger helpers.LogConfig
}

func NewProjectManagerRepository(conn *gorm.DB, logPtr helpers.LogConfig) ProjectManagerRepoUsecases {
	return &ProjectManagerRepository{db: conn, logger: logPtr}
}

func (p *ProjectManagerRepository) executeQuery(ctx echo.Context, query string, resultType interface{}) (
	response interface{}, err error) {
	p.logger.Info(ctx, fmt.Sprintf("Executing Query: %s", query))
	hasError := p.db.Raw(query).Scan(&resultType).Error
	return resultType, hasError
}

// GetAllProjects : Gets all active projects from db
func (p *ProjectManagerRepository) GetAllProjects(ctx echo.Context) (response []class.GetAllProjectsDAO, err error) {
	var result []class.GetAllProjectsDAO
	var dbResult []class.GetAllProjectsDAO
	var interimResult interface{}

	interimResult, hasError := p.executeQuery(ctx,
		"SELECT * FROM projects WHERE \"deactivatedAt\" IS NULL ORDER BY \"createdAt\" DESC", result)
	if hasError != nil {
		return result, hasError
	}
	dbResult, ok := interimResult.([]class.GetAllProjectsDAO)
	if !ok {
		errorToThrow := exceptions.UnexpectedError
		errorToThrow.UpdateRefID(ctx)
		p.logger.Info(ctx, "There is some error while casting to desired type, exiting with error")
		return response, &errorToThrow
	}
	result = append(result, dbResult...)
	return result, hasError
}

// CreateProject : Creates a new project with specified details in db
func (p *ProjectManagerRepository) CreateProject(ctx echo.Context, projectName string, projectDescription string) (err error) {

	_, hasError := p.executeQuery(ctx,
		fmt.Sprintf("INSERT INTO projects VALUES (default, '%s', '%s', %s, %s, '%s', '%s');",
			projectName, projectDescription, "NULL", "NULL", time.Now().Format(time.DateTime),
			time.Now().Format(time.DateTime)), nil)
	return hasError
}

func (p *ProjectManagerRepository) UpdateProjectByID(ctx echo.Context, projectID int, updatedProjectDescription *string, isDeactivated bool, isStarted bool) (response class.UpdateProjectResponse, err error) {
	panic("implement me")
}

func (p *ProjectManagerRepository) DeleteProjectByID(ctx echo.Context, projectID int) (response class.DeleteProjectResponse, err error) {
	panic("implement me")
}
