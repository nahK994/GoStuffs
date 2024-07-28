package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

	router.HandleFunc("/{id}", makeHTTPhandler(s.handleAccountByID))
	router.HandleFunc("/", makeHTTPhandler(s.handleAccount))
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	if r.Method == "GET" {
		return s.handleGetAccountDetails(w, r, id)
	} else if r.Method == "PUT" {
		return s.handleUpdateAccount(w, r, id)
	} else if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r, id)
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccountList(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request, id int) error {
	return nil
}

func (s *APIServer) handleGetAccountList(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return handleHTTPResponse(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountDetails(w http.ResponseWriter, r *http.Request, id int) error {
	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	return handleHTTPResponse(w, http.StatusOK, account)
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

	return handleHTTPResponse(w, http.StatusCreated, "created")
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request, id int) error {
	err := s.store.DeleteAccount(id)
	if err != nil {
		return err
	}
	return handleHTTPResponse(w, http.StatusNoContent, "deleted")
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func getID(r *http.Request) (int, error) {
	strID, ok := mux.Vars(r)["id"]
	if !ok {
		return -1, errors.New("invalid id")
	}

	id, _ := strconv.Atoi(strID)
	return id, nil
}
