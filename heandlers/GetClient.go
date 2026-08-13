package handlers

import(
    "strconv"
    "database/sql"
    "encoding/json"
    //"fmt"
    //"log"
    "net/http"
    //"os"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "studia-of-biautiful-api/db"     // или "weather-go-api/db"
)

func GetClient(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    idstr := r.URL.Query().Get("id")
    if idstr == ""{
        http.Error(w, "error, id no", http.StatusBadRequest)
        return
    }

    id,err:=strconv.ParseInt(idstr, 10,64)
    if err!=nil{
        http.Error(w, "error, id not int", http.StatusBadRequest)
        return
    }


    quotes, err := db.GetClient(dbConn,id)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(quotes)
}
