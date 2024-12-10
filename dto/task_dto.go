package dto

import (
	"errors"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/google/uuid"
)

const (
	// Failed
	MESSAGE_FAILED_REGISTER_TASK = "failed create task"
	MESSAGE_FAILED_GET_LIST_TASK = "failed get list task"
	MESSAGE_FAILED_GET_TASK      = "failed get task"
	MESSAGE_FAILED_UPDATE_TASK   = "failed update task"
	MESSAGE_FAILED_DELETE_TASK   = "failed delete task"
	MESSAGE_FAILED_ASSIGN_USER   = "failed to assign user to task"
	MESSAGE_FAILED_REMOVE_USER   = "failed to remove user from task"

	// Success
	MESSAGE_SUCCESS_REGISTER_TASK = "success create task"
	MESSAGE_SUCCESS_GET_LIST_TASK = "success get list task"
	MESSAGE_SUCCESS_GET_TASK      = "success get task"
	MESSAGE_SUCCESS_UPDATE_TASK   = "success update task"
	MESSAGE_SUCCESS_DELETE_TASK   = "success delete task"
	MESSAGE_SUCCESS_ASSIGN_USER   = "successfully assigned user to task"
	MESSAGE_SUCCESS_REMOVE_USER   = "successfully removed user from task"
)

var (
	ErrCreateTask   = errors.New("failed to create task")
	ErrGetAllTask   = errors.New("failed to get all task")
	ErrGetTaskById  = errors.New("failed to get task by id")
	ErrUpdateTask   = errors.New("failed to update task")
	ErrTaskNotFound = errors.New("task not found")
	ErrDeleteTask   = errors.New("failed to delete task")
	ErrAssignUser   = errors.New("failed to assign user to task")
	ErrRemoveUser   = errors.New("failed to remove user from task")
)

type (
	TaskCreateRequest struct {
		Title       string     `json:"title" form:"title" binding:"required"`
		Description string     `json:"description" form:"description" binding:"required"`
		Status      string     `json:"status" form:"status" binding:"required"`
		DueDate     string     `json:"due_date" form:"due_date" binding:"required"`
		TeamsID     int        `json:"teams_id" form:"teams_id" binding:"required"`
		UserID      *uuid.UUID `json:"user_id" form:"user_id"`
	}

	TaskResponse struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		DueDate     time.Time `json:"due_date"`
		TeamsID     int       `json:"teams_id"`
		UserID      *uuid.UUID `json:"user_id,omitempty"`
		User        UserResponse `json:"user,omitempty"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	TaskPaginationResponse struct {
		Data []TaskResponse `json:"data"`
		PaginationResponse
	}

	GetAllTaskRepositoryResponse struct {
		Tasks []entity.Task
		PaginationResponse
	}

	TaskUpdateRequest struct {
		Title       string     `json:"title" form:"title" binding:"required"`
		Description string     `json:"description" form:"description" binding:"required"`
		Status      string     `json:"status" form:"status" binding:"required"`
		DueDate     string     `json:"due_date" form:"due_date" binding:"required"`
		UserID      *uuid.UUID `json:"user_id"` 
	}

	TaskUpdateResponse struct {
		ID          int        `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Status      string     `json:"status"`
		DueDate     time.Time  `json:"due_date"`
		UserID      *uuid.UUID `json:"user_id,omitempty"`
		UpdatedAt   time.Time  `json:"updated_at"`
	}

	AssignUserRequest struct {
		UserID uuid.UUID `json:"user_id" binding:"required"`
	}
)
