package commands

import (
	"log"
	"time"
	"strconv"

	"github.com/go/to-do-app/database/models"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "to-do-app",
	Short: "A simple to-do app",
	Long:  `A simple to-do app that allows you to add, list, and delete tasks.`,
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Listing all tasks")
		tasks := listAllTasks()
		// print task in a table format with equal spacing
		log.Printf("%-2s | %-20s | %-55s | %-35s | %-5s | %-5s\n", "ID", "Title", "Description", "Due Date", "Completed", "Priority")
		for _, task := range tasks {
			log.Printf("%-2d | %-20s | %-55s | %-35s | %-5t | %-5d\n", task.Id, task.Title, task.Description, task.DueDate, task.Completed, task.Priority)
		}
	},
}

var AddCmd = &cobra.Command{
	Use:   "add-task",
	Short: "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Adding a task")
		dueDate, err := time.Parse(time.RFC3339, cmd.Flag("due-date").Value.String())
		if err != nil {
			log.Fatal(err)
		}
		is_completed, err := strconv.ParseBool(cmd.Flag("completed").Value.String())
		if err != nil {
			log.Fatal(err)
		}
		list_id,err := strconv.Atoi(cmd.Flag("list-id").Value.String())
		if err != nil {
			log.Fatal(err)
		}

		priority, err := strconv.Atoi(cmd.Flag("priority").Value.String())
		if err != nil {
			log.Fatal(err)
		}

		task := models.Task{
			Title:       cmd.Flag("title").Value.String(),
			Description: cmd.Flag("description").Value.String(),
			DueDate:     dueDate,
			Completed:   is_completed,
			ListID:      list_id,
			Priority:    priority,
			CreatedAt:   time.Now(),
		}
		addTask(task)
	},
}
