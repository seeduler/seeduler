package controllers

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/seeduler/seeduler/models"
    "github.com/seeduler/seeduler/services"
)

type HallController struct {
    HallService  *services.HallService
    EventService *services.EventService
}

func NewHallController(hallService *services.HallService, eventService *services.EventService) *HallController {
    return &HallController{HallService: hallService, EventService: eventService}
}

func (c *HallController) GetAllHalls(w http.ResponseWriter, r *http.Request) {
    log.Println("Getting all halls (in controller)")

    halls, err := c.HallService.GetAllHalls()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
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

    events, err := c.EventService.GetAllEvents()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    currentTime := time.Now()
    hallsWithEvents := make(map[int][]models.Event)

    for _, event := range events {
        for _, hall := range halls {
            if event.HallId == hall.ID {
                delayedEndTime := event.EndTime.Add(time.Duration(hall.DelayedTime) * time.Minute)
                if delayedEndTime.After(currentTime) {
                    hallsWithEvents[hall.ID] = append(hallsWithEvents[hall.ID], event)
                }
            }
        }
    }

    response := make([]models.HallWithEvents, 0)
    for _, hall := range halls {
        response = append(response, models.HallWithEvents{
            Hall:   hall,
            Events: hallsWithEvents[hall.ID],
        })
    }

    json.NewEncoder(w).Encode(response)
}