package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/service"
	"github.com/Caknoooo/go-gin-clean-starter/utils"
	"github.com/gin-gonic/gin"
)

type (
	TeamController interface {
		Register(ctx *gin.Context)
		Team(ctx *gin.Context)
		GetAllTeam(ctx *gin.Context)
		GetTeamById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	teamController struct {
		teamService service.TeamService
	}
)

func NewTeamController(ts service.TeamService) TeamController {
	return &teamController{
		teamService: ts,
	}
}

func (c *teamController) Register(ctx *gin.Context) {
	var team dto.TeamCreateRequest
	if err := ctx.ShouldBind(&team); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.teamService.Register(ctx.Request.Context(), team)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_TEAM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_TEAM, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *teamController) GetAllTeam(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.teamService.GetAllTeamWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TEAM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_TEAM,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *teamController) Team(ctx *gin.Context) {
	teamId := ctx.MustGet("team_id").(string)

	result, err := c.teamService.GetTeamById(ctx.Request.Context(), teamId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TEAM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TEAM, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *teamController) GetTeamById(ctx *gin.Context) {
	teamId := ctx.Param("teamId")

	result, err := c.teamService.GetTeamById(ctx.Request.Context(), teamId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TEAM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TEAM, result)
	ctx.JSON(http.StatusOK, res)
}


func (c *teamController) Update(ctx *gin.Context) {
	var req dto.TeamUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	teamId := ctx.MustGet("team_id").(string)
	result, err := c.teamService.Update(ctx.Request.Context(), req, teamId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TEAM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TEAM, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *teamController) Delete(ctx *gin.Context) {
	teamId := ctx.MustGet("team_id").(string)

	if err := c.teamService.Delete(ctx.Request.Context(), teamId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TEAM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_TEAM, nil)
	ctx.JSON(http.StatusOK, res)
}
