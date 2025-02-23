package utils

import (
	"sort"
	"time"

	"github.com/seeduler/seeduler/models"
)

func ProcessEvents(events []models.Event, delay int) []models.Event {

	// Sort events by end time in increasing order
	sort.Slice(events, func(i, j int) bool {
		return events[i].EndTime.Before(events[j].EndTime)
	})

	// Remove completed events
	filteredEvents := []models.Event{}
	now := time.Now()
	for _, event := range events {
		if event.EndTime.After(now) {
			filteredEvents = append(filteredEvents, event)
		}
	}
	events = filteredEvents

	// Add delay to the first event's start and end time
	if len(events) > 0 {
		events[0].StartTime = events[0].StartTime.Add(time.Duration(delay) * time.Minute)
		events[0].EndTime = events[0].EndTime.Add(time.Duration(delay) * time.Minute)
	}

	// Adjust subsequent events
	for i := 1; i < len(events); i++ {
		timeDiff := events[i].StartTime.Sub(events[i-1].EndTime)
		if timeDiff > 0 {
			delay -= int(timeDiff.Minutes())
			if delay < 0 {
				delay = 0
			}
		}
		events[i].StartTime = events[i].StartTime.Add(time.Duration(delay) * time.Minute)
		events[i].EndTime = events[i].EndTime.Add(time.Duration(delay) * time.Minute)
	}
	return events
}
