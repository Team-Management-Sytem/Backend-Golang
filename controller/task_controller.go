package controller

import (
	"net/http"
	"strconv"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/service"
	"github.com/Caknoooo/go-gin-clean-starter/utils"
	"github.com/gin-gonic/gin"
)

type (
	TaskController interface {
		Register(ctx *gin.Context)
		Task(ctx *gin.Context)
		GetAllTask(ctx *gin.Context)
		GetTaskById(ctx *gin.Context)
		GetTasksByTeamID(ctx *gin.Context) 
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		AssignUser(ctx *gin.Context)
		RemoveUser(ctx *gin.Context)
		GetAssignedUser(ctx *gin.Context)
	}

	taskController struct {
		taskService service.TaskService
	}
)

func NewTaskController(ts service.TaskService) TaskController {
	return &taskController{
		taskService: ts,
	}
}

func (c *taskController) Register(ctx *gin.Context) {
	var task dto.TaskCreateRequest
	if err := ctx.ShouldBind(&task); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.taskService.Register(ctx.Request.Context(), task)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_TASK, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_TASK, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *taskController) GetAllTask(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.taskService.GetAllTaskWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TASK, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_TASK,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *taskController) Task(ctx *gin.Context) {
	taskId := ctx.MustGet("task_id").(string)

	result, err := c.taskService.GetTaskById(ctx.Request.Context(), taskId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TASK, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TASK, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *taskController) GetTaskById(ctx *gin.Context) {
    taskId := ctx.Param("taskId")

    result, err := c.taskService.GetTaskById(ctx.Request.Context(), taskId)
    if err != nil {
        res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TASK, err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TASK, result)
    ctx.JSON(http.StatusOK, res)
}

func (c *taskController) GetTasksByTeamID(ctx *gin.Context) {
    teamID := ctx.Param("teamId")
    teamIDInt, err := strconv.Atoi(teamID) 
    if err != nil {
        res := utils.BuildResponseFailed("Invalid team ID", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    tasks, err := c.taskService.GetTasksByTeamID(ctx.Request.Context(), teamIDInt)
    if err != nil {
        res := utils.BuildResponseFailed("Failed to get tasks", err.Error(), nil)
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }

    res := utils.BuildResponseSuccess("Successfully fetched tasks", tasks)
    ctx.JSON(http.StatusOK, res)
}

func (c *taskController) Update(ctx *gin.Context) {
	var req dto.TaskUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	taskId := ctx.Param("taskId")

	result, err := c.taskService.Update(ctx.Request.Context(), req, taskId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TASK, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TASK, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *taskController) Delete(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	if err := c.taskService.Delete(ctx.Request.Context(), taskId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TASK, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_TASK, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *taskController) AssignUser(ctx *gin.Context) {
    taskId := ctx.Param("taskId")
    var req dto.AssignUserRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_ASSIGN_USER, err.Error(), nil)
        ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
        return
    }

    err := c.taskService.AssignUserToTask(ctx.Request.Context(), taskId, &req.UserID)
    if err != nil {
        res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_ASSIGN_USER, err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_ASSIGN_USER, nil)
    ctx.JSON(http.StatusOK, res)
}


func (c *taskController) RemoveUser(ctx *gin.Context) {
    taskId := ctx.Param("taskId")

    err := c.taskService.RemoveUserFromTask(ctx.Request.Context(), taskId)
    if err != nil {
        res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REMOVE_USER, err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REMOVE_USER, nil)
    ctx.JSON(http.StatusOK, res)
}

func (c *taskController) GetAssignedUser(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	result, err := c.taskService.GetAssignedUser(ctx.Request.Context(), taskId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get assigned user", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully fetched task and assigned user", result)
	ctx.JSON(http.StatusOK, res)
}