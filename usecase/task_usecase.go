package usecase

import (
	"todo-api/model"
	"todo-api/repository"
)

type TaskUseCase struct {
	repository repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) TaskUseCase {
	return TaskUseCase{
		repository: repo,
	}
}

func (pu *TaskUseCase) GetTasks() ([]model.Task, error) {
	return pu.repository.GetTasks()
}

func (pu *TaskUseCase) CreateTask(task model.Task) (model.Task, error) {
	taskId, err := pu.repository.CreateTask(task)
	if err != nil {
		return model.Task{}, err
	}
	task.Id = taskId
	return task, nil
}
