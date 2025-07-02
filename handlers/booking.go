package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type ScheduleRequest struct {
	BookingID string `json:"booking_id"`
	Date      string `json:"date"` // ISO 8601
}

type ScheduleResponse struct {
	Message   string    `json:"message"`
	Scheduled time.Time `json:"scheduled"`
}

func ScheduleBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	scheduledDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	resp := ScheduleResponse{
		Message:   "Booking scheduled successfully",
		Scheduled: scheduledDate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
