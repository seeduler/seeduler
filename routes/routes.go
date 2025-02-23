package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
)

func RegisterRoutes(mux *http.ServeMux, eventController *controllers.EventController, hallController *controllers.HallController, userController *controllers.UserController) {
    RegisterEventRoutes(mux, eventController)
    RegisterHallRoutes(mux, hallController)
    RegisterUserRoutes(mux, userController)
    // Future routes can be registered here
}
