package models

import (
	"time"
)

type ProfileModel struct {
	ID        int64      `gorm:"primaryKey;" json:"id"`
	OwnerID   int64      `gorm:"uniqueIndex" json:"owner_id"`
	FirstName string     `gorm:"type:varchar(255);default:null" json:"first_name"`
	LastName  string     `gorm:"type:varchar(255);default:null" json:"last_name"`
	Bio       string     `gorm:"type:varchar(255);default:null" json:"bio"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName sets the table name for GORM
func (ProfileModel) TableName() string {
	return "profile"
}
