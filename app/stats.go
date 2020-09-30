package app

import (
	"time"

	"github.com/wailsapp/wails"
)

type Stats struct {
	WaterDrank    float32
	TasksComplete int
	TimeWorking   time.Duration
	TimeChilling  time.Duration
}

func NewStats() *Stats {
	return &Stats{
		WaterDrank:    0,
		TasksComplete: 0,
		TimeWorking:   time.Duration(0),
		TimeChilling:  time.Duration(0),
	}
}

func (s *Stats) AddWater(amt float32) *Stats {
	s.WaterDrank += amt
	return s
}

func (s *Stats) CompletedTask() *Stats {
	s.TasksComplete++
	return s
}

func (s *Stats) IncompletedTask() *Stats {
	s.TasksComplete--
	return s
}

func (s *Stats) AddWorkTime(t int) *Stats {
	s.TimeWorking += (time.Duration(t) * time.Second)
	return s
}

func (s *Stats) AddChillTime(t int) *Stats {
	s.TimeChilling += (time.Duration(t) * time.Second)
	return s
}

func (s *Stats) Get() *Stats {
	return s
}

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
