package main

import (
	"net/http"
	"time"
)

func (s *Server_t) EventForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("user_id")

	result := DayEvents(id, date)
	ResponseResult(w, "Events", result, http.StatusOK)

}

func (s *Server_t) EventForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("user_id")
	result := WeekEvents(id, date)
	ResponseResult(w, "Events", result, http.StatusOK)
}

func (s *Server_t) EventForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("user_id")

	result := MonthEvents(id, date)
	ResponseResult(w, "Events", result, http.StatusOK)
}
