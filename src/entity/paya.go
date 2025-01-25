package entity

import (
	"time"
)

type Paya struct {
	Id        string    `json:"id"`
	Price     int64     `json:"price"`
	FromSheba string    `json:"fromShebaNumber"`
	ToSheba   string    `json:"toShebaNumber"`
	Note      string    `json:"note"`
	Status    string    `json:"status"` // pending, confirmed, canceled
	CreatedAt time.Time `json:"createdAt"`
}
