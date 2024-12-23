package service

import (
	"context"
	"errors"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
	"github.com/google/uuid"
)

type (
	TaskService interface {
		Register(ctx context.Context, req dto.TaskCreateRequest) (dto.TaskResponse, error)
		GetAllTaskWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TaskPaginationResponse, error)
		GetTaskById(ctx context.Context, taskId string) (dto.TaskResponse, error)
		GetTasksByTeamID(ctx context.Context, teamsID int) ([]dto.TaskResponse, error)
		Update(ctx context.Context, req dto.TaskUpdateRequest, taskId string) (dto.TaskUpdateResponse, error)
		Delete(ctx context.Context, taskId string) error
		AssignUserToTask(ctx context.Context, taskId string, userID *uuid.UUID) error
		RemoveUserFromTask(ctx context.Context, taskId string) error
		GetAssignedUser(ctx context.Context, taskId string) (dto.UserResponse, error)
		GetTasksByUserID(ctx context.Context, userID string) ([]dto.TaskResponse, error)
	}

	taskService struct {
		taskRepo repository.TaskRepository
		userRepo repository.UserRepository
	}
)

func NewTaskService(taskRepo repository.TaskRepository, userRepo repository.UserRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (s *taskService) Register(ctx context.Context, req dto.TaskCreateRequest) (dto.TaskResponse, error) {
	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		return dto.TaskResponse{}, err
	}

	task := entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     dueDate,
		TeamsID:     req.TeamsID,
	}

	if req.UserID != nil {
		task.UserID = req.UserID
	}

	taskReg, err := s.taskRepo.RegisterTask(ctx, nil, task)
	if err != nil {
		return dto.TaskResponse{}, dto.ErrCreateTask
	}

	return dto.TaskResponse{
		ID:          taskReg.ID,
		Title:       taskReg.Title,
		Description: taskReg.Description,
		Status:      taskReg.Status,
		DueDate:     taskReg.DueDate,
		UserID:      taskReg.UserID,
	}, nil
}

func (s *taskService) GetAllTaskWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TaskPaginationResponse, error) {
	dataWithPaginate, err := s.taskRepo.GetAllTaskWithPagination(ctx, nil, req)
	if err != nil {
		return dto.TaskPaginationResponse{}, err
	}

	var tasks []dto.TaskResponse
	for _, task := range dataWithPaginate.Tasks {
		tasks = append(tasks, dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     task.DueDate,
			UserID:      task.UserID,
		})
	}

	return dto.TaskPaginationResponse{
		Data: tasks,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *taskService) GetTaskById(ctx context.Context, taskId string) (dto.TaskResponse, error) {
	task, err := s.taskRepo.GetTaskById(ctx, nil, taskId)
	if err != nil {
		return dto.TaskResponse{}, dto.ErrGetTaskById
	}

	var userResponse dto.UserResponse
	if task.UserID != nil {
		user, err := s.userRepo.GetUserById(ctx, nil, task.UserID.String())
		if err != nil {
			return dto.TaskResponse{}, err
		}

		userResponse = dto.UserResponse{
			ID:         user.ID.String(),
			Name:       user.Name,
			Email:      user.Email,
			TelpNumber: user.TelpNumber,
			Role:       user.Role,
			ImageUrl:   user.ImageUrl,
			IsVerified: user.IsVerified,
		}
	}

	return dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
		TeamsID:     task.TeamsID,
		UserID:      task.UserID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		User:        userResponse,
	}, nil
}


func (s *taskService) GetTasksByTeamID(ctx context.Context, teamsID int) ([]dto.TaskResponse, error) {

    tasks, err := s.taskRepo.GetTasksByTeamID(ctx, nil, teamsID)
    if err != nil {
        return nil, err
    }

    var taskResponses []dto.TaskResponse
    for _, task := range tasks {
    
        var userResponse dto.UserResponse
        if task.UserID != nil {
            user, err := s.userRepo.GetUserById(ctx, nil, task.UserID.String())
            if err != nil {
                return nil, err
            }

            userResponse = dto.UserResponse{
                ID:         user.ID.String(),
                Name:       user.Name,
                Email:      user.Email,
                TelpNumber: user.TelpNumber,
                Role:       user.Role,
                ImageUrl:   user.ImageUrl,
                IsVerified: user.IsVerified,
            }
        }

    
        taskResponses = append(taskResponses, dto.TaskResponse{
            ID:          task.ID,
            Title:       task.Title,
            Description: task.Description,
            Status:      task.Status,
            DueDate:     task.DueDate,
            TeamsID:     task.TeamsID,
            UserID:      task.UserID,
            User:        userResponse,
        })
    }

    return taskResponses, nil
}


