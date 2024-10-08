package model

import (
	"database/sql"
)

var DB *sql.DB

type User struct {
    ID    int  
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}

func GetAllUsers() ([]User, error) {
    rows, err := DB.Query("SELECT id, name, email, password FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func GetUserByID(id int) (User, error) {
    row := DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
    var user User
    if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
        return user, err
    }
    return user, nil
}

func CreateUser(user User) (int64, error) {
    result, err := DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func UpdateUser(user User) error {
    _, err := DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
    return err
}

func DeleteUser(id int) error {
    _, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}
