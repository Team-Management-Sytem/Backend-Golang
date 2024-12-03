package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Caknoooo/go-gin-clean-starter/controller"
)

func UserTeams(route *gin.Engine, userTeamsController *controller.UserTeamsController) {
    routes := route.Group("/api/teams")
    {
        routes.POST("/:teamId/users/:userId", userTeamsController.AssignUserToTeam)
        routes.DELETE("/:teamId/users/:userId", userTeamsController.RemoveUserFromTeam)
        routes.GET("/:teamId/users", userTeamsController.GetUsersByTeamId)
    }
}
