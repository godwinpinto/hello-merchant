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

	"github.com/bwmarrin/snowflake"
	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/authn"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/embargo"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/ip_intel"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/user_intel"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/vault"
)

type ValidateResponseStruct struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Validate(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Println("failed to connect to the database", err)
		formatResponse(500, "DB connection error", w)
		return
	}

	cookie, err := request.Cookie("admin_token")
	if err != nil {
		formatResponse(500, "Invalid request", w)
		return
	}

	cookieValue := cookie.Value

	if cookieValue == "" {
		formatResponse(500, "Invalid request", w)
		return
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFn()

	token := os.Getenv("PANGEA_AUTHN_TOKEN")

	if token == "" {
		log.Println("Unauthorized: No token present")
		formatResponse(500, "DB connection error", w)
		return
	}

	// Create config and client
	authncli := authn.New(&pangea.Config{
		Token:  token,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	fmt.Println("cookieValue:::", cookieValue)
	input := authn.ClientSessionListRequest{
		Token: cookieValue,
	}

	resp, err := authncli.Client.Session.List(ctx, input)
	if err != nil {
		log.Println("Unauthorized: No session present")
		formatResponse(500, "Session error", w)
		return
	}

	responseStatus := ""
	responseStatus = *resp.Status
	if responseStatus == "InvalidToken" {
		formatResponse(500, "Invalid token. Please login again", w)
		return
	}

	input1 := authn.ClientSessionLogoutRequest{
		Token: cookieValue,
	}

	authncli.Client.Session.Logout(ctx, input1)

	var userBean beans.UserMasterStruct
	result := gormDB.Raw(`SELECT * FROM UPN_USER_MASTER WHERE USER_ID=?`, resp.Result.Sessions[0].Email).Scan(&userBean)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	//check if user is disabled
	if userBean.Active == "D" {
		formatResponse(300, "Your user id is disabled", w)
		return
	}

	fmt.Println(userBean)
	if userBean.UUMRowId == "" {
		fmt.Println("New User Registration")
		embargoStatus, err := checkEmbargo(request.RemoteAddr, ctx)
		if err != nil {
			log.Println("Error in token", err)
			formatResponse(300, "Application is not supported in your country", w)
			return
		}
		if embargoStatus {
			formatResponse(300, "Application is not supported in your country", w)
			return
		}

		ipIntelStatus, err := checkIPIntel(request.RemoteAddr, ctx)
		if err != nil {
			log.Print("Error in token", err)
			formatResponse(300, "Application is not supported for your IP address", w)
			return
		}
		if ipIntelStatus {
			formatResponse(300, "Application is not supported for your IP address", w)
			return
		}

		userIntelStatus, err := checkUserIntel(resp.Result.Sessions[0].Email, ctx)
		if err != nil {
			log.Print("Error in token", err)
			formatResponse(300, "Your User Id is detected as risk as per SpyCloud in last 30 days and hence not allowed to registered here", w)
			return
		}
		if userIntelStatus {
			formatResponse(300, "Your User Id is detected as risk as per SpyCloud in last 30 days and hence not allowed to registered here", w)
			return
		}

		id := node.Generate()

		rowId := id.String()

		insertResult := gormDB.Raw("INSERT INTO UPN_USER_MASTER (UUM_ROW_ID, USER_ID, `ROLE`, ACTIVE, CREATED_DT, CREATED_BY, UPDATED_DT, UPDATED_BY) VALUES(?, ?, 'N', 'Y', NOW(), 'SYSTEM', NOW(), 'SYSTEM')", rowId, resp.Result.Sessions[0].Email).Scan(&userBean)

		if insertResult.Error != nil {
			log.Print("Error in insert", err)
			formatResponse(300, "Unable to create registration", w)
			return
		}

		userBean.UUMRowId = rowId
		userBean.UserId = resp.Result.Sessions[0].Email

		fmt.Println("userBean after insert:", userBean.UUMRowId)
		fmt.Println("userBean after insert:", userBean.UserId)

	}

	jwtBean := beans.JwtBeanStruct{
		Id:  userBean.UUMRowId,
		Sub: userBean.UserId,
	}

	data, err := json.Marshal(jwtBean)
	if err != nil {
		log.Println("Error in token", err)
		formatResponse(300, "Unable to create session token", w)
		return
	}

	jwtToken, err := createJwt(string(data), ctx)
	if err != nil {
		log.Println("Error in token", err)
		formatResponse(300, "Unable to create session token 2", w)
		return
	}

	jwtCookie := &http.Cookie{
		Name:    "admin_jwt_token",
		Value:   jwtToken,
		Expires: time.Now().Add(time.Hour),
		Path:    "/",
		Secure:  true,
	}

	http.SetCookie(w, jwtCookie)
	tokenCookieRemove := http.Cookie{
		Name:    "admin_token", // Replace with the actual name of your cookie
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, -1), // Set expiration time to the past
	}

	// Set the cookie in the response
	http.SetCookie(w, &tokenCookieRemove)

	respStatus := ValidateResponseStruct{
		Status: 200,
	}
	data1, err := json.Marshal(respStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data1)

	/* 	http.SetCookie(w, cookie)
	   	http.Redirect(w, request, "/", http.StatusSeeOther)
	*/ /*
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	*/
}

func checkEmbargo(ipAddress string, ctx context.Context) (bool, error) {
	embargoToken := os.Getenv("VITE_PANGEA_EMBARGO_TOKEN")

	embargocli := embargo.New(&pangea.Config{
		Token:  embargoToken,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	input := &embargo.IPCheckInput{
		IP: pangea.String(ipAddress),
	}

	checkResponse, err := embargocli.IPCheck(ctx, input)
	if err != nil {
		return true, err
	}
	if len(checkResponse.Result.Sanctions) == 0 {
		return false, nil
	} else {
		return true, nil
	}

}

func checkIPIntel(ipAddress string, ctx context.Context) (bool, error) {
	ipintelToken := os.Getenv("VITE_PANGEA_IPINTEL_TOKEN")

	intelcli := ip_intel.New(&pangea.Config{
		Token:  ipintelToken,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	input := &ip_intel.IpReputationRequest{
		Ip:       ipAddress,
		Raw:      true,
		Verbose:  true,
		Provider: "crowdstrike",
	}

	resp, err := intelcli.Reputation(ctx, input)

	if err != nil {
		return true, err
	}
	if resp.Result.Data.Verdict == "unknown" {
		return false, nil
	} else {
		return true, nil
	}

}

func checkUserIntel(emailAddress string, ctx context.Context) (bool, error) {
	userintelToken := os.Getenv("VITE_PANGEA_USERINTEL_TOKEN")

	userintel := user_intel.New(&pangea.Config{
		Token:  userintelToken,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	input := &user_intel.UserBreachedRequest{
		Email:    emailAddress,
		Raw:      true,
		Verbose:  false,
		Start:    "30d",
		End:      "1d",
		Provider: "spycloud",
	}

	out, err := userintel.UserBreached(ctx, input)
	if err != nil {
		return true, err
	}
	if out.Result.Data.BreachCount == 0 {
		return false, nil
	} else {
		return true, nil

	}

}

func formatResponse(status int, message string, w http.ResponseWriter) {
	respStatus := ValidateResponseStruct{
		Status:  status,
		Message: message,
	}
	data, err := json.Marshal(respStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func createJwt(data string, ctx context.Context) (string, error) {

	vaultToken := os.Getenv("VITE_PANGEA_VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("Unauthorized: No token present")
	}

	vaultcli := vault.New(&pangea.Config{
		Token:  vaultToken,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	jwtId := os.Getenv("VITE_PANGEA_VAULT_JWT_ID")
	if vaultToken == "" {
		log.Fatal("Unauthorized: No token present")
	}

	input := &vault.JWTSignRequest{
		ID:      jwtId,
		Payload: data,
	}

	jr, err := vaultcli.JWTSign(ctx, input)

	if err != nil {
		return "", err
	}

	return jr.Result.JWS, nil

}
