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

type HallController struct {
	HallService  *services.HallService
	EventService *services.EventService
	UserService  *services.UserService
}

func NewHallController(hallService *services.HallService, eventService *services.EventService, userService *services.UserService) *HallController {
	return &HallController{HallService: hallService, EventService: eventService, UserService: userService}
}

func (c *HallController) GetAllHalls(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all halls (in controller)")

	halls, err := c.HallService.GetAllHalls()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(halls) == 0 {
		http.Error(w, "No halls found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(halls)
}

func (c *HallController) GetHallsWithEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting halls with events (in controller)")

	halls, err := c.HallService.GetAllHalls()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(halls) == 0 {
		http.Error(w, "No halls found", http.StatusNotFound)
		return
	}

	events, err := c.EventService.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(events) == 0 {
		http.Error(w, "No events found", http.StatusNotFound)
		return
	}

	currentTime := time.Now()
	hallsWithEvents := make(map[int][]models.Event)
	updatedEventsCount := 0

	for _, event := range events {
		for _, hall := range halls {
			if event.HallId == hall.ID {
				delayedEndTime := event.EndTime.Add(time.Duration(hall.DelayedTime) * time.Minute)
				if delayedEndTime.After(currentTime) {
					hallsWithEvents[hall.ID] = append(hallsWithEvents[hall.ID], event)
					updatedEventsCount++
				}
			}
		}
	}

	if len(hallsWithEvents) == 0 {
		http.Error(w, "No events found for the halls", http.StatusNotFound)
		return
	}

	log.Printf("Updated events count: %d", updatedEventsCount)

	response := make([]models.HallWithEvents, 0)
	for _, hall := range halls {
		response = append(response, models.HallWithEvents{
			Hall:   hall,
			Events: hallsWithEvents[hall.ID],
		})
	}

	json.NewEncoder(w).Encode(response)
}

func (c *HallController) UploadData(w http.ResponseWriter, r *http.Request) {
	log.Println("Uploading data (in controller)")

	var data struct {
		Halls  []models.Hall  `json:"halls"`
		Events []models.Event `json:"events"`
	}
	if err := utils.DecodeRequestBody(r, &data); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Save halls and events data
	if err := c.HallService.SaveHalls(data.Halls); err != nil {
		http.Error(w, "Error saving halls", http.StatusInternalServerError)
		return
	}
	if err := c.EventService.SaveEvents(data.Events); err != nil {
		http.Error(w, "Error saving events", http.StatusInternalServerError)
		return
	}

	// Create users for each hall
	users := make([]models.User, len(data.Halls))
	for i, hall := range data.Halls {
		users[i] = models.User{
			Username: hall.Name + "_user",
			Password: "password", // Generate a secure password in a real application
			HallID:   hall.ID,
		}
	}

	if err := c.UserService.SaveUsers(users); err != nil {
		http.Error(w, "Error saving users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
