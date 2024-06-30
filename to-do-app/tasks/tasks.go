package tasks

import (
	"log"

	"github.com/go/to-do-app/database"
	"github.com/go/to-do-app/database/models"
)

func GetTaskByID(id int) models.Task {
	db := database.GetDbConnection()
	defer db.Close()

	var task models.Task
	err := db.QueryRow(
		"SELECT * FROM tasks WHERE id = ?", id).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.DueDate,
		&task.Completed,
		&task.ListID,
		&task.Priority,
		&task.CreatedAt,
	)
	if err != nil {
		log.Fatalf("Error while fetching task by ID: %s\n", err)
	}
	return task
}

func AddTask(task models.Task) {
	db := database.GetDbConnection()
	defer db.Close()

	result, err := db.Exec(
		"INSERT INTO tasks (list_id, title, description, due_date, completed, created_at, priority) VALUES (?, ?, ?, ?, ?, ?, ?)",
		task.ListID,
		task.Title,
		task.Description,
		task.DueDate,
		task.Completed,
		task.CreatedAt,
		task.Priority,
	)
	if err != nil {
		log.Fatalf("Error while adding task: %s\n", err)
	}

	// Retrieve the auto-incremented ID
	taskId, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error while retrieving last insert ID: %s\n", err)
	}

	// Optionally, use the taskId for further processing
	log.Printf("Task added with ID: %d\n", taskId)
}
