package main

import (
	"database/sql"
	"fmt"
	"os"
)

func runSQLScript(db *sql.DB, filePath string) error {
	script, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		return fmt.Errorf("error executing SQL script: %v", err)
	}

	return nil
}