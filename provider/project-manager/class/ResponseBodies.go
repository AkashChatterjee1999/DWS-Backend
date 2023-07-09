package class

import (
	"time"
)

type ProjectResponse struct {
	ProjectID   int
	Name        string
	Description string
	IsActive    bool
	StartedAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GetAllProjectsResponse struct {
	AllProjects []ProjectResponse `json:"allProjects"`
}

type CreateProjectResponse struct{ status string }

type UpdateProjectResponse struct{ status string }

type DeleteProjectResponse struct{ status string }

type StopProjectResponse struct{ status string }

type StartProjectResponse struct{ status string }
