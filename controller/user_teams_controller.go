package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Caknoooo/go-gin-clean-starter/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserTeamsController struct {
	userTeamsService service.UserTeamsService
}

func NewUserTeamsController(userTeamsService service.UserTeamsService) *UserTeamsController {
	return &UserTeamsController{
		userTeamsService: userTeamsService,
	}
}

func (c *UserTeamsController) AssignUserToTeam(ctx *gin.Context) {
    teamId := ctx.Param("teamId")
    userId := ctx.Param("userId")

    log.Printf("Received teamId: %s", teamId)
    log.Printf("Received userId: %s", userId)

    userUuid, err := uuid.Parse(userId)
    if err != nil {
        log.Printf("Invalid userId: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }

    teamID, err := strconv.ParseUint(teamId, 10, 32)
    if err != nil {
        log.Printf("Invalid teamId: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
        return
    }

    if err := c.userTeamsService.AssignUserToTeam(userUuid, uint(teamID)); err != nil {
        log.Printf("Failed to assign user to team: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "user assigned to team successfully"})
}

func (c *UserTeamsController) RemoveUserFromTeam(ctx *gin.Context) {
    teamId := ctx.Param("teamId")
    userId := ctx.Param("userId")

    log.Printf("Received teamId: %s", teamId)
    log.Printf("Received userId: %s", userId)

    userUuid, err := uuid.Parse(userId)
    if err != nil {
        log.Printf("Invalid userId: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }

    teamID, err := strconv.ParseUint(teamId, 10, 32)
    if err != nil {
        log.Printf("Invalid teamId: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
        return
    }

    if err := c.userTeamsService.RemoveUserFromTeam(userUuid, uint(teamID)); err != nil {
        log.Printf("Failed to remove user from team: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "user removed from team successfully"})
}

func (c *UserTeamsController) GetUsersByTeamId(ctx *gin.Context) {
    teamId := ctx.Param("teamId")

    log.Printf("Received teamId: %s", teamId)

    teamID, err := strconv.ParseUint(teamId, 10, 32)
    if err != nil {
        log.Printf("Invalid teamId: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
        return
    }

    users, err := c.userTeamsService.GetUsersByTeamId(uint(teamID))
    if err != nil {
        log.Printf("Failed to get users by team id: %v", err)
        ctx.JSON(http.StatusNotFound, gin.H{"error": "no users found for this team"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"users": users})
}