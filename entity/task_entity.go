package entity

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Status      string         `gorm:"type:varchar(50);not null" json:"status"`
	DueDate     time.Time      `gorm:"type:datetime" json:"due_date"`
	TeamsID     int            `gorm:"not null" json:"teams_id"`
	UserID      int            `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	
	Team Team `gorm:"foreignKey:TeamsID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"team"`
}
