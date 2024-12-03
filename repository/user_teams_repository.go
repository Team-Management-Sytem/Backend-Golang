package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTeamsRepository interface {
	AssignUserToTeam(ctx context.Context, tx *gorm.DB, userId uuid.UUID, teamId uint) error
	RemoveUserFromTeam(ctx context.Context, tx *gorm.DB, userId uuid.UUID, teamId uint) error
	GetUsersByTeamId(ctx context.Context, tx *gorm.DB, teamId uint) ([]entity.User, error)
}

type userTeamsRepository struct {
	db *gorm.DB
}

func NewUserTeamsRepository(db *gorm.DB) UserTeamsRepository {
	return &userTeamsRepository{
		db: db,
	}
}

func (r *userTeamsRepository) AssignUserToTeam(ctx context.Context, tx *gorm.DB, userId uuid.UUID, teamId uint) error {
    if tx == nil {
        tx = r.db
    }

    var existingUserTeam entity.UserTeams
    if err := tx.WithContext(ctx).Where("user_id = ? AND team_id = ?", userId, teamId).First(&existingUserTeam).Error; err == nil {
        log.Printf("User %s already assigned to team %d", userId, teamId)
        return fmt.Errorf("user already assigned to this team")
    }

    userTeam := entity.UserTeams{
        UserID:    userId,
        TeamID:    teamId,
        CreatedAt: time.Now(),
    }
    log.Printf("Assigning user %s to team %d", userId, teamId)
    return tx.WithContext(ctx).Create(&userTeam).Error
}

func (r *userTeamsRepository) RemoveUserFromTeam(ctx context.Context, tx *gorm.DB, userId uuid.UUID, teamId uint) error {
	if tx == nil {
		tx = r.db
	}
	return tx.WithContext(ctx).Where("user_id = ? AND team_id = ?", userId, teamId).Delete(&entity.UserTeams{}).Error
}

func (r *userTeamsRepository) GetUsersByTeamId(ctx context.Context, tx *gorm.DB, teamId uint) ([]entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	var users []entity.User
	if err := tx.WithContext(ctx).
		Joins("JOIN user_teams ON users.id = user_teams.user_id").
		Where("user_teams.team_id = ?", teamId).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}