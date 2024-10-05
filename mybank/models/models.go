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

// CreateUser inserts a new user into the database
func CreateUser(user *User) error {
	conn := db.Connect()
	defer conn.Close()

	query := "INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id"
	err := conn.QueryRow(query, user.Name, user.Balance).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}
	return nil
}

// UpdateUser updates an existing user
func UpdateUser(id string, user *User) error {
	conn := db.Connect()
	defer conn.Close()

	query := "UPDATE users SET name = $1, balance = $2 WHERE id = $3"
	_, err := conn.Exec(query, user.Name, user.Balance, id)
	if err != nil {
		return fmt.Errorf("could not update user: %v", err)
	}
	return nil
}

// DeleteUser deletes a user
func DeleteUser(id string) error {
	conn := db.Connect()
	defer conn.Close()

	query := "DELETE FROM users WHERE id = $1"
	_, err := conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}
	return nil
}

// CreditBalance adds an amount to the user's balance
func CreditBalance(id string, amount float64) error {
	conn := db.Connect()
	defer conn.Close()

	query := "UPDATE users SET balance = balance + $1 WHERE id = $2"
	_, err := conn.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("could not credit balance: %v", err)
	}
	return nil
}

// DebitBalance deducts an amount from the user's balance
func DebitBalance(id string, amount float64) error {
	conn := db.Connect()
	defer conn.Close()

	query := "UPDATE users SET balance = balance - $1 WHERE id = $2"
	_, err := conn.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("could not debit balance: %v", err)
	}
	return nil
}

// TransferMoney transfers money from one user to another
func TransferMoney(from string, to string, amount float64) error {
	conn := db.Connect()
	defer conn.Close()

	// Start a transaction
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Deduct from sender
	if _, err := tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, from); err != nil {
		return fmt.Errorf("could not debit balance from user %s: %v", from, err)
	}

	// Add to receiver
	if _, err := tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, to); err != nil {
		return fmt.Errorf("could not credit balance to user %s: %v", to, err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
