package main

import (
	"log"
	"net/http"
	"os"

	"github.com/seeduler/seeduler/controllers"
	"github.com/seeduler/seeduler/middlewares"
	"github.com/seeduler/seeduler/repositories"
	"github.com/seeduler/seeduler/routes"
	"github.com/seeduler/seeduler/services"
)

func main() {
	mux := http.NewServeMux()

	eventRepo := repositories.NewEventRepository("data/event.json")
	hallRepo := repositories.NewHallRepository("data/hall.json")
	userRepo := repositories.NewUserRepository("data/user.json")

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	eventService := services.NewEventService(eventRepo, hallRepo)
	hallService := services.NewHallService(hallRepo)
	userService := services.NewUserService(userRepo, jwtKey)

	eventController := controllers.NewEventController(eventService)
	hallController := controllers.NewHallController(hallService, eventService, userService)
	userController := controllers.NewUserController(userService)

	// Register all routes without middleware
	routes.RegisterRoutes(mux, eventController, hallController, userController)

	// Add authentication middleware to all routes except /halls/upload-data and user-related routes
	authMiddleware := middlewares.AuthMiddleware(userService)
	protectedMux := http.NewServeMux()
	protectedMux.Handle("/", authMiddleware(mux))

	// Register /halls/upload-data and user-related routes without middleware
	protectedMux.Handle("/halls/upload-data", http.HandlerFunc(hallController.UploadData))
	protectedMux.Handle("/authenticate", http.HandlerFunc(userController.Authenticate))

	log.Printf("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", protectedMux))
}
