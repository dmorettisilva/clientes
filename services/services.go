package services

import (
	"log"
	"net/http"

	"github.com/dmorettisilva/clientes/config"
	"github.com/dmorettisilva/clientes/database"
	"github.com/gorilla/mux"

	_ "github.com/jmoiron/sqlx"
)

//Service - struct que representa os services
type Service struct {
	cfg *config.Config
	db  *database.DAO
}

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

func New(cfg *config.Config, dao *database.DAO) *Service {
	return &Service{cfg: cfg, db: dao}
}

func (s *Service) RouteDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "message": "OK" }`))
}

func (s *Service) authAPI(next httpHandlerFunc) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header != s.cfg.Authorization {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{ "errors": "401", "message": "invalid authorization" }`))
			return
		}
		next(w, r)
	}
}

func (s *Service) Listen() {
	router := mux.NewRouter()

	router.HandleFunc("/", (s.RouteDefault)).Methods("GET")
	router.HandleFunc("/cliente/{idpessoa}", s.authAPI(s.ConsultaClienteHandler)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8089", router))
}
