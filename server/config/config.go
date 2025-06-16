package config

import (
    "database/sql"
    "log"
    "os"
    "errors"

    "github.com/joho/godotenv"
    _ "github.com/jackc/pgx/v5/stdlib"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println(".env file not found, reading environment variables instead")
    }
}

func ConnectDB() (*sql.DB, error) {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        return nil, errors.New("DATABASE_URL is not set")
    }
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, err
    }
    // Test connection
    if err := db.Ping(); err != nil {
        return nil, errors.New("Failed to connect to the database: " + err.Error())
    }
    return db, nil
}
