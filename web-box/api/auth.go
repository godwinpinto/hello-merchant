package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"hello-merchant-web-box/beans"
	"hello-merchant-web-box/database"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/authn"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/vault"
)

func Auth(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queryParams := request.URL.Query()
	code := queryParams.Get("code")
	signup := queryParams.Get("signup")
	state := queryParams.Get("state")

	ctx, cancelFn := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFn()

	token := os.Getenv("PANGEA_AUTHN_TOKEN")

	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	jwtId := os.Getenv("VITE_PANGEA_JWT_ID")

	if jwtId == "" {
		log.Fatal("Unauthorized: No JWT ID present")
	}

	// Create config and client
	client := authn.New(&pangea.Config{
		Token:  token,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	input := authn.ClientUserinfoRequest{
		Code: code,
	}
	fmt.Println(signup)
	fmt.Println(state)
	resp, err := client.Client.Userinfo(ctx, input)
	if err != nil {
		log.Fatal("resp error in token check", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(resp)
	fmt.Println(resp.Result.ActiveToken.Token)

	/* 	data, err := json.Marshal(resp.Result)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */
	var userBean beans.UserMasterStruct
	result := gormDB.Raw(`SELECT * FROM UPN_USER_MASTER WHERE USER_ID=?`, resp.Result.ActiveToken.Email).Scan(&userBean)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(userBean)

	if result.RowsAffected <= 0 {
		http.Error(w, "No Records", http.StatusNotFound)
		return
	}

	jwtBean := beans.JwtBeanStruct{
		UserId:   userBean.UserId,
		UUMRowId: userBean.UUMRowId,
	}

	// Convert the array of structs to JSON
	jwtJsonString, err := json.Marshal(jwtBean)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonString := string(jwtJsonString)

	inputJWTSign := &vault.JWTSignRequest{
		ID:      jwtId,
		Payload: jsonString,
	}
	vaultToken := os.Getenv("VITE_PANGEA_CUBE_VAULT_TOKEN")

	if vaultToken == "" {
		log.Fatal("Unauthorized: No token present")
	}

	vaultcli := vault.New(&pangea.Config{
		Token:  vaultToken,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	jr, err := vaultcli.JWTSign(ctx, inputJWTSign)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	layout := "2006-01-02T15:04:05.999999Z"

	// Parse the string into a time.Time object
	expiryTime, err := time.Parse(layout, resp.Result.ActiveToken.Expire)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cookie := &http.Cookie{
		Name:    "email",
		Value:   resp.Result.ActiveToken.Email,
		Expires: expiryTime,
		Path:    "/",
		Secure:  true,
	}

	http.SetCookie(w, cookie)

	cookie = &http.Cookie{
		Name:    "jwt_token",
		Value:   jr.Result.JWS,
		Path:    "/",
		Expires: expiryTime,
		Secure:  true,
	}

	http.SetCookie(w, cookie)

	cookie = &http.Cookie{
		Name:    "uum_row_id",
		Value:   userBean.UUMRowId,
		Path:    "/",
		Expires: expiryTime,
		Secure:  true,
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, request, "/", http.StatusSeeOther)
	/*
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	*/
}
