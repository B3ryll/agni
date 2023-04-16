package sqlite

import (
    "testing"
)

func TestMigrate(t *testing.T) {
    repo, db, err := MemorySQLiteRepo()
    if err != nil {
        t.Fatalf("failed to setup database: %s", err.Error())
    }
    
    defer db.Close()

    if err := repo.Migrate(); err != nil {
        t.Fatalf("failed to migrate: %s", err.Error())
    }
    
    const query = `
        SELECT COUNT(*)
        FROM SQLITE_MASTER
        WHERE type='table' and name=?;
    `

    var userTableCount int
    row := db.QueryRow(query, "USERS")
    if err := row.Scan(&userTableCount); err != nil {
        t.Errorf("failed to query user table: %s", err.Error())
    }
    
    if userTableCount != 1 {
        t.Errorf(
            "failed to assert user table existence: count = %d",
            userTableCount,
        )
    } 
}
