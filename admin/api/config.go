package handler

import (
	"context"
	"encoding/json"
	"hello-merchant/beans"
	"hello-merchant/database"
	"log"
	"net/http"
	"time"
)

func Config(w http.ResponseWriter, request *http.Request) {

	cookie, err := request.Cookie("admin_jwt_token")
	if err != nil {
		formatResponse(500, "Invalid request", w)
		return
	}

	cookieValue := cookie.Value

	if cookieValue == "" {
		formatResponse(500, "Invalid request", w)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFn()

	jwtClaims, err := verifyJWT(cookieValue, ctx)

	if err != nil {
		formatResponse(500, "Invalid login", w)
		return
	}

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queryParams := request.URL.Query()
	pgType := queryParams.Get("type")

	configResponse := beans.ConfigResponseStruct{}
	if pgType == "XRPL" {
		var xrplStructs beans.XrplMappingStruct
		result := gormDB.Raw(`SELECT * FROM UPN_XRPL_MAPPING WHERE UUM_ROW_ID=?`, jwtClaims.Id).Scan(&xrplStructs)

		if result.Error != nil {
			formatResponse(300, "Invalid Count", w)
			return
		}

		if result.RowsAffected <= 0 {
			formatResponse(300, "No records", w)
			return
		}
		configResponse.AcNo = xrplStructs.XrplAcNo
	} else if pgType == "SQUAREUP" {
		var squareUpStructs beans.SquareUpMappingStruct
		result := gormDB.Raw(`SELECT * FROM UPN_SQUAREUP_MAPPING WHERE UUM_ROW_ID=?`, jwtClaims.Id).Scan(&squareUpStructs)

		if result.Error != nil {
			formatResponse(300, "Invalid Count", w)
			return
		}

		if result.RowsAffected <= 0 {
			formatResponse(300, "No records", w)
			return
		}
		configResponse.AcNo = squareUpStructs.UrlUuid
	}

	// Convert the array of structs to JSON
	data, err := json.Marshal(configResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
