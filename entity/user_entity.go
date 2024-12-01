package entity

import (
	"github.com/Caknoooo/go-gin-clean-starter/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name       string         `json:"name"`
	TelpNumber string         `json:"telp_number"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Role       string         `json:"role"`
	ImageUrl   string         `json:"image_url"`
	IsVerified bool           `json:"is_verified"`
	CreatedAt  time.Time      `gorm:"type:datetime" json:"created_at"` // Tipe waktu disesuaikan untuk MySQL
	UpdatedAt  time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return nil
}
