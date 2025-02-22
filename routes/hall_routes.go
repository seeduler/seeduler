package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
)

func RegisterHallRoutes(mux *http.ServeMux, hallController *controllers.HallController) {
    mux.HandleFunc("/halls", hallController.GetAllHalls)
    mux.HandleFunc("/halls/with-events", hallController.GetHallsWithEvents)
}