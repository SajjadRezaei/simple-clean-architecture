package dto

import (
	"time"
)

type CreateNewPaya struct {
	Id        string
	Price     int64
	From      string
	To        string
	Status    string
	Note      string
	CreatedAt time.Time
}
