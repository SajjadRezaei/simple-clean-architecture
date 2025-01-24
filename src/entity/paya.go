package entity

import (
	"errors"
	"regexp"
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

func (s *Paya) Validate() error {
	if s.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	shebaRegex := `^IR\d{24}$`
	if !regexp.MustCompile(shebaRegex).MatchString(s.FromSheba) || !regexp.MustCompile(shebaRegex).MatchString(s.ToSheba) {
		return errors.New("invalid sheba number format")
	}

	return nil
}
