package main

import (
	"log"
	"net/http"

	"github.com/seeduler/seeduler/controllers"
	"github.com/seeduler/seeduler/middlewares"
	"github.com/seeduler/seeduler/repositories"
	"github.com/seeduler/seeduler/routes"
	"github.com/seeduler/seeduler/services"
	"github.com/seeduler/seeduler/utils"
)

func main() {
	mux := http.NewServeMux()

	eventRepo := repositories.NewEventRepository("data/event.json")
	hallRepo := repositories.NewHallRepository("data/hall.json")
	userRepo := repositories.NewUserRepository("data/user.json")

	// Load configuration
	config, err := utils.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	eventService := services.NewEventService(eventRepo, hallRepo)
	hallService := services.NewHallService(hallRepo)
	userService := services.NewUserService(userRepo, []byte(config.JWT.SecretKey))

	eventController := controllers.NewEventController(eventService)
	hallController := controllers.NewHallController(hallService, eventService, userService)
	userController := controllers.NewUserController(userService)

	// Create a chain of middlewares
	handler := middlewares.CorsMiddleware(
		middlewares.HallCheckMiddleware(hallService)(mux),
	)

	// Register all routes
	routes.RegisterRoutes(mux, eventController, hallController, userController, userService)

	log.Printf("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
