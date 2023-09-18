package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
)

type DbStatsStruct struct {
	TotalUsers   int `json:"total_users" gorm:"column:TOTAL_USERS"`
	TotalPaste   int `json:"total_paste" gorm:"column:TOTAL_PASTE"`
	TotalCutCopy int `json:"total_cut_copy" gorm:"column:TOTAL_CUT_COPY"`
	TotalVisit   int `json:"total_visit" gorm:"column:TOTAL_VISITS"`
}

func Stats(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []DbStatsStruct
	result := gormDB.Raw(`select 
	(select count(distinct user_id) from AUDIT_REQUEST_MASTER arm 
	WHERE 
		arm.CREATED_DATETIME >= CURDATE() 
		AND arm.CREATED_DATETIME < CURDATE() + INTERVAL 1 DAY
		AND TIME(arm.CREATED_DATETIME) >= '00:00:00' 
		AND TIME(arm.CREATED_DATETIME) <= '23:59:59') as TOTAL_USERS
	,(select count(USER_ACTION) from AUDIT_REQUEST_MASTER arm 
	WHERE 
		arm.CREATED_DATETIME >= CURDATE() 
		AND arm.CREATED_DATETIME < CURDATE() + INTERVAL 1 DAY
		AND TIME(arm.CREATED_DATETIME) >= '00:00:00' 
		AND TIME(arm.CREATED_DATETIME) <= '23:59:59'
		AND USER_ACTION ='I') as TOTAL_VISITS,
	(select count(USER_ACTION)  from AUDIT_REQUEST_MASTER arm 
	WHERE 
		arm.CREATED_DATETIME >= CURDATE() 
		AND arm.CREATED_DATETIME < CURDATE() + INTERVAL 1 DAY
		AND TIME(arm.CREATED_DATETIME) >= '00:00:00' 
		AND TIME(arm.CREATED_DATETIME) <= '23:59:59'
		AND USER_ACTION ='P') as TOTAL_PASTE
	,(select count(USER_ACTION) from AUDIT_REQUEST_MASTER arm 
	WHERE 
		arm.CREATED_DATETIME >= CURDATE() 
		AND arm.CREATED_DATETIME < CURDATE() + INTERVAL 1 DAY
		AND TIME(arm.CREATED_DATETIME) >= '00:00:00' 
		AND TIME(arm.CREATED_DATETIME) <= '23:59:59'
		AND USER_ACTION IN('C','X')) as TOTAL_CUT_COPY
	
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
