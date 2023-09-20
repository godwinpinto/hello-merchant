package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/redact"
)

type RequestBody struct {
	Body string `json:"body"`
}

func Redact(w http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Create a struct to unmarshal the JSON into
	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Access the JSON variables
	fmt.Println("Key 1:", requestBody.Body)

	var data map[string]interface{}

	err = json.Unmarshal([]byte(requestBody.Body), &data)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	input := &redact.StructuredInput{
		Data:  data,
		Rules: []string{"Account"},
	}

	ctx := context.Background()

	token := os.Getenv("VITE_PANGEA_REDACT_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	redactcli := redact.New(&pangea.Config{
		Token:  token,
		Domain: os.Getenv("PANGEA_DOMAIN"),
	})

	redactResponse, err := redactcli.RedactStructured(ctx, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the array of structs to JSON
	redactedDataJSON, err := json.Marshal(redactResponse.Result.RedactedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(redactedDataJSON)
}
