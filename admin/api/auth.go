package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/authn"
)

type AuthStruct struct {
	RowId     string    `json:"id" gorm:"column:ROW_ID"`
	Domain    string    `json:"domain_name" gorm:"column:DOMAIN_NAME"`
	CreatedDt time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
}

func Auth(w http.ResponseWriter, request *http.Request) {

	/* 	gormDB, err := database.InitializeDB()
	   	if err != nil {
	   		log.Fatal("failed to connect to the database", err)
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	*/
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
	/* var results []DomainStruct
	result := gormDB.Raw(`SELECT ROW_ID , DOMAIN_NAME , CREATED_DT
	FROM AUDIT_BLACKLIST_DOMAIN_MASTER
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
	} */

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
		Name:    "token",
		Value:   resp.Result.ActiveToken.Token,
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
