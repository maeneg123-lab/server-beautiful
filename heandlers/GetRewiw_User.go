package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    //"strconv"

    "studia-of-biautiful-api/db"
    //"studia-of-biautiful-api/models"
)

func GetRewiew_User(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    username := r.URL.Query().Get("username")
    if username == ""{
        http.Error(w, "error, username not found", http.StatusBadRequest)
    }
    quotes, err := db.GetRewiew_User(dbConn, username)
    if err!=nil{
        http.Error(w, "error db", http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(quotes)
}
