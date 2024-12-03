package service

import (
	"context"
	"strconv"
	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
)

type (
	TeamService interface {
		Register(ctx context.Context, req dto.TeamCreateRequest) (dto.TeamResponse, error)
		GetAllTeamWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TeamPaginationResponse, error)
		GetTeamById(ctx context.Context, teamId string) (dto.TeamResponse, error)
		Update(ctx context.Context, req dto.TeamUpdateRequest, teamId string) (dto.TeamUpdateResponse, error)
		Delete(ctx context.Context, teamId string) error
	}

	teamService struct {
		teamRepo   repository.TeamRepository
	}
)

func NewTeamService(teamRepo repository.TeamRepository) TeamService {
	return &teamService{
		teamRepo:   teamRepo,
	}
}

func (s *teamService) Register(ctx context.Context, req dto.TeamCreateRequest) (dto.TeamResponse, error) {
	team := entity.Team{
		Name:       	req.Name,
		Description: 	req.Description,
	}

	teamReg, err := s.teamRepo.RegisterTeam(ctx, nil, team)
	if err != nil {
		return dto.TeamResponse{}, dto.ErrCreateTeam
	}

	return dto.TeamResponse{
		ID:         	strconv.Itoa(teamReg.ID),
		Name:       	teamReg.Name,
		Description: 	req.Description,
	}, nil
}

func (s *teamService) GetAllTeamWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TeamPaginationResponse, error) {
	dataWithPaginate, err := s.teamRepo.GetAllTeamWithPagination(ctx, nil, req)
	if err != nil {
		return dto.TeamPaginationResponse{}, err
	}

	var datas []dto.TeamResponse
	for _, team := range dataWithPaginate.Teams {
		data := dto.TeamResponse{
			ID:         	strconv.Itoa(team.ID),
			Name:       	team.Name,
			Description: 	team.Description,
		}

		datas = append(datas, data)
	}

	return dto.TeamPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *teamService) GetTeamById(ctx context.Context, teamId string) (dto.TeamResponse, error) {
	team, err := s.teamRepo.GetTeamById(ctx, nil, teamId)
	if err != nil {
		return dto.TeamResponse{}, dto.ErrGetTeamById
	}

	return dto.TeamResponse{
		ID:         	strconv.Itoa(team.ID),
		Name:       	team.Name,
		Description: 	team.Description,
	}, nil
}

func (s *teamService) Update(ctx context.Context, req dto.TeamUpdateRequest, teamId string) (dto.TeamUpdateResponse, error) {
	team, err := s.teamRepo.GetTeamById(ctx, nil, teamId)
	if err != nil {
		return dto.TeamUpdateResponse{}, dto.ErrTeamNotFound
	}

	data := entity.Team{
		ID:         	team.ID,
		Name:       	req.Name,
		Description: 	req.Description,
	}

	teamUpdate, err := s.teamRepo.UpdateTeam(ctx, nil, data)
	if err != nil {
		return dto.TeamUpdateResponse{}, dto.ErrUpdateTeam
	}

	return dto.TeamUpdateResponse{
		ID:         	strconv.Itoa(teamUpdate.ID),
		Name:       	teamUpdate.Name,
		Description: 	teamUpdate.Description,
	}, nil
}

func (s *teamService) Delete(ctx context.Context, teamId string) error {
	team, err := s.teamRepo.GetTeamById(ctx, nil, teamId)
	if err != nil {
		return dto.ErrTeamNotFound
	}

	err = s.teamRepo.DeleteTeam(ctx, nil, strconv.Itoa(team.ID))
	if err != nil {
		return dto.ErrDeleteTeam
	}

	return nil
}