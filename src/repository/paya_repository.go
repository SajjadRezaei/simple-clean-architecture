package repository

import (
	"simpleBank/src/entity"
)

type IPayaRepository interface {
	Save(req *entity.Paya) error
	GetByID(id string) (*entity.Paya, error)
	GetBalance(sheba string) (int64, error)
	Update(req *entity.Paya) error
	GetAll() []entity.Paya
	DeductBalance(sheba string, amount int64) error
	IncrementBalance(sheba string, amount int64) error
}
