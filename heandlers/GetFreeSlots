package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    //"strconv"

    "studia-of-biautiful-api/db"
    //"studia-of-biautiful-api/models"
)

func GetFreeSlots(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    quotes, err := db.GetFreeSlots(dbConn)
    if err!=nil{
        http.Error(w, "error db", http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(quotes)
}
