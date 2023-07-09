package class

import "time"

type GetAllProjectsDAO struct {
	Id            int        `json:"id" gorm:"column:id"`
	Name          string     `json:"name" gorm:"column:name"`
	Description   string     `json:"description" gorm:"column:description"`
	DeactivatedAt *time.Time `json:"deactivatedAt" gorm:"column:deactivatedAt"`
	StartedAt     *time.Time `json:"startedAt" gorm:"column:name"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"column:updatedAt"`
}
