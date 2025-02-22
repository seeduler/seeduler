package models

import (
    "time"
    "encoding/json"
)

type Event struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    HallId    int       `json:"hall_id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Info      Info      `json:"info"`
}

type Info struct {
    // info about the event
    // to be changed later
}

// Custom time format for JSON marshaling/unmarshaling
const timeFormat = time.RFC3339

func (t *Event) MarshalJSON() ([]byte, error) {
    type Alias Event
    return json.Marshal(&struct {
        StartTime string `json:"start_time"`
        EndTime   string `json:"end_time"`
        *Alias
    }{
        StartTime: t.StartTime.Format(timeFormat),
        EndTime:   t.EndTime.Format(timeFormat),
        Alias:     (*Alias)(t),
    })
}

func (t *Event) UnmarshalJSON(data []byte) error {
    type Alias Event
    aux := &struct {
        StartTime string `json:"start_time"`
        EndTime   string `json:"end_time"`
        *Alias
    }{
        Alias: (*Alias)(t),
    }
    if err := json.Unmarshal(data, aux); err != nil {
        return err
    }
    var err error
    t.StartTime, err = time.Parse(timeFormat, aux.StartTime)
    if err != nil {
        return err
    }
    t.EndTime, err = time.Parse(timeFormat, aux.EndTime)
    if err != nil {
        return err
    }
    return nil
}
