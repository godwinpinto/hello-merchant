package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
	"time"
)

type UserStruct struct {
	Username     string    `json:"username" gorm:"column:USER_NAME"`
	AccessSecret string    `json:"access_secret" gorm:"column:ACCESS_SECRET"`
	CreatedDt    time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
}

func UserView(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []UserStruct
	result := gormDB.Raw(`SELECT USER_NAME , CREATED_DT 
	FROM AUDIT_USER_MASTER
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
