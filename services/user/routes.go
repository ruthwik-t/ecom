package user

import "net/http"

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /login", h.handleLogin)
	router.HandleFunc("GET /register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to LOGIN route"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to REGISTER route"))
}
