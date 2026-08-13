package handlers

import(
    "database/sql"
    //"encoding/json"
    "fmt"
    //"log"
    "net/http"
    //"os"
    "strconv"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "studia-of-biautiful-api/db"     // или "weather-go-api/db"
)

func AddRewiew(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    value := r.URL.Query()

    username := value.Get("username")
    text := value.Get("text")
    ratingstr := value.Get("rating")
    operation := value.Get("operation")

    if text == ""{
        http.Error(w, "error, text no", http.StatusBadRequest)
        return
    }

    if username == ""{
        http.Error(w, "error, username no", http.StatusBadRequest)
        return
    }

    rating ,err := strconv.ParseFloat(ratingstr, 64)
    if err != nil{
        http.Error(w, "error, client no int", http.StatusBadRequest)
    }

    if operation == ""{
        http.Error(w, "error, operation no", http.StatusBadRequest)
        return
    }



    err = db.AddRewiew(dbConn, username, text, rating, operation)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! rewiew added!")
}
