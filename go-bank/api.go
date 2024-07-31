package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

const secret = "asdf"

type APIServer struct {
	listenAddr string
	store      *PostgresStore
}

type APIError struct {
	Error string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func handleResponse(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func handleError(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			handleResponse(w, http.StatusBadRequest, APIError{Error: err.Error()})
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

	router.HandleFunc("/{id}", middleware(handleError(s.handleAccountByID)))
	router.HandleFunc("/", handleError(s.handleAccount))
	router.HandleFunc("/transfer", handleError(s.handleTransfer))
	http.ListenAndServe(s.listenAddr, router)
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		token, err := jwt.ParseWithClaims(tokenString, &AuthToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, fmt.Errorf("token not valid")
		}

		claims, ok := token.Claims.(*AuthToken)
		if !ok {
			return nil, fmt.Errorf("cannot get auth info")
		}

		if !(claims.ExpiresAt.After(time.Now()) || claims.ExpiresAt.Equal(time.Now())) {
			return nil, fmt.Errorf("token expired")
		}

		return []byte(secret), nil
	})
}

func createJWT(account *Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        account.ID,
		"expiresAt": time.Now().Add(time.Minute),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middlware initiated")

		token := strings.Split(r.Header.Get("Authorization"), " ")[1]
		if _, err := validateJWT(token); err != nil {
			handleResponse(w, http.StatusBadRequest, APIError{Error: "authorization failed"})
			return
		}

		f(w, r)
	}
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

	return handleResponse(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountDetails(w http.ResponseWriter, r *http.Request, id int) error {
	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	return handleResponse(w, http.StatusOK, account)
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

	// fmt.Println("Token:", token)

	return handleResponse(w, http.StatusCreated, "created")
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request, id int) error {
	if _, err := s.store.GetAccountByID(id); err != nil {
		return err
	}

	err := s.store.DeleteAccount(id)
	if err != nil {
		return err
	}
	return handleResponse(w, http.StatusOK, "deleted")
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}
	defer r.Body.Close()
	return handleResponse(w, http.StatusOK, transferReq)
}

func getID(r *http.Request) (int, error) {
	strID, ok := mux.Vars(r)["id"]
	if !ok {
		return -1, errors.New("invalid id")
	}

	id, _ := strconv.Atoi(strID)
	return id, nil
}
