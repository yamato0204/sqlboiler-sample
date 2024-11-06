package database

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

type DBConfig struct {
	User     string
	DBName   string
	RootPass string
	Host     string
	Port     string
}

func NewDB(cfg *DBConfig) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.RootPass, cfg.Host, cfg.DBName)
	log.Printf("dsn: %s", dsn)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	log.Println("db connection success")

	return &DB{db}
}

// ListTables retrieves all table names in the database and returns them as a comma-separated string.
func (db *DB) ListTables() (string, error) {
	var tables []string
	query := "SHOW TABLES"
	err := db.Select(&tables, query)
	if err != nil {
		return "", fmt.Errorf("failed to list tables: %w", err)
	}

	// Join table names with comma and space
	tableNames := strings.Join(tables, ", ")
	return tableNames, nil
}
