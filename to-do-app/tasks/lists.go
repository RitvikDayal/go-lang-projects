package tasks

import (
	"log"

	"github.com/go/to-do-app/database"
	"github.com/go/to-do-app/database/models"
)

func AddList(list models.List) {
	db := database.GetDbConnection()
	defer db.Close()

	result, err := db.Exec(
		"INSERT INTO lists (name, created_at) VALUES (?, ?)",
		list.Name,
		list.CreatedAt,
	)
	if err != nil {
		log.Fatalf("Error while adding list: %s\n", err)
	}

	// Get the ID of the newly created list
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error while getting the ID of the newly created list: %s\n", err)
	}

	log.Printf("List added successfully with ID: %d\n", id)
}

func GetTasksByListID(listID int) []models.Task {
	db := database.GetDbConnection()
	defer db.Close()

	rows, err := db.Query(
		"SELECT * FROM tasks WHERE list_id = ?", listID)
	if err != nil {
		log.Fatalf("Error while fetching tasks by list ID: %s\n", err)
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
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
			log.Fatalf("Error while scanning task: %s\n", err)
		}
		tasks = append(tasks, task)
	}

	return tasks
}

