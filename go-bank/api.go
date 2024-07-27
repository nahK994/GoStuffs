package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      *PostgresStore
}

type APIError struct {
	Error string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func handleHTTPResponse(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPhandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			handleHTTPResponse(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string, store *PostgresStore) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", makeHTTPhandler(s.handleAccount))
	router.HandleFunc("/{id}", makeHTTPhandler(s.handleGetAccount))
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	// vars, ok := mux.Vars(r)["id"]
	// if ok {
	// 	return handleHTTPResponse(w, http.StatusOK, vars)
	// } else {
	// 	account := NewAccount("Shomi", "Khan")
	// 	return handleHTTPResponse(w, http.StatusOK, account)
	// }
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return handleHTTPResponse(w, http.StatusOK, accounts)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	accountInfo := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(accountInfo); err != nil {
		return err
	}
	account := NewAccount(accountInfo.FirstName, accountInfo.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return handleHTTPResponse(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
