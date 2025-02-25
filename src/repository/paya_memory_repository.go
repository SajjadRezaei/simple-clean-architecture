package repository

import (
	"errors"
	"simpleBank/src/entity"
	"slices"
	"sort"
	"sync"
)

type InMemoryPayaRepository struct {
	requests     []entity.Paya
	mu           sync.RWMutex
	userBalances map[string]int64
}

func NewInMemoryPayaRepository() *InMemoryPayaRepository {
	return &InMemoryPayaRepository{
		requests: make([]entity.Paya, 0),
		userBalances: map[string]int64{
			"IR123456789012345678901234": 50000000000000,
		},
	}
}

func (r *InMemoryPayaRepository) Save(req *entity.Paya) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.requests = append(r.requests, *req)

	return nil
}

func (r *InMemoryPayaRepository) GetByID(id string) (*entity.Paya, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	idx := slices.IndexFunc(r.requests, func(p entity.Paya) bool { return p.Id == id })

	if idx == -1 {
		return nil, errors.New("request not found")
	}

	return &r.requests[idx], nil
}

func (r *InMemoryPayaRepository) Update(req *entity.Paya) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	idx := slices.IndexFunc(r.requests, func(p entity.Paya) bool { return p.Id == req.Id })

	if idx == -1 {
		return errors.New("request not found")
	}

	r.requests[idx] = *req
	return nil
}

func (r *InMemoryPayaRepository) GetAll() []entity.Paya {
	r.mu.RLock()
	defer r.mu.RUnlock()

	sort.Slice(r.requests, func(i, j int) bool {
		return r.requests[i].CreatedAt.Before(r.requests[j].CreatedAt)
	})

	return r.requests
}

func (r *InMemoryPayaRepository) GetBalance(sheba string) (int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	balance, ok := r.userBalances[sheba]
	if !ok {
		return 0, errors.New("account not found")
	}

	return balance, nil
}

func (r *InMemoryPayaRepository) DeductBalance(sheba string, amount int64) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	balance, ok := r.userBalances[sheba]
	if !ok {
		return errors.New("account not found")
	}

	if balance >= amount {
		r.userBalances[sheba] -= amount
	}

	return nil

}

func (r *InMemoryPayaRepository) IncrementBalance(sheba string, amount int64) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.userBalances[sheba]

	if !ok {
		return errors.New("account not found")
	}

	r.userBalances[sheba] += amount

	return nil
}
