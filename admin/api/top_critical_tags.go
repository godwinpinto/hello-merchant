package handler

import (
	"encoding/json"
	"hello-merchant/database"
	"log"
	"net/http"
)

type TopCriticalTagsStruct struct {
	Count   int    `json:"count" gorm:"column:COUNT_TAGS"`
	TagName string `json:"tag" gorm:"column:TAG_NAME"`
}

func TopCriticalTags(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []TopCriticalTagsStruct
	result := gormDB.Raw(`SELECT COUNT(aust.TAG_NAME) COUNT_TAGS, aust.TAG_NAME  
	FROM AUDIT_USER_SEARCH_TAGS aust, AUDIT_TAGS_MASTER atm 
	WHERE aust.CREATED_DT >=DATE_SUB(NOW(), INTERVAL 24 HOUR)
	AND aust.TAG_NAME =atm.TAG_NAME AND atm.TAG_TYPE ='C'
	GROUP BY TAG_NAME 
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
