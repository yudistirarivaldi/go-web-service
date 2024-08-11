package formatter

import (
	"web-service/model"
)

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func FormatUsers(users []model.User) []UserFormatter {
	var formattedUsers []UserFormatter
	for _, u := range users {
		formatter := UserFormatter{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Password: u.Password,
		}
		formattedUsers = append(formattedUsers, formatter)
	}
	return formattedUsers
}
