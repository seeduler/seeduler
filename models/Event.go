package models

import "time"

type Event struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Hall      string    `json:"hall"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Info      Info      `json:"info"`
}

type Info struct {
	// info about the event
	// to be changed later
}
