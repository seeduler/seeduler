package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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

func (c *EventController) MarkEventCompleted(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EventID int `json:"event_id"`
	}
	if err := utils.DecodeRequestBody(r, &req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.EventService.MarkEventCompleted(req.EventID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *EventController) MarkEventUncompleted(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EventID int `json:"event_id"`
	}
	if err := utils.DecodeRequestBody(r, &req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.EventService.MarkEventUncompleted(req.EventID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *EventController) AddDelay(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EventID int           `json:"event_id"`
		Delay   time.Duration `json:"delay"`
	}
	if err := utils.DecodeRequestBody(r, &req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.EventService.AddDelay(req.EventID, req.Delay); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *EventController) UpdateDelay(w http.ResponseWriter, r *http.Request) {
	if err := c.EventService.UpdateDelay(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
