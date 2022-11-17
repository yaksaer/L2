package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ResponseResult_t struct {
	Message string `json:"message"`
	Events  []Info `json:"events"`
}

func ResponseResult(w http.ResponseWriter, mess string, e []Info, status int) {
	var respResult = ResponseResult_t{
		Message: mess,
		Events:  e}

	json, err := json.Marshal(respResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func RequestToJson(r *http.Request) (*Event, error) {
	var info Info
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		return nil, errors.New("Wrong json")
	}
	event.UserId = info.UserId
	event.inf = append(event.inf, info)
	return &event, nil

}

type Server_t struct {
	httpServ *http.Server
	mux      *http.ServeMux
	port     string
}

func ResponseError(w http.ResponseWriter, mess string, status int) {
	errResp := ResponseResult_t{Message: mess}
	json, err := json.Marshal(errResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
