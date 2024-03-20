package entity

import "time"

type Comment struct {
	ID        string    `json:"-"`
	Comment   string    `json:"comment"`
	PostId    string    `json:"-"`
	UserId    string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
