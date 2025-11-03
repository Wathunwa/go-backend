package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/register", h.Register).Methods("POST")
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Handle user login
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// Handle user registration
}
