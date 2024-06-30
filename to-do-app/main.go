/*
To list app
---
author: ritvikdayal
*/

package main

import (
	"log"
	"time"

	"github.com/go/to-do-app/database"
	"github.com/go/to-do-app/database/models"
	"github.com/go/to-do-app/tasks"
)

func main() {
	database.InitDB()

	// Add a list
	log.Println("Adding a list...")
	list := models.List{
		Name:      "Work",
		CreatedAt: time.Now(),
	}
	tasks.AddList(list)

	// Add a task
	log.Println("Adding a task...")
	task := models.Task{
		ListID:      1,
		Title:       "Write a blog post",
		Description: "Write a blog post on how to use Go to build a to-do app",
		DueDate:     time.Now().Format(time.RFC3339),
		Completed:   false,
		CreatedAt:   time.Now().Format(time.RFC3339),
		Priority:    1,
	}
	tasks.AddTask(task)

	// Get a task by ID
	log.Println("Getting a task by ID...")
	taskByID := tasks.GetTaskByID(1)
	log.Printf("Task: %+v\n", taskByID)

	// Get tasks by list ID
	log.Println("Getting tasks by list ID...")
	tasksByListID := tasks.GetTasksByListID(1)
	log.Printf("Tasks: %+v\n", tasksByListID)
}
