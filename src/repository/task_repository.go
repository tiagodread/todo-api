package repository

import (
	"database/sql"
	"fmt"
	"todo-api/src/model"
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
		return []model.Task{}, err
	}
	defer rows.Close()

	var taskList []model.Task
	var taskObj model.Task

	if !rows.Next() { // Verifica se há pelo menos uma linha retornada
		fmt.Println("Nenhuma tarefa encontrada")
		return []model.Task{}, nil
	}

	// Se houve linhas, processa a primeira
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
		return []model.Task{}, err
	}
	taskList = append(taskList, taskObj)

	// Processa as demais linhas
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
			return []model.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return []model.Task{}, err
	}

	return taskList, nil
}

func (pr *TaskRepository) GetTask(id int) (model.Task, error) {
	var task model.Task
	query := "SELECT * FROM task WHERE id=$1"
	row := pr.connection.QueryRow(query, id)
	switch err := row.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt, &task.CompletedAt, &task.IsCompleted, &task.RewardInSats); err {
	case sql.ErrNoRows:
		return model.Task{}, sql.ErrNoRows
	}

	return task, nil
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

func (pr *TaskRepository) DeleteTask(id int) error {
	query := fmt.Sprintf("DELETE FROM task WHERE id = %d", id)

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Print(rows)
	return nil
}

func (pr *TaskRepository) UpdateTask(task model.Task) (int, error) {
	var id int

	query, err := pr.connection.Prepare("UPDATE task " +
		"SET title = $1, description = $2, created_at = $3, is_completed = $4, reward_in_sats = $5" +
		" WHERE id = $6 RETURNING id")
	if err != nil {
		fmt.Println(err)
	}

	err = query.QueryRow(task.Title, task.Description, task.CreatedAt, task.IsCompleted, task.RewardInSats, task.Id).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}
