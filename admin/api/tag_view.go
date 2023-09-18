package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
	"time"
)

type TagStruct struct {
	TagName   string    `json:"tag_name" gorm:"column:TAG_NAME"`
	TagType   string    `json:"tag_type" gorm:"column:TAG_TYPE"`
	CreatedDt time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
}

func TagView(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []TagStruct
	result := gormDB.Raw(`SELECT TAG_NAME, TAG_TYPE , CREATED_DT 
	FROM AUDIT_TAGS_MASTER
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
