package usecase

import (
	"simpleBank/src/entity"
	"simpleBank/src/pkg/service_errors"
	"simpleBank/src/repository"
	"simpleBank/src/usecase/dto"
	"time"
)

type PayaUseCase struct {
	repo repository.IPayaRepository
}

func NewShebaUseCase(repo repository.IPayaRepository) *PayaUseCase {
	return &PayaUseCase{repo: repo}
}

func (u *PayaUseCase) CreatePayaRequest(dto *dto.CreateNewPaya) (*entity.Paya, error) {

	userBalance, err := u.repo.GetBalance(dto.From)
	if err != nil {
		return nil, err
	}

	if dto.Price > userBalance {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.InsufficientBalanceErr}
	}

	err = u.repo.DeductBalance(dto.From, dto.Price)
	if err != nil {
		return nil, err
	}

	dto.Id = GenerateID()
	dto.Status = entity.PendingStatus
	dto.CreatedAt = time.Now()

	paya := entity.Paya{
		Id:        dto.Id,
		Price:     dto.Price,
		FromSheba: dto.From,
		ToSheba:   dto.To,
		Status:    dto.Status,
		CreatedAt: dto.CreatedAt,
	}

	// Save to repository
	if err = u.repo.Save(&paya); err != nil {
		return nil, err
	}

	return &paya, nil
}

func (u *PayaUseCase) UpdatePayaRequest(id string, status, note string) (*entity.Paya, error) {
	paya, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if paya.Status != entity.PendingStatus {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.InvalidStatusErr}
	}

	paya.Status = status
	paya.Note = note

	if paya.Status == entity.CanceledStatus {
		err = u.repo.IncrementBalance(paya.FromSheba, paya.Price)
		if err != nil {
			return nil, err
		}
	}

	err = u.repo.Update(paya)

	if err != nil {
		return nil, err
	}

	return paya, nil
}

func (u *PayaUseCase) GetPayaRequests() []entity.Paya {
	return u.repo.GetAll()
}

func GenerateID() string {

	return time.Now().Format("20060102150405")
}
