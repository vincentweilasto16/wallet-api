package domain

import (
	"time"
)

type User struct {
    ID        string
    Name      string
    Email     string
    Balance   string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}