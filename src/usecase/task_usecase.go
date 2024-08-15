package usecase

import (
	"todo-api/src/model"
	"todo-api/src/repository"
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

func (pu *TaskUseCase) GetTask(id int) (model.Task, error) {
	task, err := pu.repository.GetTask(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (pu *TaskUseCase) CreateTask(task model.Task) (model.Task, error) {
	taskId, err := pu.repository.CreateTask(task)
	if err != nil {
		return model.Task{}, err
	}
	task.Id = taskId
	return task, nil
}

func (pu *TaskUseCase) DeleteTask(id int) error {
	err := pu.repository.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (pu *TaskUseCase) DeleteTasks() error {
	err := pu.repository.DeleteTasks()
	if err != nil {
		return err
	}
	return nil
}

func (pu *TaskUseCase) UpdateTask(task model.Task) (int, error) {
	id, err := pu.repository.UpdateTask(task)
	if err != nil {
		return 0, err
	}
	return id, nil
}
