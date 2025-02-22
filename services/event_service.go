package services

import (
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

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	log.Println("Getting all events (in service)")
	return s.EventRepository.GetAllEvents()
}
