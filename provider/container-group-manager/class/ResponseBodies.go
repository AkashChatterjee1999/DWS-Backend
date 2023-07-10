package class

import (
	"time"
)

type ProjectResponse struct {
	ProjectID   int        `json:"projectID"`
	Name        string     `json:"projectName"`
	Description string     `json:"projectDescription"`
	IsActive    bool       `json:"isActive"`
	StartedAt   *time.Time `json:"startedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type GetAllProjectsResponse struct {
	AllProjects []ProjectResponse `json:"allProjects"`
}

type CreateProjectResponse struct {
	Status string `json:"status"`
}

type UpdateProjectResponse struct {
	Status string `json:"status"`
}

type DeleteProjectResponse struct {
	Status string `json:"status"`
}

type StopProjectResponse struct {
	Status string `json:"status"`
}

type StartProjectResponse struct {
	Status string `json:"status"`
}
