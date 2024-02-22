package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string `json:"error"`
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type Server struct {
	PortNumber string
}

func NewApiServer(port string) *Server {
	return &Server{
		PortNumber: port,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))
	log.Println("JSON API server running on port:", s.PortNumber)
	http.ListenAndServe(s.PortNumber, router)
}

func (s *Server) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return nil
	}
}

func (s *Server) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
