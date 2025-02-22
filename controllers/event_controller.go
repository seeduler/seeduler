package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/seeduler/seeduler/models"
	"github.com/seeduler/seeduler/services"
	"github.com/seeduler/seeduler/utils"
)

type EventController struct {
	EventService *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{EventService: service}
}

func (c *EventController) GetHallEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting events by hall IDs (in controller)")

	hallIds, err := utils.ParseQueryParams(r, "hall_ids")
	if err != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	req := models.GetEventsRequest{HallIds: hallIds}

	events, err := c.EventService.GetEventsByHallIds(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}
