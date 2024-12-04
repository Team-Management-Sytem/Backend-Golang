package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
)

func Task(route *gin.Engine, taskController controller.TaskController) {
	routes := route.Group("/api/tasks")
	{
		routes.POST("", taskController.Register)
		routes.GET("", taskController.GetAllTask)
		routes.GET("/:taskId", taskController.GetTaskById)
		routes.PATCH("/:taskId", taskController.Update)
		routes.DELETE("/:taskId", taskController.Delete)

		routes.POST("/:taskId/assign", taskController.AssignUser)
		routes.POST("/:taskId/remove", taskController.RemoveUser)
	}
}

