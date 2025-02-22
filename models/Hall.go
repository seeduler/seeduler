package models

type Hall struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    DelayedTime int      `json:"delayed_time"` // in minutes
    Info        HallInfo `json:"info"`
}

type HallInfo struct {
    // Additional info about the hall
    // to be changed later
}

type HallWithEvents struct {
    Hall   Hall    `json:"hall"`
    Events []Event `json:"events"`
}