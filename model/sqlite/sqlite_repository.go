package sqlite

import (
    "database/sql"
    "errors"
    "fmt"

    _ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
    db *sql.DB
}

func NewSQLiteRepo(db *sql.DB ) SQLiteRepository{
    return SQLiteRepository{db: db}
}

func MemorySQLiteRepo() (*SQLiteRepository, *sql.DB, error) {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        msg :=  fmt.Sprintf(
            "failed to setup database: %s", err.Error(),
        )
        return nil, nil, errors.New(msg) 
    }

    repo := NewSQLiteRepo(db)
    return &repo, db, nil
}

func (repo *SQLiteRepository) Migrate() error {
    if err := createUsersTable(repo.db); err != nil {
        msg := fmt.Sprintf("migration failed: %s", err.Error())
        return errors.New(msg)
    }

    return nil
}

func createUsersTable(db *sql.DB) error {
    const query = `
        CREATE TABLE IF NOT EXISTS USERS(
            ID TEXT PRIMARYKEY NOT NULL,
            EMAIL TEXT NOT NULL,
            NAME TEXT NOT NULL,
            PASSWD TEXT NOT NULL
        );
    `

    _, err := db.Exec(query)
    if err != nil {
        return err
    }

    return nil
}
