package main

import (
	"todo-api/controller"
	"todo-api/db"
	"todo-api/repository"
	"todo-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	// Repository
	TaskRepository := repository.NewTaskRepository(dbConnection)

	// Use cases
	TaskUseCase := usecase.NewTaskUseCase(TaskRepository)

	// Controllers
	TaskController := controller.NewTaskController(TaskUseCase)

	server.GET("/tasks", TaskController.GetTasks)
	server.GET("/task/:id", TaskController.GetTask)
	server.POST("/task", TaskController.CreateTask)
	server.DELETE("/task/:id", TaskController.DeleteTask)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})

	server.Run(":8000")

}
