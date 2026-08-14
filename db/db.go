package db

import(
    "database/sql"
    "log"
    "os"
    "fmt"

    _"github.com/lib/pq"
    "studia-of-biautiful-api/models"
)

//функции клиента

func NewClient(db *sql.DB, username string, userID string) error {
    // Проверяем, существует ли клиент
    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM clients1 WHERE user_id = $1)", userID).Scan(&exists)
    if err != nil {
        return err
    }
    if exists {
        return nil // Клиент уже есть
    }

    // Вставляем нового клиента
    _, err = db.Exec("INSERT INTO clients1 (username, user_id) VALUES ($1, $2)", username, userID)
    return err
}

func AddRewiew(db *sql.DB, username string, text string, rating float64, operation string) error{
    _, err := db.Exec(`INSERT INTO reviews1 (username, text, operation, rating) VALUES ($1, $2, $3, $4)`, username, text, operation, rating)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    return err
}

func GetRewiew_User(db *sql.DB, username string) ([]models.RatingResponse, error){
    rows, err:=db.Query(`SELECT id, username, text, created_at, rating, operation FROM reviews1 WHERE username=$1`, username)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    defer rows.Close()

    var ratings []models.RatingResponse // ← ОСНОВНОЙ СЛАЙС
    for rows.Next() {
        var r models.RatingResponse // ← ОДНА ЗАПИСЬ
        err := rows.Scan(&r.ID, &r.Username,&r.Text, &r.CreatedAt, &r.Rating, &r.Operation) // ← ЧИТАЕМ В ПОЛЯ q
        if err != nil {
            log.Println("Ошибка сканирования:", err)
            continue
        }
        ratings = append(ratings, r) // ← ДОБАВЛЯЕМ ЗАПИСЬ В СЛАЙС
    }
    return ratings, nil
}

func PutSlot(db *sql.DB, userID int64, slotID int64) error {
    // Находим клиента по user_id
    var clientID int
    err := db.QueryRow("SELECT id FROM clients1 WHERE user_id = $1", userID).Scan(&clientID)
    if err != nil {
        return fmt.Errorf("клиент не найден: %v", err)
    }

    _, err = db.Exec("UPDATE slots1 SET client=$1, status='booked' WHERE id=$2", clientID, slotID)
    return err
}

func GetFreeSlots(db *sql.DB) ([]models.Slot, error){
    rows, err:=db.Query(`SELECT id, start_time, end_time, client, status FROM slots1 WHERE status='free'`)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    defer rows.Close()

    var slots []models.Slot // ← ОСНОВНОЙ СЛАЙС
    for rows.Next() {
        var s models.Slot // ← ОДНА ЗАПИСЬ
        err := rows.Scan(&s.ID, &s.StartTime, &s.EndTime, &s.ClientID, &s.Status) // ← ЧИТАЕМ В ПОЛЯ q
        if err != nil {
            log.Println("Ошибка сканирования:", err)
            continue
        }
        slots = append(slots, s) // ← ДОБАВЛЯЕМ ЗАПИСЬ В СЛАЙС
    }
    return slots, nil
}

//функции администратора

func GetRewiew(db *sql.DB, id int64) ([]models.RatingResponse, error){
    rows, err:=db.Query(`SELECT id, username, text, created_at, operation, rating FROM reviews1 WHERE id=$1`, id)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    defer rows.Close()

    var ratings []models.RatingResponse // ← ОСНОВНОЙ СЛАЙС
    for rows.Next() {
        var r models.RatingResponse // ← ОДНА ЗАПИСЬ
        err := rows.Scan(&r.ID, &r.Username, &r.Text, &r.CreatedAt, &r.Operation, &r.Rating) // ← ЧИТАЕМ В ПОЛЯ q
        if err != nil {
            log.Println("Ошибка сканирования:", err)
            continue
        }
        ratings = append(ratings, r) // ← ДОБАВЛЯЕМ ЗАПИСЬ В СЛАЙС
    }
    return ratings, nil
}

func AddSlot(db *sql.DB, start_time string, end_time string, status string) error{
    _, err := db.Exec(`INSERT INTO slots1 (start_time, end_time, status) VALUES ($1, $2, $3)`, start_time, end_time, status)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    return err
}

