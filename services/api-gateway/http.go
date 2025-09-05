package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	// "time"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(time.Second * 9)
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to pass json data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// validation
	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}
	log.Println("Trip preview success")
	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	log.Printf("Response from trip-service: %v, %v", resp, err)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	var respBody any

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w, "failed to pass json data from trip-service", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: respBody}
	writeJSON(w, http.StatusOK, response)
}
