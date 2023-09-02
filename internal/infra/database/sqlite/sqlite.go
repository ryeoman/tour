package sqlite

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// Database represents a database connection.
type Database struct {
	*sql.DB
}

// NewDatabase creates a new Database instance.
func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Set up connection pool parameters
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	return &Database{
		db,
	}, nil
}

// Close closes the database connection.
func (d *Database) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}

// ExecuteQuery executes a SQL query that doesn't return rows.
func (d *Database) ExecuteQuery(query string, args ...interface{}) error {
	_, err := d.DB.Exec(query, args)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return err
	}
	return nil
}

// QueryRow executes a SQL query that is expected to return a single row.
func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRow(query, args)
}

// QueryRows executes a SQL query that is expected to return multiple rows.
func (d *Database) QueryRows(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.DB.Query(query, args)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err
	}
	return rows, nil
}
