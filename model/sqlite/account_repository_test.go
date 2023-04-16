package sqlite

import (
    // "fmt"
    "testing"

    "agni.org/core/account"
)

func TestGetUserById(t *testing.T) {
    repo, db, err := MemorySQLiteRepo()
    if err != nil {
        t.Fatalf("failed to setup database: %s", err.Error())
    }

    defer db.Close()
    
    repo.Migrate()
    // [...]
}

func TestRegisterAccount(t *testing.T) {
    fixtureUser := account.User{
        Email: "sed.facilisis@outlook.org",
        Name:  "Neville Hooper",
        Id:    "nevhooper",
    }
    const fixturePasswd = "12345"

    repo, db, err := MemorySQLiteRepo()
    if err != nil {
        t.Fatalf("failed to setup database: %s", err.Error())
    }

    defer db.Close()
    if err := repo.Migrate(); err != nil {
        t.Fatalf("falied to migrate database: %s", err.Error())
    }
     
    err = repo.RegisterUser(fixtureUser, fixturePasswd) 
    if err != nil {
        t.Errorf(
            "failed to register user account: %s",
            err.Error(),
        )
    } 

    const query = `
        SELECT NAME, EMAIL
        FROM USERS 
        WHERE ID = ?
    `
    
    var user account.User

    row := db.QueryRow(query, fixtureUser.Id)
    if  err := row.Scan(&user.Name, &user.Email); err != nil {
        t.Errorf(
            "failed to retrieve user account: %s",
            err.Error(), 
        )
    }

    if user.Name != fixtureUser.Name {
        t.Errorf(
            "failed to assert name persistence\n expected: %s, actual: %s",
            fixtureUser.Name, user.Name,
        )
    }

    if user.Email != fixtureUser.Email {
        t.Errorf(
            "failed to assert email persistence\n expected: %s, actual: %s",
            fixtureUser.Email, user.Email,
        )
    }
}

func TestIsPasswdValid(t *testing.T) {
    fixtureUser := account.User{
        Email: "sed.facilisis@outlook.org",
        Name:  "Neville Hooper",
        Id:    "nevhooper",
    }
    const fixturePasswd  = "12345"
    const fixturePassMod = "7_6"

    repo, db, err := MemorySQLiteRepo()
    if err != nil {
        t.Fatalf("failed to setup database: %s", err.Error())
    }

    defer db.Close()
    if err := repo.Migrate(); err != nil {
        t.Fatalf("falied to migrate database: %s", err.Error())
    }
     
    err = repo.RegisterUser(fixtureUser, fixturePasswd) 
    if err != nil {
        t.Errorf(
            "failed to register user account: %s",
            err.Error(),
        )
    } 

    isPassValid, err := repo.IsPasswdValid(
        fixtureUser.Email,
        fixturePasswd, 
    )
    if err != nil {
        t.Errorf(
            "failed to authenticate password: %s",
            err.Error(),
        )
    }

    if !isPassValid {
       t.Errorf("failed to assert true password is valid")
    }

    isPassValid, err = repo.IsPasswdValid(
        fixtureUser.Email,
        fixturePasswd + fixturePassMod,
    )
    if err != nil {
        t.Errorf(
            "failed to authenticate password: %s",
            err.Error(),
        )
    }

    if isPassValid {
       t.Errorf("failed to assert false password is invalid")
    }
}
