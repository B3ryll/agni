package account

import (
    "encoding/base64"
    "crypto/sha256"
    "bytes"
)

// type for password hash
type Password string

type Account struct {
    User
    Passwd Password
}

type AccountRepository interface {
    RegisterAccount(user User, passwd string) error
    
    IsPasswordValid(email, passwd string) (bool, error)
    ChangePassword(passwd string) (error)
}

// convert plain password string to base64 sha256 encrypted string
func HashPassword(plain string) Password {
    hash := sha256.New()
    hash.Write([]byte(plain))

    buf     := new(bytes.Buffer)
    encoder := base64.NewEncoder(base64.StdEncoding, buf)

    encoder.Write(hash.Sum(nil))
    
    return Password(buf.String())
}

func (account *Account) DoesPasswdMatch(passwd string) bool {
    return HashPassword(passwd) == account.Passwd
}
