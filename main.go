package main

import (
	"log"
	"net/http"

	"github.com/seeduler/seeduler/controllers"
	"github.com/seeduler/seeduler/repositories"
	"github.com/seeduler/seeduler/routes"
	"github.com/seeduler/seeduler/services"
)

func main() {
	mux := http.NewServeMux()

	eventRepo := repositories.NewEventRepository("data/event.json")
	eventService := services.NewEventService(eventRepo)
	eventController := controllers.NewEventController(eventService)
	routes.RegisterEventRoutes(mux, eventController)

	log.Printf("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
