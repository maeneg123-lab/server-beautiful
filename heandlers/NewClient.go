package handlers

import(
    "database/sql"
    //"encoding/json"
    "fmt"
    //"log"
    "net/http"
    //"os"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "studia-of-biautiful-api/db"     // или "weather-go-api/db"
)

func NewClient(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    value := r.URL.Query()

    username := value.Get("username")
    user := value.Get("user_id")

    if username == ""{
        http.Error(w, "error, username no", http.StatusBadRequest)
        return
    }

    err := db.NewClient(dbConn, username, user)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! client added!")
}
