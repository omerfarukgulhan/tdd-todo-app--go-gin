package domain

import (
	"time"
)

type Todo struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
