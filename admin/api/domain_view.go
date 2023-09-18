package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
	"time"
)

type DomainStruct struct {
	RowId     string    `json:"id" gorm:"column:ROW_ID"`
	Domain    string    `json:"domain_name" gorm:"column:DOMAIN_NAME"`
	CreatedDt time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
}

func DomainView(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []DomainStruct
	result := gormDB.Raw(`SELECT ROW_ID , DOMAIN_NAME , CREATED_DT 
	FROM AUDIT_BLACKLIST_DOMAIN_MASTER
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
