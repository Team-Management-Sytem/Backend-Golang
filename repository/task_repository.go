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
	TaskRepository interface {
		RegisterTask(ctx context.Context, tx *gorm.DB, task entity.Task) (entity.Task, error)
		GetAllTaskWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllTaskRepositoryResponse, error)
		GetTaskById(ctx context.Context, tx *gorm.DB, taskId string) (entity.Task, error)
		UpdateTask(ctx context.Context, tx *gorm.DB, task entity.Task) (entity.Task, error)
		DeleteTask(ctx context.Context, tx *gorm.DB, taskId string) error
	}

	taskRepository struct {
		db *gorm.DB
	}
)

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) RegisterTask(ctx context.Context, tx *gorm.DB, task entity.Task) (entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	log.Printf("Registering task: %+v", task) 
	if err := tx.WithContext(ctx).Create(&task).Error; err != nil {
		log.Printf("Failed to register task: %v", err) 
		return entity.Task{}, err
	}

	log.Printf("Task registered successfully: %+v", task)
	return task, nil
}

func (r *taskRepository) GetAllTaskWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllTaskRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var tasks []entity.Task
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 20
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Task{}).Count(&count).Error; err != nil {
		return dto.GetAllTaskRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&tasks).Error; err != nil {
		return dto.GetAllTaskRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllTaskRepositoryResponse{
		Tasks: tasks,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *taskRepository) GetTaskById(ctx context.Context, tx *gorm.DB, taskId string) (entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	var task entity.Task
	if err := tx.WithContext(ctx).Where("id = ?", taskId).Take(&task).Error; err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, tx *gorm.DB, task entity.Task) (entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&task).Error; err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, tx *gorm.DB, taskId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Task{}, "id = ?", taskId).Error; err != nil {
		return err
	}

	return nil
}