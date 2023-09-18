package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
)

type TransactionOverviewStruct struct {
	Count  int    `json:"count" gorm:"column:COUNT_ACTION"`
	Action string `json:"action" gorm:"column:USER_ACTION"`
}

func TransactionOverview(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []TransactionOverviewStruct
	result := gormDB.Raw(`SELECT COUNT(arm.USER_ACTION) COUNT_ACTION, arm.USER_ACTION
	FROM AUDIT_REQUEST_MASTER arm, AUDIT_BLACKLIST_DOMAIN_MASTER abdm 
	WHERE arm.URL_DOMAIN =abdm.DOMAIN_NAME 
	AND arm.CREATED_DATETIME >=DATE_SUB(NOW(), INTERVAL 24 HOUR)
	GROUP BY arm.USER_ACTION
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
