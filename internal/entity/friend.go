package entity

import "time"

type Friend struct {
	ID        string
	Email     string
	Phone     string
	Name      string
	Password  string
	CreatedAt time.Time
}
