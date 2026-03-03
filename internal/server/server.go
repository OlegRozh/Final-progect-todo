package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OlegRozh/Final-progect-todo/internal/api"
	"github.com/OlegRozh/Final-progect-todo/internal/storage"
)

type Server struct {
	store  storage.TaskStorage
	port   string
	router *http.ServeMux
}

func New(store storage.TaskStorage, port string) *Server {
	return &Server{
		store:  store,
		port:   port,
		router: http.NewServeMux(),
	}
}

func (s *Server) SetupRoutes() {
	s.router.Handle("/", http.FileServer(http.Dir("./web")))
	s.router.HandleFunc("/api/task", api.TaskHandler(s.store))
	s.router.HandleFunc("/api/tasks", api.GetListTasks(s.store))
	s.router.HandleFunc("/api/task/done", api.StatusDoneHandler(s.store))
	s.router.HandleFunc("/api/nextdate", api.NextDateHandler)
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.port)
	log.Printf("Server starting on %s", addr)
	start := http.ListenAndServe(addr, s.router)
	return start
}
