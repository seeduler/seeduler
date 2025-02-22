package routes

import (
	"log"
	"net/http"

	"github.com/seeduler/seeduler/controllers"
)

func RegisterEventRoutes(mux *http.ServeMux, eventController *controllers.EventController) {
	log.Println("Registering event routes")
	mux.HandleFunc("/events", eventController.GetAllEvents)
}
