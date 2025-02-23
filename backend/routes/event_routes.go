package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
    "github.com/seeduler/seeduler/middlewares"
    "github.com/seeduler/seeduler/services"
)

func RegisterEventRoutes(mux *http.ServeMux, eventController *controllers.EventController, userService *services.UserService) {
    authMiddleware := middlewares.AuthMiddleware(userService)

    mux.HandleFunc("/events/by-hall-ids", eventController.GetHallEvents)
    mux.HandleFunc("/events/first-event-of-each-hall", eventController.GetFirstEventOfEachHall)
    mux.Handle("/events/mark-completed", authMiddleware(http.HandlerFunc(eventController.MarkEventCompleted)))
    mux.Handle("/events/mark-uncompleted", authMiddleware(http.HandlerFunc(eventController.MarkEventUncompleted)))
    mux.Handle("/events/add-delay", authMiddleware(http.HandlerFunc(eventController.AddDelay)))
    mux.Handle("/events/update-delay", authMiddleware(http.HandlerFunc(eventController.UpdateDelay)))
    mux.Handle("/events/add", authMiddleware(http.HandlerFunc(eventController.AddEvent)))
    mux.Handle("/events/remove", authMiddleware(http.HandlerFunc(eventController.RemoveEvent)))
}
