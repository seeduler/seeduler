package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/seeduler/seeduler/services"
)

type EventController struct {
	EventService *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{EventService: service}
}

func (c *EventController) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all events (in controller)")
	events, err := c.EventService.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}