func GetClients(db *sql.DB) ([]models.ClientResponse, error){
    rows, err:=db.Query(`SELECT id, user_id, username, created_at FROM clients1`)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    defer rows.Close()

    var clients []models.ClientResponse // ← ОСНОВНОЙ СЛАЙС
    for rows.Next() {
        var c models.ClientResponse // ← ОДНА ЗАПИСЬ
        err := rows.Scan(&c.ID, &c.User_id, &c.Username, &c.CreatedAt) // ← ЧИТАЕМ В ПОЛЯ q
        if err != nil {
            log.Println("Ошибка сканирования:", err)
            continue
        }
        clients = append(clients, c) // ← ДОБАВЛЯЕМ ЗАПИСЬ В СЛАЙС
    }
    return clients, nil
}

func GetClient(db *sql.DB, id int64) ([]models.ClientResponse, error){
    rows, err:=db.Query(`SELECT id, user_id, username, created_at FROM clients1 WHERE id=$1`, id)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    defer rows.Close()

    var clients []models.ClientResponse // ← ОСНОВНОЙ СЛАЙС
    for rows.Next() {
        var c models.ClientResponse // ← ОДНА ЗАПИСЬ
        err := rows.Scan(&c.ID, &c.User_id, &c.Username, &c.CreatedAt) // ← ЧИТАЕМ В ПОЛЯ q
        if err != nil {
            log.Println("Ошибка сканирования:", err)
            continue
        }
        clients = append(clients, c) // ← ДОБАВЛЯЕМ ЗАПИСЬ В СЛАЙС
    }
    return clients, nil
}

func GetRewiews(db *sql.DB) ([]models.RatingResponse, error){
    rows, err:=db.Query(`SELECT id, username, text, created_at, operation, rating FROM reviews1`)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    defer rows.Close()

    var ratings []models.RatingResponse // ← ОСНОВНОЙ СЛАЙС
    for rows.Next() {
        var r models.RatingResponse // ← ОДНА ЗАПИСЬ
        err := rows.Scan(&r.ID, &r.Username, &r.Text, &r.CreatedAt, &r.Operation, &r.Rating) // ← ЧИТАЕМ В ПОЛЯ q
        if err != nil {
            log.Println("Ошибка сканирования:", err)
            continue
        }
        ratings = append(ratings, r) // ← ДОБАВЛЯЕМ ЗАПИСЬ В СЛАЙС
    }
    return ratings, nil
}

func GetSlots(db *sql.DB) ([]models.Slot, error){
    rows, err:=db.Query(`SELECT id, start_time, end_time, client, status FROM slots1`)

    if err!=nil{
        log.Fatal("error: ", err)
    }

    defer rows.Close()

    var slots []models.Slot // ← ОСНОВНОЙ СЛАЙС
    for rows.Next() {
        var s models.Slot // ← ОДНА ЗАПИСЬ
        err := rows.Scan(&s.ID, &s.StartTime, &s.EndTime, &s.ClientID, &s.Status) // ← ЧИТАЕМ В ПОЛЯ q
        if err != nil {
            log.Println("Ошибка сканирования:", err)
            continue
        }
        slots = append(slots, s) // ← ДОБАВЛЯЕМ ЗАПИСЬ В СЛАЙС
    }
    return slots, nil
}

func InitDB() *sql.DB{
    connStr := os.Getenv("DATABASE_URL")
    if connStr == ""{
        connStr = "user=postgres password=36863686 dbname=studia_of_beautiful_api sslmode=disable"
    }
    db,err := sql.Open("postgres", connStr)
    if err != nil{
        log.Fatal("error db open: ", err)
    }
    CreateTableClientSQL := `
    CREATE TABLE IF NOT EXISTS clients1 (
        id SERIAL PRIMARY KEY,
        user_id TEXT NOT NULL,
        username TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW()
    );`

    _,err = db.Exec(CreateTableClientSQL)
    if err != nil{
        log.Fatal("error create table client: ", err)
    }

    CreateTableRatingSQL :=`
    CREATE TABLE IF NOT EXISTS reviews1 (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL,
        text TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW(),
        operation TEXT DEFAULT NULL,
        rating FLOAT NOT NULL
    );`

    _,err = db.Exec(CreateTableRatingSQL)
    if err != nil{
        log.Fatal("error create table rating: ", err)
    }

    createTableSlotsSQL := `
    CREATE TABLE IF NOT EXISTS slots1 (
        id SERIAL PRIMARY KEY,
        start_time TIMESTAMP NOT NULL,
        end_time TIMESTAMP NOT NULL,
        status TEXT DEFAULT 'free',
        client INT REFERENCES clients1(id) ON DELETE SET NULL,
        created_at TIMESTAMP DEFAULT NOW()
    );`
    _, err = db.Exec(createTableSlotsSQL)
    if err != nil {
        log.Fatal("error create table slots: ", err)
    }

    return db
}
