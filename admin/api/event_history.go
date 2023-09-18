package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
	"time"
)

type EventHistoryStruct struct {
	RowId      string    `json:"id" gorm:"column:ARM_ROW_ID"`
	UserId     string    `json:"user_id" gorm:"column:USER_ID"`
	UserAction string    `json:"action" gorm:"column:USER_ACTION"`
	CreatedDt  time.Time `json:"created_dt" gorm:"column:CREATED_DATETIME"`
	Domain     string    `json:"domain" gorm:"column:URL_DOMAIN"`
	Content    string    `json:"content" gorm:"column:CONTENT_SHORT"`
}

func EventHistory(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []EventHistoryStruct
	result := gormDB.Raw(`SELECT ARM_ROW_ID , USER_ID,USER_ACTION , CREATED_DATETIME , URL_DOMAIN , CONTENT_SHORT 
	FROM AUDIT_REQUEST_MASTER arm 
	WHERE arm.CREATED_DATETIME >=DATE_SUB(NOW(), INTERVAL 1000 HOUR)
	ORDER BY CREATED_DATETIME DESC 
	LIMIT 200
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
