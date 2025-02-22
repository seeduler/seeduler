package services

import (
	"errors"
	"log"

	"github.com/seeduler/seeduler/models"
	"github.com/seeduler/seeduler/repositories"
)

type EventService struct {
	EventRepository *repositories.EventRepository
}

func NewEventService(eventRepository *repositories.EventRepository) *EventService {
	return &EventService{EventRepository: eventRepository}
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
