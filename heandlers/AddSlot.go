package handlers

import(
    "database/sql"
    //"encoding/json"
    "fmt"
    //"log"
    "net/http"
    //"os"
    //"strconv"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "studia-of-biautiful-api/db"     // или "weather-go-api/db"
)

func AddSlot(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    value := r.URL.Query()

    start_time := value.Get("start_time")
    end_time := value.Get("end_time")
    status := value.Get("status")

    if start_time == ""{
        http.Error(w, "error, start_time no", http.StatusBadRequest)
        return
    }

    if end_time == ""{
        http.Error(w, "error, end_time no", http.StatusBadRequest)
        return
    }

    err := db.AddSlot(dbConn, start_time, end_time, status)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! slot added!")
}
