package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {

	router := http.NewServeMux()
	subrouter := http.NewServeMux()

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", subrouter))

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	fmt.Println("Server up and running at port", s.addr)

	return server.ListenAndServe()
}
