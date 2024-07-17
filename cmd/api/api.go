package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ruthwik-t/ecom/services/user"
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

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	fmt.Println("Server up and running at port", s.addr)

	return server.ListenAndServe()
}
