package user

import "time"

type User struct {
	ID              int
	Username        string
	Name            string
	Email           string
	Password        string
	Avatar          string
	Phone_Number     string
	City            string
	Address         string
	Roles           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
