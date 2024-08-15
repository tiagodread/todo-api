package main

import (
	"todo-api/src/controller"
	"todo-api/src/db"
	"todo-api/src/repository"
	"todo-api/src/usecase"

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
	server.PUT("/task/:id", TaskController.UpdateTask)
	server.DELETE("/task/:id", TaskController.DeleteTask)
	server.DELETE("/tasks", TaskController.DeleteTasks)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})

	server.Run(":8080")

}
