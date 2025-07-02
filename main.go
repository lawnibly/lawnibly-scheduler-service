package main

import (
	"log"
	"net/http"

	"github.com/lawnibly/scheduler-service/handlers"
)

func main() {
	http.HandleFunc("/api/schedule", handlers.ScheduleBooking)

	port := ":8080"
	log.Println("Scheduler service running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
