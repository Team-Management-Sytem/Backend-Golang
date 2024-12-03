package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
)

func Team(route *gin.Engine, teamController controller.TeamController) {
	routes := route.Group("/api/teams")
	{
		routes.POST("", teamController.Register)
		routes.GET("", teamController.GetAllTeam)
		routes.GET("/:teamId", teamController.GetTeamById)
		routes.PATCH("/:teamId", teamController.Update)
		routes.DELETE("/:teamId", teamController.Delete)
	}
}

