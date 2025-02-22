package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
)

func RegisterRoutes(mux *http.ServeMux, eventController *controllers.EventController, hallController *controllers.HallController) {
    RegisterEventRoutes(mux, eventController)
    RegisterHallRoutes(mux, hallController)
    // Future routes can be registered here
}
