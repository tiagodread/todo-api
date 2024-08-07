package controller

import (
	"net/http"
	"strconv"
	"todo-api/model"
	"todo-api/usecase"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUseCase usecase.TaskUseCase
}

func NewTaskController(usecase usecase.TaskUseCase) TaskController {
	return TaskController{
		TaskUseCase: usecase,
	}
}

func (p *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := p.TaskUseCase.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (p *TaskController) GetTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	task, err := p.TaskUseCase.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (p *TaskController) CreateTask(ctx *gin.Context) {
	var task model.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTask, err := p.TaskUseCase.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTask)
}

func (p *TaskController) DeleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	task, err := p.TaskUseCase.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	err = p.TaskUseCase.DeleteTask(task.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, model.Task{})
}
