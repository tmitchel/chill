package app

import (
	"time"

	"github.com/pkg/errors"
)

// Task holds a single task to be completed.
type Task struct {
	ID        int
	Content   string
	CreatedAt time.Time
	Completed bool
}

// Tasks gives a type for a slice of Task's
type Tasks []*Task

// NewTasks initializes the slice and returns a pointer
// to the Tasks
func NewTasks() *Tasks {
	t := make([]*Task, 0)
	task := Tasks(t)
	return &task
}

// Tasks returns the slice of tasks.
func (t *Tasks) Tasks() []*Task {
	return []*Task(*t)
}

// Create makes a new task and adds to the slice.
func (t *Tasks) Create(content string) *Task {
	task := &Task{
		ID:        len(*t),
		Content:   content,
		CreatedAt: time.Now(),
		Completed: false,
	}
	*t = append(*t, task)
	return task
}

// GetTask returns the task with the given ID.
func (t *Tasks) GetTask(id int) (*Task, error) {
	if len(*t) <= id {
		return nil, errors.Errorf("No task with id: %v", id)
	}
	return (*t)[id], nil
}

// ToggleStatus flips the Completed bool.
func (t *Tasks) ToggleStatus(id int) (*Task, error) {
	if len(*t) <= id {
		return nil, errors.Errorf("No task with id: %v", id)
	}

	(*t)[id].Completed = !(*t)[id].Completed
	return (*t)[id], nil
}
