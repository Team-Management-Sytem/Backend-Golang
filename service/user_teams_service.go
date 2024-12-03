package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
	"github.com/google/uuid"
)

type UserTeamsService interface {
	AssignUserToTeam(userId uuid.UUID, teamId uint) error
	RemoveUserFromTeam(userId uuid.UUID, teamId uint) error
	GetUsersByTeamId(teamId uint) ([]entity.User, error)
}

type userTeamsService struct {
	userTeamsRepo repository.UserTeamsRepository
}

func NewUserTeamsService(userTeamsRepo repository.UserTeamsRepository) UserTeamsService {
	return &userTeamsService{
		userTeamsRepo: userTeamsRepo,
	}
}

func (s *userTeamsService) AssignUserToTeam(userId uuid.UUID, teamId uint) error {
	return s.userTeamsRepo.AssignUserToTeam(context.Background(), nil, userId, teamId)
}

func (s *userTeamsService) RemoveUserFromTeam(userId uuid.UUID, teamId uint) error {
	return s.userTeamsRepo.RemoveUserFromTeam(context.Background(), nil, userId, teamId)
}

func (s *userTeamsService) GetUsersByTeamId(teamId uint) ([]entity.User, error) {
	return s.userTeamsRepo.GetUsersByTeamId(context.Background(), nil, teamId)
}