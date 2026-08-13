package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"      // ✅ Правильно
    "os"

    "studia-of-biautiful-api/db"
    "studia-of-biautiful-api/handlers"

    "github.com/joho/godotenv"
)

var dbConn *sql.DB

func main(){
    err:= godotenv.Load()
    if err!=nil{
        log.Println("file .env not found")
    }

    dbConn := db.InitDB()
    defer dbConn.Close()

    http.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request){
        handlers.GetClients(w, r, dbConn)
    })//получение всех клиентов✅

    http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request){
        handlers.GetClient(w, r, dbConn)
    })//получение одного клиента✅

    http.HandleFunc("/new_client", func(w http.ResponseWriter, r *http.Request){
        handlers.NewClient(w, r, dbConn)
    })//добавление клиента✅

    http.HandleFunc("/rewiews", func(w http.ResponseWriter, r *http.Request){
        handlers.GetRewiews(w, r, dbConn)
    })//получение всех отзывов✅

    http.HandleFunc("/rewiew", func(w http.ResponseWriter, r *http.Request){
        handlers.GetRewiew(w, r, dbConn)
    })//получение одного отзыва✅

    http.HandleFunc("/new_rewiew", func(w http.ResponseWriter, r *http.Request){
        handlers.AddRewiew(w, r, dbConn)
    })//добавление отзыва✅

    http.HandleFunc("/slot", func(w http.ResponseWriter, r *http.Request){
        handlers.GetSlots(w, r, dbConn)
    })//получение всех записей✅

    http.HandleFunc("/new_slot", func(w http.ResponseWriter, r *http.Request){
        handlers.AddSlot(w, r, dbConn)
    })//добавление слота для записи✅

    http.HandleFunc("/free_slot", func(w http.ResponseWriter, r *http.Request){
        handlers.GetFreeSlots(w, r, dbConn)
    })//получение свободных слотов для записи✅

    http.HandleFunc("/user_get_rewiew", func(w http.ResponseWriter, r *http.Request){
        handlers.GetRewiew_User(w, r, dbConn)
    })//пользователь получает свой отзыв✅

    http.HandleFunc("/slot_boocked", func(w http.ResponseWriter, r *http.Request){
        handlers.PutSlot(w, r, dbConn)
    })//занимание слота для записи пользователем✅

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Printf("Сервер запущен на порту %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}