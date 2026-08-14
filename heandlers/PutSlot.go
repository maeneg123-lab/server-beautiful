package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    "strconv"

    "studia-of-biautiful-api/db"
)

func PutSlot(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
    // 1. Получаем client_id
    clientIDStr := r.URL.Query().Get("client_id")
    if clientIDStr == "" {
        http.Error(w, "error: client_id is required", http.StatusBadRequest)
        return
    }
    clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
    if err != nil {
        http.Error(w, "error: client_id must be a number", http.StatusBadRequest)
        return
    }

    // 2. Получаем slot_id
    slotIDStr := r.URL.Query().Get("slot_id")
    if slotIDStr == "" {
        http.Error(w, "error: slot_id is required", http.StatusBadRequest)
        return
    }
    slotID, err := strconv.ParseInt(slotIDStr, 10, 64)
    if err != nil {
        http.Error(w, "error: slot_id must be a number", http.StatusBadRequest)
        return
    }

    // 3. Вызываем функцию из db
    err = db.PutSlot(dbConn, clientID, slotID)
    if err != nil {
        http.Error(w, "error db: "+err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "success! slot updated!")
}
