package dto

import (
	"errors"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
)

const (
	// Failed
	MESSAGE_FAILED_REGISTER_TEAM           = "failed create team"
	MESSAGE_FAILED_GET_LIST_TEAM           = "failed get list team"
	MESSAGE_FAILED_GET_TEAM                = "failed get team"
	MESSAGE_FAILED_UPDATE_TEAM             = "failed update team"
	MESSAGE_FAILED_DELETE_TEAM             = "failed delete team"

	// Success
	MESSAGE_SUCCESS_REGISTER_TEAM           = "success create team"
	MESSAGE_SUCCESS_GET_LIST_TEAM           = "success get list team"
	MESSAGE_SUCCESS_GET_TEAM                = "success get team"
	MESSAGE_SUCCESS_UPDATE_TEAM             = "success update team"
	MESSAGE_SUCCESS_DELETE_TEAM             = "success delete team"
)

var (
	ErrCreateTeam             = errors.New("failed to create team")
	ErrGetAllTeam             = errors.New("failed to get all team")
	ErrGetTeamById            = errors.New("failed to get team by id")
	ErrUpdateTeam             = errors.New("failed to update team")
	ErrTeamNotFound           = errors.New("team not found")
	ErrDeleteTeam             = errors.New("failed to delete team")
)

type (
	TeamCreateRequest struct {
		Name        string `json:"name" form:"name" binding:"required"`
    	Description string `json:"description" form:"description" binding:"required"`
	}

	TeamResponse struct {
		ID         	string `json:"id"`
		Name       	string `json:"name"`
		Description string `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	TeamPaginationResponse struct {
		Data []TeamResponse `json:"data"`
		PaginationResponse
	}

	GetAllTeamRepositoryResponse struct {
		Teams []entity.Team
		PaginationResponse
	}

	TeamUpdateRequest struct {
		Name       	string `json:"name" form:"name" binding:"required"`
		Description string `json:"description" form:"description" binding:"required"`
	}

	TeamUpdateResponse struct {
		ID         	string `json:"id"`
		Name       	string `json:"name"`
		Description string `json:"description"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
