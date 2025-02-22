package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
)

func RegisterRoutes(mux *http.ServeMux, eventController *controllers.EventController) {
    RegisterEventRoutes(mux, eventController)
    // Future routes can be registered here
}
