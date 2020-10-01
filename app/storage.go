package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Storage interface {
	Store(*Entry) error
	Reset() error
	AddTask(*Task) error
	UpdateStats(*Stats) error
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

func (s *storage) UpdateStats(stats *Stats) error {
	entry, ok := s.Entries[time.Now().Format("2006/01/02")]
	if !ok {
		return errors.Errorf("No entry")
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
