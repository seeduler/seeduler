package services

import (
	"errors"
	"log"
	"time"

	"github.com/seeduler/seeduler/models"
	"github.com/seeduler/seeduler/repositories"
	"github.com/seeduler/seeduler/utils"
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
	hall, err := s.HallRepository.GetHallByID(req.HallIds[0])
	if err != nil {
		return resp, err
	}
	delay := hall.DelayedTime
	resp = utils.ProcessEvents(resp, delay)
	return
}

func (s *EventService) GetFirstEventOfEachHall() ([]models.Event, error) {
	log.Println("Getting first event of each hall (in service)")
	events, err := s.EventRepository.GetEvents()
	if err != nil {
		return nil, err
	}

	halls, err := s.HallRepository.GetHalls()
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	firstEvents := make([]models.Event, 0)

	for _, hall := range halls {
		var firstEvent *models.Event
		for _, event := range events {
			if event.HallId == hall.ID && !event.IsCompleted {
				delayedEndTime := event.EndTime.Add(time.Duration(hall.DelayedTime) * time.Minute)
				if delayedEndTime.After(currentTime) {
					if firstEvent == nil || delayedEndTime.Before(firstEvent.EndTime) {
						firstEvent = &event
						firstEvent.EndTime = delayedEndTime
						firstEvent.StartTime = firstEvent.StartTime.Add(time.Duration(hall.DelayedTime) * time.Minute)
					}
				}
			}
		}

		if firstEvent != nil {
			adjustedEvent := *firstEvent
			adjustedEvent.StartTime = adjustedEvent.StartTime.Add(time.Duration(hall.DelayedTime) * time.Minute)
			adjustedEvent.EndTime = adjustedEvent.EndTime.Add(time.Duration(hall.DelayedTime) * time.Minute)
			firstEvents = append(firstEvents, adjustedEvent)
		}
	}

	return firstEvents, nil
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

func (s *EventService) SaveEvents(events []models.Event) error {
	log.Println("Saving events (in service)")
	return s.EventRepository.SaveEvents(events)
}
