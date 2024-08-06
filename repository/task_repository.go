package repository

import (
	"database/sql"
	"fmt"
	"todo-api/model"
)

type TaskRepository struct {
	connection *sql.DB
}

func NewTaskRepository(connection *sql.DB) TaskRepository {
	return TaskRepository{
		connection: connection,
	}
}

func (pr *TaskRepository) GetTasks() ([]model.Task, error) {
	query := "SELECT * FROM task"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Task{}, nil
	}

	defer rows.Close()

	var taskList []model.Task
	var taskObj model.Task

	for rows.Next() {
		err = rows.Scan(
			&taskObj.Id,
			&taskObj.Title,
			&taskObj.Description,
			&taskObj.CreatedAt,
			&taskObj.CompletedAt,
			&taskObj.IsCompleted,
			&taskObj.RewardInSats,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Task{}, nil
		}

		taskList = append(taskList, taskObj)
	}
	return taskList, nil
}

func (pr *TaskRepository) CreateTask(task model.Task) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO task" +
		"(title, description, created_at, is_completed, reward_in_sats)" +
		" VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		fmt.Println(err)
	}

	err = query.QueryRow(task.Title, task.Description, task.CreatedAt, task.IsCompleted, task.RewardInSats).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil

}
