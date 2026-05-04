package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

type User struct {
	BaseModel
	Username           string              `gorm:"unique; not null" json:"username" `
	DisplayName        string              `json:"displayName"`
	Email              string              `gorm:"unique" json:"email"`
	GlobalRoles        []Role              `gorm:"many2many:user_global_roles;" json:"globalRoles"`
	ProjectAssignments []ProjectAssignment `json:"projectAssignments"`
}

type Role struct {
	BaseModel
	Name        string       `gorm:"unique; not null" json:"name"`
	Description string       `json:"description"`
	IsProject   bool         `json:"isProject"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
}

type Permission struct {
	BaseModel
	Name      string `gorm:"unique; not null" json:"name"`
	IsProject bool   `json:"isProject"`
	OwnedBy   string `gorm:"unique; not null" json:"ownedBy"`
}

type Project struct {
	BaseModel
	Name        string   `gorm:"unique; not null" json:"name"`
	Description string   `json:"description"`
	Plugins     []Plugin `gorm:"many2many:project_plugins;" json:"plugins"`
}

type ProjectAssignment struct {
	BaseModel
	UserID    uint    `gorm:"primaryKey"`
	ProjectID uint    `gorm:"primaryKey"`
	RoleID    uint    `gorm:"primaryKey"`
	User      User    `json:"-"`
	Project   Project `json:"project"`
	Role      Role    `json:"role"`
}

type Plugin struct {
	BaseModel
	Name  string `gorm:"unique; not null" json:"name"`
	Url   string `gorm:"unique; not null" json:"url"`
	Local bool   `json:"isLocal"`
}
