package models

import (
	"fmt"
	"mybank/db"
)

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func CreateUser(user *User) error {
	conn := db.DB

	query := "INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id"
	err := conn.QueryRow(query, user.Name, user.Balance).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}
	return nil
}

func UpdateUser(id string, user *User) error {
	conn := db.DB

	query := "UPDATE users SET name = $1, balance = $2 WHERE id = $3"
	_, err := conn.Exec(query, user.Name, user.Balance, id)
	if err != nil {
		return fmt.Errorf("could not update user: %v", err)
	}
	return nil
}

func DeleteUser(id string) error {
	conn := db.DB

	query := "DELETE FROM users WHERE id = $1"
	_, err := conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}
	return nil
}

func CreditBalance(id string, amount float64) error {
	conn := db.DB

	query := "UPDATE users SET balance = balance + $1 WHERE id = $2"
	_, err := conn.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("could not credit balance: %v", err)
	}
	return nil
}

func DebitBalance(id string, amount float64) error {
	conn := db.DB

	query := "UPDATE users SET balance = balance - $1 WHERE id = $2"
	_, err := conn.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("could not debit balance: %v", err)
	}
	return nil
}

func TransferMoney(from string, to string, amount float64) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("could not start transaction: %v", err)
	}

	var fromBalance float64
	if err := tx.QueryRow("SELECT balance FROM users WHERE id = $1", from).Scan(&fromBalance); err != nil {
		tx.Rollback()
		return fmt.Errorf("could not retrieve balance for user %s: %v", from, err)
	}

	if fromBalance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient funds in user %s's account", from)
	}

	if _, err := tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, from); err != nil {
		tx.Rollback()
		return fmt.Errorf("could not debit balance from user %s: %v", from, err)
	}

	if _, err := tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, to); err != nil {
		tx.Rollback()
		return fmt.Errorf("could not credit balance to user %s: %v", to, err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}
