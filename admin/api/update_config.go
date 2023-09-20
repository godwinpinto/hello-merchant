package handler

import (
	"context"
	"fmt"
	"hello-merchant/beans"
	"hello-merchant/database"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/pangeacyber/pangea-go/pangea-sdk/v2/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/v2/service/audit"
)

func UpdateConfig(w http.ResponseWriter, request *http.Request) {

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
	acNo := queryParams.Get("ac_no")

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	//	configResponse := beans.ConfigResponseStruct{}
	if pgType == "XRPL" {
		var xrplStructs beans.XrplMappingStruct
		result := gormDB.Raw(`SELECT * FROM UPN_XRPL_MAPPING WHERE UUM_ROW_ID=?`, jwtClaims.Id).Scan(&xrplStructs)
		if result.Error != nil {
			formatResponse(300, "Invalid Count", w)
			return
		}

		if result.RowsAffected == 0 && acNo == "" {
			formatResponse(300, "Invalid Request", w)
			return
		}
		if result.RowsAffected == 0 && acNo != "" {
			createSecureAuditEvent("I", jwtClaims.Id, acNo, ctx)
			result := gormDB.Raw(`INSERT INTO UPN_DB.UPN_XRPL_MAPPING
			(UXM_ROW_ID, UUM_ROW_ID, XRPL_AC_NO, ACTIVE, CREATED_DT, CREATED_BY, UPDATED_DT, UPDATED_BY)
			VALUES(?, ?, ?, 'Y', NOW(), 'SYSTEM', NOW(), 'SYSTEM')
			`, node.Generate().String(), jwtClaims.Id, acNo).Scan(&xrplStructs)
			if result.Error != nil {
				formatResponse(300, "Cannot insert", w)
				return
			}
		} else if acNo != "" {
			createSecureAuditEvent("UDATED", jwtClaims.Id, acNo, ctx)
			result := gormDB.Raw(`UPDATE UPN_XRPL_MAPPING SET XRPL_AC_NO=? WHERE UUM_ROW_ID=?`, acNo, jwtClaims.Id).Scan(&xrplStructs)
			if result.Error != nil {
				formatResponse(300, "Cannot update", w)
				return
			}
		} else if acNo == "" {
			createSecureAuditEvent("REMOVED", jwtClaims.Id, acNo, ctx)
			result := gormDB.Raw(`DELETE FROM UPN_XRPL_MAPPING WHERE UUM_ROW_ID=?`, jwtClaims.Id).Scan(&xrplStructs)
			if result.Error != nil {
				formatResponse(300, "Cannot update", w)
				return
			}
		}

		formatResponse(200, "Record udpated", w)
		return

	}

	/*
		 	// Convert the array of structs to JSON
			data, err := json.Marshal(configResponse)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
	*/
}

func createSecureAuditEvent(eventType string, userId string, acNo string, ctx context.Context) {
	secureAuditToken := os.Getenv("VITE_PANGEA_SECURE_AUDIT_LOG_TOKEN")

	auditcli, err := audit.New(&pangea.Config{
		Token:  secureAuditToken,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})
	if err != nil {
		log.Fatal("failed to create audit client")
	}

	event := &audit.StandardEvent{
		Message: "Hello, World!",
	}
	_, err = auditcli.Log(ctx, event, true)

	if err != nil {
		log.Println(err)
	}
}
