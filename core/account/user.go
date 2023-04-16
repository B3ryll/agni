package account

type ID string

type User struct {
    Id    ID
    Email string
    Name  string
}

type UserRepository interface {
    GetUserById(Id ID) (*User, error)
    GetUserByEmail(email string) (*User, error)

    UpdateUserName(name string) (*User, error)
    UpdateUserEmail(email string) (*User, error)

    RemoveUser(id ID) error
}
