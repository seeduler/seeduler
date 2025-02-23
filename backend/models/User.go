package models

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    HallID   int    `json:"hall_id"`
}