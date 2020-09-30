package app

import (
	"time"

	"github.com/wailsapp/wails"
)

// Stats wraps a few statistics to monitor and display.
type Stats struct {
	WaterDrank    float32
	TasksComplete int
	TimeWorking   time.Duration
	TimeChilling  time.Duration
}

// NewStats returns an initialized Stats.
func NewStats() *Stats {
	return &Stats{
		WaterDrank:    0,
		TasksComplete: 0,
		TimeWorking:   time.Duration(0),
		TimeChilling:  time.Duration(0),
	}
}

// AddWater increments the water of the day by an amount.
func (s *Stats) AddWater(amt float32) *Stats {
	s.WaterDrank += amt
	return s
}

// CompletedTask increments the number of completed tasks.
func (s *Stats) CompletedTask() *Stats {
	s.TasksComplete++
	return s
}

// IncompletedTask decrements the number of completed tasks.
func (s *Stats) IncompletedTask() *Stats {
	s.TasksComplete--
	return s
}

// AddWorkTime increments the amount of time worked by a number of seconds.
func (s *Stats) AddWorkTime(t int) *Stats {
	s.TimeWorking += (time.Duration(t) * time.Second)
	return s
}

// AddChillTime increments the amount of time chilling by a number of seconds.
func (s *Stats) AddChillTime(t int) *Stats {
	s.TimeChilling += (time.Duration(t) * time.Second)
	return s
}

// Get lets the frontend grab the Stats.
func (s *Stats) Get() *Stats {
	return s
}

// WailsInit sets listeners for events emitted by the Timer.
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	// listen for event the Timer will emit
	runtime.Events.On("add-work-time", func(data ...interface{}) {
		addTime, ok := data[0].(int)
		if ok {
			s.TimeWorking += (time.Duration(addTime) * time.Second)
		}
	})

	// listen for event the Timer will emit
	runtime.Events.On("add-chill-time", func(data ...interface{}) {
		addTime, ok := data[0].(int)
		if ok {
			s.TimeChilling += (time.Duration(addTime) * time.Second)
		}
	})

	return nil
}
