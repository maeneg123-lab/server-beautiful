package models

import(
 "time"
)

type ClientResponse struct{
    ID int `json: "id"`
    User_id string `json: "user"`
    Username string `json: "username"`
    CreatedAt time.Time `json: "created_at"`
}

type RatingResponse struct{
    ID int `json: "id"`
    Username string `json: "username"`
    Text string `json: "text"`
    CreatedAt time.Time `json: "created_at"`
    Rating float64 `json: "rating"`
    Operation string `json: "operation"`
}

type Slot struct {
    ID        int       `json:"id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Status    string    `json:"status"`
    ClientID  *int      `json:"client_id,omitempty"` // может быть NULL
}
