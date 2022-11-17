package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

var RWmtx sync.RWMutex
var wg sync.WaitGroup

type Info struct {
	UserId      string    `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	taskid      int
}

func Logs(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

type Event struct {
	UserId string
	inf    []Info
}

var Events = make(map[string]*Event, 5)

func AddEvent(event Event) error {
	RWmtx.Lock()
	defer RWmtx.Unlock()
	isAlready := false
	if _, ok := Events[event.UserId]; ok {
		for _, val := range Events[event.UserId].inf {
			if val.Description == event.inf[0].Description && val.Date == event.inf[0].Date {
				isAlready = true
				break
			}
		}
		if !isAlready {
			Events[event.UserId].inf = append(Events[event.UserId].inf, event.inf[0])
			return nil
		} else {
			return errors.New("Event already exists")
		}

	}

	Events[event.UserId] = &event
	return nil
}

func UpdateInfo(event Event) error {
	RWmtx.Lock()
	defer RWmtx.Unlock()
	isUpdate := false
	if _, ok := Events[event.UserId]; ok {
		for _, v := range Events[event.UserId].inf {
			for _, ev := range event.inf {
				if v.Date == ev.Date {
					Events[event.UserId] = &event
					isUpdate = true
					return nil
				}
			}

		}
		if !isUpdate {
			Events[event.UserId].inf = append(Events[event.UserId].inf, event.inf[0])
			return nil
		}

	}
	return errors.New("Event no find")
}

func DeleteEvent(event Event) error {
	RWmtx.Lock()
	defer RWmtx.Unlock()
	if _, ok := Events[event.UserId]; ok {
		delete(Events, event.UserId)
		return nil
	}
	return errors.New("Event no find")
}

func (e *Event) Validate() error {
	id, err := strconv.Atoi(e.UserId)
	if err != nil {
		return errors.New("Wrong id")
	}
	if e.UserId == "" || id < 0 {
		return errors.New("Wrong event")
	}

	for _, v := range e.inf {
		if v.Description == "" {
			return errors.New("Wrong event")
		}
	}

	return nil
}

func NewServer(port string) Server_t {
	return Server_t{
		mux:  http.NewServeMux(),
		port: port,
	}
}

func (s *Server_t) Run() {
	defer wg.Done()
	s.mux.HandleFunc("/create_event", Logs(http.HandlerFunc(s.CreateEvent)))
	s.mux.HandleFunc("/update_event", Logs(http.HandlerFunc(s.UpdateEvent)))
	s.mux.HandleFunc("/delete_event", Logs(http.HandlerFunc(s.DeleteEvent)))

	s.mux.HandleFunc("/events_for_day", Logs(http.HandlerFunc(s.EventForDay)))
	s.mux.HandleFunc("/events_for_week", Logs(http.HandlerFunc(s.EventForWeek)))
	s.mux.HandleFunc("/events_for_month", Logs(http.HandlerFunc(s.EventForMonth)))
	s.httpServ = &http.Server{
		Addr:    s.port,
		Handler: s.mux,
	}

	fmt.Println("Server start")
	if err := s.httpServ.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}

}

func (s *Server_t) WaitStop() {
	defer wg.Done()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	select {
	case <-quit:
		fmt.Println("Server stopped")
		s.httpServ.Shutdown(nil)
	}
}

func main() {
	serv := NewServer(":16666")
	wg.Add(2)
	go serv.WaitStop()
	serv.Run()
	wg.Wait()
}
