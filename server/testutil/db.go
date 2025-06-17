package testutil

import (
	"database/sql"
	"os"
	"sync"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	testDB     *sql.DB
	testDBOnce sync.Once
)

func GetTestDB(t *testing.T) *sql.DB {
	testDBOnce.Do(func() {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			t.Fatal("DATABASE_URL not set")
		}
		var err error
		testDB, err = sql.Open("pgx", dsn)
		if err != nil {
			t.Fatalf("Failed to connect to test DB: %v", err)
		}
	})
	return testDB
}
