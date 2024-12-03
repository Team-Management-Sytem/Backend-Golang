package repository

import (
	"context"
	"log"
	"math"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

type (
	TeamRepository interface {
		RegisterTeam(ctx context.Context, tx *gorm.DB, team entity.Team) (entity.Team, error)
		GetAllTeamWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllTeamRepositoryResponse, error)
		GetTeamById(ctx context.Context, tx *gorm.DB, teamId string) (entity.Team, error)
		UpdateTeam(ctx context.Context, tx *gorm.DB, team entity.Team) (entity.Team, error)
		DeleteTeam(ctx context.Context, tx *gorm.DB, teamId string) error
	}

	teamRepository struct {
		db *gorm.DB
	}
)

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{
		db: db,
	}
}

func (r *teamRepository) RegisterTeam(ctx context.Context, tx *gorm.DB, team entity.Team) (entity.Team, error) {
	if tx == nil {
		tx = r.db
	}

	log.Printf("Registering team: %+v", team) 
	if err := tx.WithContext(ctx).Create(&team).Error; err != nil {
		log.Printf("Failed to register team: %v", err) 
		return entity.Team{}, err
	}

	log.Printf("Team registered successfully: %+v", team)
	return team, nil
}

func (r *teamRepository) GetAllTeamWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllTeamRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var teams []entity.Team
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 20
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Team{}).Count(&count).Error; err != nil {
		return dto.GetAllTeamRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&teams).Error; err != nil {
		return dto.GetAllTeamRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllTeamRepositoryResponse{
		Teams: teams,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *teamRepository) GetTeamById(ctx context.Context, tx *gorm.DB, teamId string) (entity.Team, error) {
	if tx == nil {
		tx = r.db
	}

	var team entity.Team
	if err := tx.WithContext(ctx).Where("id = ?", teamId).Take(&team).Error; err != nil {
		return entity.Team{}, err
	}

	return team, nil
}

func (r *teamRepository) UpdateTeam(ctx context.Context, tx *gorm.DB, team entity.Team) (entity.Team, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&team).Error; err != nil {
		return entity.Team{}, err
	}

	return team, nil
}

func (r *teamRepository) DeleteTeam(ctx context.Context, tx *gorm.DB, teamId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Team{}, "id = ?", teamId).Error; err != nil {
		return err
	}

	return nil
}
