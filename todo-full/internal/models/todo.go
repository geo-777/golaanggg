package models

import "time"

type Todo struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Completed bool      `json:"completed" db:"completed"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
