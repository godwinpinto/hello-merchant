package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"hello-merchant/beans"
	"hello-merchant/database"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/vault"
)

func Dashboard(w http.ResponseWriter, request *http.Request) {

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

	var summaryCount beans.DashboardSummaryStruct
	result := gormDB.Raw(`SELECT 
	(SELECT COUNT(UUM_ROW_ID) from UPN_TRANSACTION_MASTER utm WHERE utm.UUM_ROW_ID=? AND CHANNEL='Ripple') AS RIPPLE_COUNT,
	(SELECT COUNT(UUM_ROW_ID) from UPN_TRANSACTION_MASTER utm WHERE utm.UUM_ROW_ID=? AND CHANNEL='SquareUp') AS SQUAREUP_COUNT
	`, jwtClaims.Id, jwtClaims.Id).Scan(&summaryCount)

	if result.Error != nil {
		formatResponse(300, "Invalid Count", w)
		return
	}

	if result.RowsAffected <= 0 {
		formatResponse(300, "No records", w)
		return
	}

	var reportData []beans.DashboardReportStruct
	results := gormDB.Raw(`SELECT 
    DATE_FORMAT(DATE_SUB(CURRENT_DATE(), INTERVAL t.n DAY), '%d-%b') as DATE_RANGE,
    COALESCE(COUNT(UTM_ROW_ID), 0) as RECORD_COUNT,    channels.CHANNEL
	FROM 
    (SELECT 0 as n UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t
	CROSS JOIN 
    (SELECT DISTINCT CHANNEL FROM UPN_TRANSACTION_MASTER) channels
	LEFT JOIN 
    UPN_TRANSACTION_MASTER
	ON 
    DATE(UPN_TRANSACTION_MASTER.CREATED_DT) = DATE_SUB(CURRENT_DATE(), INTERVAL t.n DAY)
    AND channels.CHANNEL = UPN_TRANSACTION_MASTER.CHANNEL AND UPN_TRANSACTION_MASTER.UUM_ROW_ID =?
	GROUP BY 
    channels.CHANNEL, date_range
	ORDER BY 
    channels.CHANNEL, DATE_RANGE
	`, jwtClaims.Id).Scan(&reportData)

	if results.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if results.RowsAffected <= 0 {
		http.Error(w, "No Records", http.StatusNotFound)
		return
	}

	response := beans.DashboardResponseStruct{
		RippleCount:      summaryCount.RippleCount,
		SquareUpCount:    summaryCount.SquareUpCount,
		LastTransactions: reportData,
	}

	// Convert the array of structs to JSON
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func verifyJWT(jwtString string, ctx context.Context) (beans.JwtBeanStruct, error) {
	vaultToken := os.Getenv("VITE_PANGEA_VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("Unauthorized: No token present")
	}

	vaultcli := vault.New(&pangea.Config{
		Token:  vaultToken,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	input := &vault.JWTVerifyRequest{
		JWS: jwtString,
	}

	jwtBean := beans.JwtBeanStruct{}

	jr, err := vaultcli.JWTVerify(ctx, input)
	if err != nil {
		log.Println("failed to verify JWT", err)
		return jwtBean, err
	}

	if jr.Result.ValidSignature {
		token, _, err := new(jwt.Parser).ParseUnverified(jwtString, jwt.MapClaims{})
		if err != nil {
			fmt.Printf("Error %s", err)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// obtains claims
			sub := fmt.Sprint(claims["sub"])
			id := fmt.Sprint(claims["id"])
			jwtBean.Id = id
			jwtBean.Sub = sub

		}
	}

	if jr.Result.ValidSignature {
		return jwtBean, nil
	} else {
		return jwtBean, nil
	}

}
