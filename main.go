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
	hallRepo := repositories.NewHallRepository("data/hall.json")
	eventService := services.NewEventService(eventRepo, hallRepo)
	eventController := controllers.NewEventController(eventService)

	hallService := services.NewHallService(hallRepo)
	hallController := controllers.NewHallController(hallService, eventService)

	// Register all routes
	routes.RegisterRoutes(mux, eventController, hallController)

	log.Printf("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
