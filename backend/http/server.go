package http

import (
	"encoding/json"
	"errors"
	"github.com/rs/cors"
	"log"
	"net/http"

	patient "manage-patients"

	"github.com/go-chi/chi/v5"
)

// Server is an HTTP server which embeds a chi router
type Server struct {
	router      *chi.Mux
	patientRepo patient.PatientRepository
}

// NewServer creates and configures a new Server instance
func NewServer(patientRepo patient.PatientRepository, p ...Pinger) *Server {
	s := Server{
		router:      chi.NewRouter(),
		patientRepo: patientRepo,
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // 允许所有来源，根据您的需求调整
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	s.router.Use(corsMiddleware.Handler)
	s.routes(p...)
	return &s
}

// ErrorResponse is responsible to include multiple errors
type ErrorResponse struct {
	Errors []string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(r *http.Request, w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	var body interface{}
	switch v := data.(type) {
	case error:
		res := ErrorResponse{
			Errors: []string{},
		}

		for ce := v; ce != nil; ce = errors.Unwrap(ce) {
			res.Errors = append(res.Errors, ce.Error())
		}
		body = res
	case *[]patient.Patient, *patient.Order:
		body = data
	default:
		body = nil
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Fatalln(err)
		}
	}
}
