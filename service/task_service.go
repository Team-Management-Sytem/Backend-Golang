package service

import (
	"context"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
)

type (
	TaskService interface {
		Register(ctx context.Context, req dto.TaskCreateRequest) (dto.TaskResponse, error)
		GetAllTaskWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TaskPaginationResponse, error)
		GetTaskById(ctx context.Context, taskId string) (dto.TaskResponse, error)
		Update(ctx context.Context, req dto.TaskUpdateRequest, taskId string) (dto.TaskUpdateResponse, error)
		Delete(ctx context.Context, taskId string) error
	}

	taskService struct {
		taskRepo repository.TaskRepository
	}
)

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
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
		UserID:      req.UserID,
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

	return dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	}, nil
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
