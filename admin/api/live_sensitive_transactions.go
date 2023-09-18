package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
	"time"
)

type LiveSensitiveTransactionsStruct struct {
	UniqueID        string    `json:"uuid" gorm:"column:ARM_ROW_ID"`
	UserID          string    `json:"user_id" gorm:"column:USER_ID"`
	UserAction      string    `json:"action" gorm:"column:USER_ACTION"`
	Domain          string    `json:"domain" gorm:"column:URL_DOMAIN"`
	CreatedDateTime time.Time `json:"created_dt" gorm:"type:timestamp;column:CREATED_DATETIME"`
	// Add more fields as needed from your database query result
}

func LiveSensitiveTransactions(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []LiveSensitiveTransactionsStruct
	result := gormDB.Raw(`SELECT arm.ARM_ROW_ID, arm.USER_ID, arm.USER_ACTION, arm.CREATED_DATETIME, arm.URL_DOMAIN 
	FROM AUDIT_REQUEST_MASTER arm, AUDIT_BLACKLIST_DOMAIN_MASTER abdm 
	WHERE arm.URL_DOMAIN =abdm.DOMAIN_NAME 
	AND arm.CREATED_DATETIME >=DATE_SUB(NOW(), INTERVAL 10 SECOND)
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
