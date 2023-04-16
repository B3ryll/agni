package sqlite

import (
    "errors"
    "fmt"

    "agni.org/core/account"
)


func (repo *SQLiteRepository) GetUserById(id account.ID) (*account.User, error) {
    const query = `
        SELECT NAME, EMAIL
        FROM USERS
        WHERE ID = ?;
    `

    user := account.User{Id: id}

    row := repo.db.QueryRow(query, string(id))
    if err := row.Scan(&user.Name, &user.Email); err != nil {
        msg := fmt.Sprintf(
            "failed to fetch user from db: %s",
            err.Error(),
        )
        return nil, errors.New(msg)
    }

    return &user, nil
}

func (repo *SQLiteRepository) RegisterUser(usr account.User, pswd string) error {
    const query = `
        INSERT INTO USERS(NAME, EMAIL, ID, PASSWD)
        VALUES (?, ?, ?, ?);
    `

    passwdHash := account.HashPassword(pswd)
    _, err := repo.db.Exec(
        query,
        usr.Name,
        usr.Email,
        usr.Id,
        passwdHash,
    )
    if err != nil {
        msg := fmt.Sprintf("failed to register user: %s", err.Error())
        return errors.New(msg)
    }

    return nil
}

func (repo *SQLiteRepository) IsPasswdValid(email, passwd string) (bool, error) {
    const query = `
        SELECT PASSWD
        FROM USERS
        WHERE EMAIL = ?;
    `

    passwdHash := account.HashPassword(passwd)

    var persistedHash string
    row := repo.db.QueryRow(query, email)
    if err := row.Scan(&persistedHash); err != nil {
        msg := fmt.Sprintf(
            "failed retrived password hash from db: %s",
            err.Error(),
        )
        return false, errors.New(msg)
    }
    
    return passwdHash == account.Password(persistedHash), nil
}
