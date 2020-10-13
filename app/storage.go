package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Storage interface {
	Store(*Entry) error
	Reset() error
	AddTask(*Task) error
	UpdateTasks(*Tasks) error
	UpdateStats(*Stats) error
	GetTasks() Tasks
	GetStats() *Stats
}

type storage struct {
	Name    string
	Entries map[string]Entry
}

type Entry struct {
	Date  string
	Stats Stats
	Tasks Tasks
}

// Open checks if the file and reads it if it exists.
// Otherwise, create the empty storage.
func Open(file string) (Storage, error) {
	s := &storage{Name: file}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		s.Entries = make(map[string]Entry)
		return s, nil
	}

	// read the file
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// parse the content
	err = json.Unmarshal(content, &s.Entries)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *storage) GetTasks() Tasks {
	tasks := s.Entries[time.Now().Format("2006/01/02")].Tasks
	for _, e := range s.Entries {
		for _, t := range e.Tasks.T {
			if !t.Completed {
				tasks.T = append(tasks.T, t)
			}
		}
	}
	return tasks
}

func (s *storage) GetStats() *Stats {
	stats, ok := s.Entries[time.Now().Format("2006/01/02")]
	if !ok {
		return nil
	}
	return &stats.Stats
}

func (s *storage) AddTask(task *Task) error {
	entry, ok := s.Entries[time.Now().Format("2006/01/02")]

	// if entry doesn't exist, create it
	if !ok {
		date := time.Now().Format("2006/01/02")
		entry := Entry{
			Date: date,
			Tasks: Tasks{
				T: []*Task{task},
			},
		}
		s.Entries[date] = entry
		s.Store(&entry)
		return nil
	}

	// update existing entry
	entry.Tasks.T = append(entry.Tasks.T, task)
	s.Store(&entry)
	return nil
}

// Reset deletes everything.
func (s *storage) Reset() error {
	s.Entries = make(map[string]Entry)
	return os.Remove(s.Name)
}

func (s *storage) UpdateTasks(tasks *Tasks) error {
	entry, ok := s.Entries[time.Now().Format("2006/01/02")]
	if !ok {
		date := time.Now().Format("2006/01/02")
		entry := Entry{
			Date:  date,
			Tasks: *tasks,
		}
		s.Entries[date] = entry
		s.Store(&entry)
		return nil
	}
	entry.Tasks = *tasks
	s.Store(&entry)
	return nil
}

func (s *storage) UpdateStats(stats *Stats) error {
	entry, ok := s.Entries[time.Now().Format("2006/01/02")]
	if !ok {
		date := time.Now().Format("2006/01/02")
		entry := Entry{
			Date:  date,
			Stats: *stats,
		}
		s.Entries[date] = entry
		s.Store(&entry)
		return nil
	}
	entry.Stats = *stats
	s.Store(&entry)
	return nil
}

// AddEntry persists a new entry to the database.
func (s *storage) Store(e *Entry) error {
	date := time.Now().Format("2006/01/02")
	e.Date = date
	s.Entries[date] = *e

	// write file in the background
	go func() {
		file, err := os.Create(s.Name)
		if err != nil {
			logrus.Fatal(err)
		}
		defer file.Close()

		content, err := json.MarshalIndent(s.Entries, "", "  ")
		if err != nil {
			logrus.Fatal(err)
		}

		_, err = file.Write(content)
	}()
	return nil
}
