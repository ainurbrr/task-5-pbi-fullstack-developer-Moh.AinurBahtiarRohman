package user

import "time"

type UserFormatter struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Token       string    `json:"token"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.Phone_Number,
		Token:       token,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return formatter
}
