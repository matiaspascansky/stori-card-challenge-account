package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDBConnection() (*sql.DB, error) {
	dbHost := os.Getenv("rds_host")
	dbUser := os.Getenv("rds_user")
	dbPassword := os.Getenv("rds_password")
	dbName := os.Getenv("rds_name")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	//SE SUPONE QUE SE CONECTA
	return db, nil
}
