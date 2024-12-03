package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserTeams struct {
	UserID    uuid.UUID `gorm:"type:char(36);primaryKey" json:"user_id"`
	TeamID    uint      `gorm:"primaryKey" json:"team_id"`
	CreatedAt time.Time `json:"created_at"`
}