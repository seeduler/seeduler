package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
)

func RegisterEventRoutes(mux *http.ServeMux, eventController *controllers.EventController) {
    mux.HandleFunc("/events/by-hall-ids", eventController.GetHallEvents)
}
