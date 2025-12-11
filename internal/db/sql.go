package db

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "github.com/jmoiron/sqlx"
)

func Open(path string) (*sqlx.DB, error) {
    dsn := fmt.Sprintf("file:%s?_busy_timeout=5000&_fk=1", path)
    db, err := sqlx.Open("sqlite3", dsn)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
