package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
	"todo-api/src/model"

	"github.com/stretchr/testify/assert"
)

var jsonPayload = map[string]any{
	"title":          "Study golang",
	"description":    "super cool",
	"created_at":     time.Now(),
	"is_completed":   false,
	"reward_in_sats": 0,
}

func TestHealth(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var result struct {
		Status string `json:"status"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if result.Status != "OK" {
		t.Errorf("Expected status to be 'OK', got '%s'", result.Status)
	}
}

func TestCreateTask(t *testing.T) {
	// Arrange
	cleanDB(t)

	// Act
	task := createTask(t, jsonPayload)

	// Assert
	assert.Equal(t, jsonPayload["title"], task.Title)
	assert.Equal(t, jsonPayload["description"], task.Description.Val)
	assert.Equal(t, jsonPayload["is_completed"], task.IsCompleted)
	assert.Equal(t, jsonPayload["reward_in_sats"], task.RewardInSats)
}

func TestGetEmptyTaskList(t *testing.T) {
	// Arrange
	cleanDB(t)

	// Act
	tasks := getTasks(t)

	// Assert
	assert.Empty(t, tasks)
}

func TestGetTasks(t *testing.T) {
	// Arrange
	cleanDB(t)

	// Act
	task := createTask(t, jsonPayload)
	tasks := getTasks(t)

	// Assert
	assert.NotEmpty(t, tasks)
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, jsonPayload["title"], task.Title)
	assert.Equal(t, jsonPayload["description"], task.Description.Val)
	assert.Equal(t, jsonPayload["is_completed"], task.IsCompleted)
	assert.Equal(t, jsonPayload["reward_in_sats"], task.RewardInSats)
}

func TestGetTasksById(t *testing.T) {
	// Arrange
	cleanDB(t)

	// Act
	task := createTask(t, jsonPayload)
	returnedTask := getTasksById(t, task.Id)

	// Assert
	assert.NotEmpty(t, returnedTask)
	assert.Equal(t, task.Id, returnedTask.Id)
	assert.Equal(t, jsonPayload["title"], task.Title)
	assert.Equal(t, jsonPayload["description"], task.Description.Val)
	assert.Equal(t, jsonPayload["is_completed"], task.IsCompleted)
	assert.Equal(t, jsonPayload["reward_in_sats"], task.RewardInSats)
}

func TestUpdateTask(t *testing.T) {
	// Arrange
	cleanDB(t)

	// Act
	task := createTask(t, jsonPayload)
	var customJsonPayload = map[string]any{
		"id":             task.Id,
		"title":          "Study kubernets",
		"description":    "awesome",
		"created_at":     time.Now(),
		"is_completed":   false,
		"reward_in_sats": 0,
	}
	updatedTask := updateTask(t, customJsonPayload)

	// Assert
	assert.NotEmpty(t, updatedTask)
	assert.Equal(t, jsonPayload["title"], task.Title)
	assert.Equal(t, jsonPayload["description"], task.Description.Val)
	assert.Equal(t, jsonPayload["is_completed"], task.IsCompleted)
	assert.Equal(t, jsonPayload["reward_in_sats"], task.RewardInSats)
}

func TestDeleteTask(t *testing.T) {
	// Arrange
	cleanDB(t)

	// Act
	task := createTask(t, jsonPayload)
	deleteTask(t, task.Id)
	tasks := getTasks(t)

	// Assert
	assert.Empty(t, tasks)
}

func cleanDB(t *testing.T) {
	var client = &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/tasks", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code 204, got %d", resp.StatusCode)
	}
}

func deleteTask(t *testing.T, id int) {
	var client = &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/task/%d", id), nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code 204, got %d", resp.StatusCode)
	}
}

func createTask(t *testing.T, jsonPayload map[string]any) model.Task {
	payload, _ := json.Marshal(jsonPayload)
	resp, err := http.Post("http://localhost:8080/task", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var task model.Task

	err = json.Unmarshal(body, &task)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	return task
}

func getTasks(t *testing.T) []model.Task {
	resp, err := http.Get("http://localhost:8080/tasks")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var tasks []model.Task

	err = json.Unmarshal(body, &tasks)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	return tasks
}

func getTasksById(t *testing.T, id int) model.Task {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/task/%d", id))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var task model.Task

	err = json.Unmarshal(body, &task)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	return task
}

func updateTask(t *testing.T, jsonPayload map[string]any) model.Task {
	client := &http.Client{}

	payload, _ := json.Marshal(jsonPayload)
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/task/%d", jsonPayload["id"]), bytes.NewBuffer(payload))
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var task model.Task

	err = json.Unmarshal(body, &task)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	return task
}