func (s *taskService) Update(ctx context.Context, req dto.TaskUpdateRequest, taskId string) (dto.TaskUpdateResponse, error) {
	task, err := s.taskRepo.GetTaskById(ctx, nil, taskId)
	if err != nil {
		return dto.TaskUpdateResponse{}, dto.ErrTaskNotFound
	}

	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		return dto.TaskUpdateResponse{}, err
	}

	data := entity.Task{
		ID:          task.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     dueDate,
	}

	if req.UserID != nil {
		data.UserID = req.UserID
	}

	taskUpdate, err := s.taskRepo.UpdateTask(ctx, nil, data)
	if err != nil {
		return dto.TaskUpdateResponse{}, dto.ErrUpdateTask
	}

	return dto.TaskUpdateResponse{
		ID:          taskUpdate.ID,
		Title:       taskUpdate.Title,
		Description: taskUpdate.Description,
		Status:      taskUpdate.Status,
		DueDate:     taskUpdate.DueDate,
		UserID:      taskUpdate.UserID,
	}, nil
}

func (s *taskService) Delete(ctx context.Context, taskId string) error {
	_, err := s.taskRepo.GetTaskById(ctx, nil, taskId)
	if err != nil {
		return dto.ErrTaskNotFound
	}

	err = s.taskRepo.DeleteTask(ctx, nil, taskId)
	if err != nil {
		return dto.ErrDeleteTask
	}

	return nil
}

func (s *taskService) AssignUserToTask(ctx context.Context, taskId string, userID *uuid.UUID) error {
	task, err := s.taskRepo.GetTaskById(ctx, nil, taskId)
	if err != nil {
		return dto.ErrTaskNotFound
	}

	if task.UserID != nil {
		return errors.New("task already assigned to another user")
	}

	err = s.taskRepo.AssignUserToTask(ctx, nil, taskId, userID)
	if err != nil {
		return dto.ErrAssignUser
	}

	return nil
}

func (s *taskService) RemoveUserFromTask(ctx context.Context, taskId string) error {
	err := s.taskRepo.RemoveUserFromTask(ctx, nil, taskId)
	if err != nil {
		return dto.ErrUpdateTask
	}
	return nil
}

func (s *taskService) GetAssignedUser(ctx context.Context, taskId string) (dto.UserResponse, error) {
	task, err := s.taskRepo.GetTaskById(ctx, nil, taskId)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if task.UserID != nil {
		user, err := s.userRepo.GetUserById(ctx, nil, task.UserID.String())
		if err != nil {
			return dto.UserResponse{}, err
		}

		userResponse := dto.UserResponse{
			ID:         user.ID.String(),
			Name:       user.Name,
			Email:      user.Email,
			TelpNumber: user.TelpNumber,
			Role:       user.Role,
			ImageUrl:   user.ImageUrl,
			IsVerified: user.IsVerified,
		}

		return userResponse, nil
	}

	return dto.UserResponse{}, errors.New("no user assigned to this task")
}

func (s *taskService) GetTasksByUserID(ctx context.Context, userID string) ([]dto.TaskResponse, error) {
	tasks, err := s.taskRepo.GetTasksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var taskResponses []dto.TaskResponse
	for _, task := range tasks {
		var userResponse dto.UserResponse
		if task.UserID != nil {
			user, err := s.userRepo.GetUserById(ctx, nil, task.UserID.String())
			if err != nil {
				return nil, err
			}

			userResponse = dto.UserResponse{
				ID:         user.ID.String(),
				Name:       user.Name,
				Email:      user.Email,
				TelpNumber: user.TelpNumber,
				Role:       user.Role,
				ImageUrl:   user.ImageUrl,
				IsVerified: user.IsVerified,
			}
		}

		taskResponses = append(taskResponses, dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     task.DueDate,
			TeamsID:     task.TeamsID,
			UserID:      task.UserID,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			User:        userResponse,
		})
	}

	return taskResponses, nil
}
