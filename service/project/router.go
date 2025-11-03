package project

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	DB *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Handler { return &Handler{DB: db} }

// Mount ผูกเส้นทาง /api/v1/project
func (h *Handler) Mount(r *mux.Router) {
	s := r.PathPrefix("/project").Subrouter()
	s.HandleFunc("", h.List).Methods(http.MethodGet) // GET /api/v1/project
	// s.HandleFunc("", h.Create).Methods(http.MethodPost)  // POST /api/v1/users
	// s.HandleFunc("/{id}", h.Get).Methods(http.MethodGet) // GET /api/v1/users/{id}
}

type dataProject struct { // create struct keep data from db
	Project_id   string `json:"project_id"`
	Project_name string `json:"project_name"`
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	rows, err := h.DB.Query(ctx, `
	  SELECT project_id,project_name
	  FROM tbl_project_raw
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer rows.Close()

	fmt.Println("Check data => ", rows)
	var out []dataProject
	for rows.Next() {
		var u dataProject
		if err := rows.Scan(&u.Project_id, &u.Project_name); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		out = append(out, u)
	}

	_ = json.NewEncoder(w).Encode(out)

}
