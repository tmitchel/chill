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
type Tasks struct {
	T     []*Task `json:"tasks"`
	Store Storage `json:"-"`
}

// NewTasks initializes the slice and returns a pointer
// to the Tasks
func NewTasks(store Storage) *Tasks {
	return &Tasks{
		T:     make([]*Task, 0),
		Store: store,
	}
}

// Tasks returns the slice of tasks.
func (t *Tasks) Tasks() []*Task {
	return t.T
}

// Create makes a new task and adds to the slice.
func (t *Tasks) Create(content string) *Task {
	task := &Task{
		ID:        len(t.T),
		Content:   content,
		CreatedAt: time.Now(),
		Completed: false,
	}
	t.T = append(t.T, task)
	t.Store.AddTask(task)
	return task
}

// GetTask returns the task with the given ID.
func (t *Tasks) GetTask(id int) (*Task, error) {
	if len(t.T) <= id {
		return nil, errors.Errorf("No task with id: %v", id)
	}
	return t.T[id], nil
}

// ToggleStatus flips the Completed bool.
func (t *Tasks) ToggleStatus(id int) (*Task, error) {
	if len(t.T) <= id {
		return nil, errors.Errorf("No task with id: %v", id)
	}

	t.T[id].Completed = !t.T[id].Completed
	return t.T[id], nil
}
