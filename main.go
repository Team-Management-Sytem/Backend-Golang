package main

import (
	"log"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/command"
	"github.com/Caknoooo/go-gin-clean-starter/config"
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/middleware"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
	"github.com/Caknoooo/go-gin-clean-starter/routes"
	"github.com/Caknoooo/go-gin-clean-starter/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	err := db.AutoMigrate(&entity.User{},
		&entity.Team{},
		&entity.Task{},)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		teamRepository repository.TeamRepository = repository.NewTeamRepository(db)
		taskRepository repository.TaskRepository = repository.NewTaskRepository(db)

		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)
		teamService service.TeamService = service.NewTeamService(teamRepository)
		taskService service.TaskService = service.NewTaskService(taskRepository)

		// Controller
		userController controller.UserController = controller.NewUserController(userService)
		teamController controller.TeamController = controller.NewTeamController(teamService)
		taskController controller.TaskController = controller.NewTaskController(taskService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)
	routes.Team(server, teamController)
	routes.Task(server, taskController)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server:%v", err)
	}
}
