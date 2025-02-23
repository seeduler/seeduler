package services

import (
	"errors"
	"log"
	"time"

	"github.com/seeduler/seeduler/models"
	"github.com/seeduler/seeduler/repositories"
)

type EventService struct {
	EventRepository *repositories.EventRepository
	HallRepository  *repositories.HallRepository
}

func NewEventService(eventRepository *repositories.EventRepository, hallRepository *repositories.HallRepository) *EventService {
	return &EventService{EventRepository: eventRepository, HallRepository: hallRepository}
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	log.Println("Getting all events (in service)")
	return s.EventRepository.GetEvents()
}

func (s *EventService) GetEventsByHallIds(req models.GetEventsRequest) (resp []models.Event, err error) {
	log.Println("Getting all events (in service)")
	events, err := s.EventRepository.GetEvents()
	if err != nil {
		return resp, err
	}
	if len(events) == 0 {
		return resp, errors.New("No events found")
	}
	for _, event := range events {
		for _, hallID := range req.HallIds {
			if event.HallId == hallID {
				resp = append(resp, event)
				break
			}
		}
	}
	return
}

func (s *EventService) MarkEventCompleted(eventID int) error {
	events, err := s.EventRepository.GetEvents()
	if err != nil {
		return err
	}

	for i, event := range events {
		if event.ID == eventID {
			event.IsCompleted = true
			delay := time.Now().Sub(event.EndTime)
			if delay < 0 {
				delay = 0
			}
			event.EndTime = time.Now()
			events[i] = event
			break
		}
	}

	return s.EventRepository.SaveEvents(events)
}

func (s *EventService) MarkEventUncompleted(eventID int) error {
	events, err := s.EventRepository.GetEvents()
	if err != nil {
		return err
	}

	var uncompletedEvent *models.Event
	for i, event := range events {
		if event.ID == eventID {
			event.IsCompleted = false
			uncompletedEvent = &event
			events[i] = event
			break
		}
	}

	if uncompletedEvent != nil {
		delay := time.Now().Sub(uncompletedEvent.EndTime)
		if delay < 0 {
			delay = 0
		}

		for i, event := range events {
			if event.EndTime.After(uncompletedEvent.EndTime) {
				event.IsStarted = false
				events[i] = event
			}
		}

		uncompletedEvent.EndTime = uncompletedEvent.EndTime.Add(delay)
	}

	return s.EventRepository.SaveEvents(events)
}

func (s *EventService) AddDelay(eventID int, delay time.Duration) error {
	events, err := s.EventRepository.GetEvents()
	if err != nil {
		return err
	}

	for i, event := range events {
		if event.ID == eventID {
			event.EndTime = event.EndTime.Add(delay)
			events[i] = event
			break
		}
	}

	return s.EventRepository.SaveEvents(events)
}

func (s *EventService) UpdateDelay() error {
	events, err := s.EventRepository.GetEvents()
	if err != nil {
		return err
	}

	var previousEvent *models.Event
	for _, event := range events {
		if !event.IsCompleted {
			if previousEvent == nil || event.EndTime.Before(previousEvent.EndTime) {
				previousEvent = &event
			}
		}
	}

	if previousEvent != nil {
		delay := time.Now().Sub(previousEvent.EndTime) + time.Minute
		if delay < 0 {
			delay = 0
		}
		previousEvent.EndTime = previousEvent.EndTime.Add(delay)
	}

	return s.EventRepository.SaveEvents(events)
}
