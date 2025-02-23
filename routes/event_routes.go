package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
)

func RegisterEventRoutes(mux *http.ServeMux, eventController *controllers.EventController) {
    mux.HandleFunc("/events/by-hall-ids", eventController.GetHallEvents)
    mux.HandleFunc("/events/mark-completed", eventController.MarkEventCompleted)
    mux.HandleFunc("/events/mark-uncompleted", eventController.MarkEventUncompleted)
    mux.HandleFunc("/events/add-delay", eventController.AddDelay)
    mux.HandleFunc("/events/update-delay", eventController.UpdateDelay)
}
