package models

import "time"

// Mode of a List
type List struct {
	Name      string
	CreatedAt time.Time
}

// Mode of a Task
type Task struct {
	Id 		int
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
	ListID      int
	Priority    int
	CreatedAt   time.Time
}

// Mode of a tag
type Tag struct {
	Name string
}

// Mode of a task_tag
type TaskTag struct {
	TaskId int
	TagId  int
}

// Tasks of a List
type ListTasks struct {
	List
	Tasks []Task
}

// Tags of a Task
type TaskTags struct {
	Task
	Tags []Tag
}
