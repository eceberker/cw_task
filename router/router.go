package router

import (
	"github.com/eceberker/cw_task/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/logs", middleware.PostLog).Methods("POST", "OPTIONS")
	router.HandleFunc("/logs/durations/average", middleware.GetDailyAverageDurations).Methods("GET", "OPTIONS")
	router.HandleFunc("/logs/daily/active", middleware.GetDailyActiveUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/logs/users", middleware.GetTotalUsers).Methods("GET", "OPTIONS")

	return router
}
