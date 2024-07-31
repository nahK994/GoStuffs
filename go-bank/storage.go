package main

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(int) error
	DeleteAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgres://user:password@localhost:5432/go_bank?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (p *PostgresStore) Init() error {
	return p.CreateAccountTable()
}

func (p *PostgresStore) CreateAccountTable() error {
	query := `
		create table if not exists account(
			id serial primary key,
			first_name varchar(50),
			last_name varchar(50),
			number serial,
			balance serial,
			password varchar(200),
			created_at timestamp
		)
	`
	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresStore) CreateAccount(a *Account) error {
	query := `
		insert into account (first_name, last_name, number, balance, password, created_at)
		values ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.Query(query, a.FirstName, a.LastName, a.Number, a.Balance, a.Password, a.CreatedAt)
	return err
}

func (p *PostgresStore) UpdateAccount(id int) error {
	return nil
}

func (p *PostgresStore) DeleteAccount(id int) error {
	query := `delete from account where id=$1`
	if _, err := p.db.Query(query, id); err != nil {
		return err
	}
	return nil
}

func (p *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := `select * from account where id=$1`
	rows, err := p.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("account not found")
	}

	account, err := scanRow(rows)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (p *PostgresStore) GetAccountByAccNumber(number int) (*Account, error) {
	query := `select id, password from account where number=$1`
	rows, err := p.db.Query(query, number)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("account not found")
	}

	account := new(Account)
	if err := rows.Scan(&account.ID, &account.Password); err != nil {
		return nil, err
	}
	return account, nil
}

func (p *PostgresStore) GetAccounts() ([]*Account, error) {
	query := `select * from account`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	var accounts []*Account
	for rows.Next() {
		account, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanRow(row *sql.Rows) (*Account, error) {
	account := new(Account)
	if err := row.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.Password, &account.CreatedAt); err != nil {
		return nil, err
	}
	return account, nil
}
