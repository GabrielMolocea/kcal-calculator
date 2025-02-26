package utils

import (
    "database/sql"
    "fmt"
    "os"
)

// RunSQLScript reads and executes an SQL script from a file
func RunSQLScript(db *sql.DB, filePath string) error {
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