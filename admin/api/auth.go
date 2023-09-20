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

func Auth(w http.ResponseWriter, request *http.Request) {

	queryParams := request.URL.Query()
	code := queryParams.Get("code")

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
	resp, err := client.Client.Userinfo(ctx, input)
	if err != nil {
		log.Fatal("resp error in token check", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//	fmt.Println(resp)
	fmt.Println(resp.Result)

	layout := "2006-01-02T15:04:05.999999Z"

	// Parse the string into a time.Time object
	expiryTime, err := time.Parse(layout, resp.Result.ActiveToken.Expire)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cookie := &http.Cookie{
		Name:    "admin_email",
		Value:   resp.Result.ActiveToken.Email,
		Expires: expiryTime,
		Path:    "/",
		Secure:  true,
	}

	http.SetCookie(w, cookie)

	cookie = &http.Cookie{
		Name:    "admin_token",
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
