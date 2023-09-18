package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
	"time"
)

type LiveCntrlEventsStruct struct {
	CountAction int       `json:"count" gorm:"column:COUNT_ACTION"`
	UserAction  string    `json:"action" gorm:"column:USER_ACTION"`
	ServerTime  time.Time `json:"server_time" gorm:"type:timestamp;column:CURRENT_SERVER_TIME"`
	// Add more fields as needed from your database query result
}

func LiveCntrlEvents(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []LiveCntrlEventsStruct
	result := gormDB.Raw(`SELECT USER_ACTION, count(USER_ACTION) as COUNT_ACTION, NOW() AS CURRENT_SERVER_TIME 
	FROM AUDIT_REQUEST_MASTER
	WHERE CREATED_DATETIME >=DATE_SUB(NOW(), INTERVAL 60 SECOND)
	group by USER_ACTION
	`).Scan(&results)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected <= 0 {
		http.Error(w, "No Records", http.StatusNotFound)
		return
	}

	// Convert the array of structs to JSON
	data, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
