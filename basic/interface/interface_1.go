package interface_example

import "fmt"

// DBConnection defines the interface for database connections
type DBConnection interface {
	Connect() error
	Close() error
}

// PostgreSQL implements DBConnection interface
type PostgreSQL struct {
	dsn string
}

func (p *PostgreSQL) Connect() error {
	fmt.Println("Connecting to PostgreSQL with DSN:", p.dsn)
	// Actual connection logic...
	return nil
}

func (p *PostgreSQL) Close() error {
	fmt.Println("Closing PostgreSQL connection")
	// Actual close logic...
	return nil
}

// MySQL implements DBConnection interface
type MySQL struct {
	dsn string
}

func (m *MySQL) Connect() error {
	fmt.Println("Connecting to MySQL with DSN:", m.dsn)
	// Actual connection logic...
	return nil
}

func (m *MySQL) Close() error {
	fmt.Println("Closing MySQL connection")
	// Actual close logic...
	return nil
}

func Interface1() {
	var db DBConnection

	db = &PostgreSQL{dsn: "postgres://user:pass@localhost/dbname"}
	db.Connect()
	db.Close()

	db = &MySQL{dsn: "mysql://user:pass@localhost/dbname"}
	db.Connect()
	db.Close()
}
