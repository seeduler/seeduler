package routes

import (
	"net/http"

	"github.com/seeduler/seeduler/controllers"
	"github.com/seeduler/seeduler/services"
)

func RegisterRoutes(mux *http.ServeMux, eventController *controllers.EventController, hallController *controllers.HallController, userController *controllers.UserController, userService *services.UserService) {
	RegisterEventRoutes(mux, eventController, userService)
	RegisterHallRoutes(mux, hallController)
	RegisterUserRoutes(mux, userController)
	// Future routes can be registered here
}
