package controller

import (
	"net/http"
	"strconv"
	"time"
	"todo-api/src/model"
	"todo-api/src/usecase"

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
		ctx.JSON(http.StatusBadRequest, err.Error())
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

func (p *TaskController) DeleteTasks(ctx *gin.Context) {
	err := p.TaskUseCase.DeleteTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, model.Task{})
}

func (p *TaskController) UpdateTask(ctx *gin.Context) {
	var newTask = model.Task{}
	id := ctx.Params.ByName("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	parsedId, _ := strconv.Atoi(id)

	currentTask, err := p.TaskUseCase.GetTask(parsedId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	err = ctx.BindJSON(&newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	currentTask.Title = newTask.Title
	currentTask.Description = newTask.Description
	currentTask.RewardInSats = newTask.RewardInSats
	currentTask.IsCompleted = newTask.IsCompleted

	if newTask.IsCompleted {
		currentTask.CompletedAt.Val = time.Now()
		currentTask.CompletedAt.Valid = true
	}

	if !newTask.IsCompleted {
		currentTask.CompletedAt.Val = time.Time{}
	}

	updatedTaskId, err := p.TaskUseCase.UpdateTask(currentTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	updatedTask, _ := p.TaskUseCase.GetTask(updatedTaskId)

	ctx.JSON(http.StatusOK, updatedTask)
}
