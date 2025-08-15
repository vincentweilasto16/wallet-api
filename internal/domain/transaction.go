package domain

import (
	"time"
)

type Transaction struct {
	ID              string
	UserID          string
	Amount          string
	Type            string
	Status          string
	Description     string
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
