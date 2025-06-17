package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func ExecSQLFile(t *testing.T, db *sql.DB, filepathStr string) {
	abs, _ := filepath.Abs(filepathStr)
	fmt.Printf("Trying to open: %s\n", abs)

	sqlBytes, err := os.ReadFile(filepathStr)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", filepathStr, err)
	}
	sqlText := string(sqlBytes)
	stmts := strings.Split(sqlText, ";")
	for _, stmt := range stmts {
		stmt = strings.TrimSpace(stmt)
		if stmt != "" {
			if _, err := db.Exec(stmt); err != nil {
				t.Fatalf("Failed to execute statement from %s: %v\nSQL: %s", filepathStr, err, stmt)
			}
		}
	}
}
