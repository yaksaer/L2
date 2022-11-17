package main

import (
	"net/http"
)

func (s *Server_t) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	event, err := RequestToJson(r)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = event.Validate()
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := AddEvent(*event); err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseResult(w, "Event add", event.inf, http.StatusCreated)
}

func (s *Server_t) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	event, err := RequestToJson(r)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = event.Validate()
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = UpdateInfo(*event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseResult(w, "Data update", event.inf, http.StatusOK)
	return
}

func (s *Server_t) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	event, err := RequestToJson(r)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = DeleteEvent(*event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseResult(w, "Data delete", event.inf, http.StatusOK)
	return

}
