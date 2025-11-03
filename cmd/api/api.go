package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux" // ตัวทำ router
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Wathunwa/go-backend.git/service/project"
)

type APIServer struct {
	addr   string
	router *mux.Router
	db     *pgxpool.Pool
} // Create struct เพื่อเก็บข้อมูล DB

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	if db == nil {
		panic("DB pool is nil")
	}
	serv := &APIServer{
		addr:   addr,
		router: mux.NewRouter(),
		db:     db,
	}

	serv.registerRouters()
	return serv
}

func (s *APIServer) registerRouters() {
	//middleware  CORS/JSON
	s.router.Use(commonHeaders)

	//Health check
	s.router.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	}).Methods(http.MethodGet)

	//version.
	v1 := s.router.PathPrefix("/api/v1").Subrouter()

	//Module
	project.New(s.db).Mount(v1) // -> api/v1/project
}

// basic middleware
func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Run() error {
	srv := &http.Server{
		Addr:         s.addr,
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Println("HTTP listening on", s.addr)
	return srv.ListenAndServe()
}
