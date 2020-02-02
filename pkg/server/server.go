package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andreylm/vm_test/pkg/storage"

	customHanlders "github.com/andreylm/vm_test/pkg/server/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server - server
type Server struct {
	host     string
	port     string
	jsonFile string
	router   *mux.Router
	db       storage.IStorage
}

// NewServer - creates new server
func NewServer(host, port, jsonFile, connString string, useDB bool) (*Server, error) {
	var err error
	s := &Server{
		host:     host,
		port:     port,
		jsonFile: jsonFile,
		router:   mux.NewRouter(),
	}
	s.db, err = storage.GetStorage(useDB, connString)

	return s, err
}

// Start - startign server
func (s *Server) Start() error {
	defer s.db.Close()

	if err := s.init(); err != nil {
		return err
	}

	srv := &http.Server{
		Handler:      s.getHandler(),
		Addr:         fmt.Sprintf("%s:%s", s.host, s.port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s *Server) init() (err error) {
	if err = s.db.Load(s.jsonFile); err != nil {
		return
	}

	s.router.HandleFunc("/users", customHanlders.UserList(s.db)).Methods("GET")
	s.router.HandleFunc("/users/{id}", customHanlders.UserByID(s.db)).Methods("GET")
	return nil
}

func (s *Server) getHandler() http.Handler {
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control"}),
		handlers.ExposedHeaders([]string{"*"}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(s.router))

	return handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
}
