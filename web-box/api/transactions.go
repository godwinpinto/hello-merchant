package handler

import (
	"encoding/json"
	"hello-merchant-web-box/beans"
	"hello-merchant-web-box/database"
	"log"
	"net/http"
)

func Transactions(w http.ResponseWriter, request *http.Request) {

	gormDB, err := database.InitializeDB()
	if err != nil {
		log.Fatal("failed to connect to the database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie, err := request.Cookie("uum_row_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	cookieValue := cookie.Value

	/* 	fmt.Fprintf(w, "Cookie Value: %s", cookieValue)

	   	if cookieValue == "" {
	   		log.Fatal("resp error in token check", err)
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	token, err := jwt.Parse(cookieValue, func(token *jwt.Token) (interface{}, error) {
	   		return nil, nil // Return nil for both interface and error to skip validation
	   	})

	   	if err != nil {
	   		fmt.Printf("Error decoding with key: %v\n", err)
	   		return
	   	}

	   	var customClaims beans.JwtBeanStruct
	   	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

	   		customClaims.UUMRowId = claims["uum_row_id"].(string)
	   		customClaims.UserId = claims["user_id"].(string)

	   		fmt.Println("Username:", customClaims.UUMRowId)
	   		fmt.Println("Role:", customClaims.UserId)
	   	} */

	/* 	for _, key := range keys {
		tryDecode([]byte(key))
	} */
	/*
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
		   	}*/
	var transactions []beans.TransactionMasterStruct
	result := gormDB.Raw(`SELECT * FROM UPN_TRANSACTION_MASTER WHERE UUM_ROW_ID=? ORDER BY CREATED_DT DESC LIMIT 3`, cookieValue).Scan(&transactions)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	/* 	if result.RowsAffected <= 0 {

	   		return
	   	}
	*/
	response := beans.ResponseStruct{
		Status: "200",
		Data:   transactions,
	}
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
