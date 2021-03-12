package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	clients "github.com/eceberker/cw_task/gcp_clients"
	services "github.com/eceberker/cw_task/services"
)

// PostLog publishes the log in request body in GCP Pub/Sub service
func PostLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read the request body. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message := string(b)

	client, err := clients.NewPublisher()
	if err != nil {
		log.Fatalf("error:  %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := client.Publish(message)
	if response.Status == 500 {
		log.Fatal(response.MessageText)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

// GetTotalUsers gets total unique users ID and their last online date from GCP BigQuery
func GetTotalUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	response := services.GetTotalUsers()
	if response.Status == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

// GetDailyAverageDurations gets average session durations of daily active users according to day from GCP BigQuery
func GetDailyAverageDurations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	response := services.GetDailyAverageDurations()
	if response.Status == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

// GetDailyActiveUsers gets count of daily active unique users according to day from GCP BigQuery
func GetDailyActiveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	response := services.GetUsersPerDay()

	if response.Status == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
