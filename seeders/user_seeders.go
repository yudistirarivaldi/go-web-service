package seeders

import (
	"database/sql"
)

func SeedUsers(db *sql.DB) error {
	users := []struct {
		Name     string
		Email    string
		Password string
	}{
		{Name: "Yudis", Email: "yudis@gmail.com", Password: "password123"},
		{Name: "Budi", Email: "budi@gmail.com", Password: "password456"},
	}

	for _, user := range users {
		_, err := db.Exec("INSERT IGNORE INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
		if err != nil {
			return err
		}
	}

	return nil
}
