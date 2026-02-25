package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/OlegRozh/Final-progect-todo/internal/service"
)

type NextDateRequest struct {
	Now    string
	Date   string
	Repeat string
}

func (r *NextDateRequest) ParseFromRequest(req *http.Request) error {
	r.Now = req.FormValue("now")
	r.Date = req.FormValue("date")
	r.Repeat = req.FormValue("repeat")
	if r.Date == "" {
		return errors.New("date start cannot be empty")
	}
	if r.Repeat == "" {
		return errors.New("repeat cannot be empty")
	}
	return nil
}

func NextDate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	var req NextDateRequest
	if err := req.ParseFromRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var now time.Time
	var err error
	if req.Now == "" {
		now = time.Now()
	} else {
		now, err = time.Parse("20060102", req.Now)
		if err != nil {
			http.Error(w, "Invalid now format", http.StatusBadRequest)
			return
		}
	}
	nextDate, err := service.NextDate(now, req.Date, req.Repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
