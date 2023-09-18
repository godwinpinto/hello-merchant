package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
)

type TopUserActivityStruct struct {
	Count  int    `json:"count" gorm:"column:COUNT_ACTIVITY"`
	UserId string `json:"user_id" gorm:"column:USER_ID"`
}

func TopUserActivity(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []TopUserActivityStruct
	result := gormDB.Raw(`SELECT COUNT(USER_ID) COUNT_ACTIVITY, USER_ID
	FROM AUDIT_REQUEST_MASTER
	WHERE CREATED_DATETIME >=DATE_SUB(NOW(), INTERVAL 24 HOUR)
	GROUP BY USER_ID  
	ORDER BY 1 DESC 
	LIMIT 10
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
