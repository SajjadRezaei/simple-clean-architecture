package dto

import (
	"errors"
	"regexp"
	"simpleBank/src/usecase/dto"
)

type PayaRequest struct {
	Price int64  `json:"price"`
	From  string `json:"fromShebaNumber"`
	To    string `json:"toShebaNumber"`
	Note  string `json:"note"`
}

type UpdatePayaRequest struct {
	Status string `json:"status"`
	Note   string `json:"note"`
}

func (r *PayaRequest) Validate() error {
	if r.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	return validateSheba(r)
}

func validateSheba(r *PayaRequest) error {
	shebaRegex := `^IR\d{24}$`
	if !regexp.MustCompile(shebaRegex).MatchString(r.From) || !regexp.MustCompile(shebaRegex).MatchString(r.To) {
		return errors.New("invalid sheba number format")
	}
	return nil
}

func ToCreatePaya(paya PayaRequest) *dto.CreateNewPaya {
	return &dto.CreateNewPaya{
		Price: paya.Price,
		From:  paya.From,
		To:    paya.To,
		Note:  paya.Note,
	}
}
