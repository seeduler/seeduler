package routes

import (
    "net/http"

    "github.com/seeduler/seeduler/controllers"
)

func RegisterUserRoutes(mux *http.ServeMux, userController *controllers.UserController) {
    mux.HandleFunc("/authenticate", userController.Authenticate)
}