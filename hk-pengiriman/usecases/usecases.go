package usecases

import (
	"errors"
	"hk-pengiriman/helpers/request"
	"hk-pengiriman/model"
	"hk-pengiriman/repositories"
)

// Usecases ...
type Usecases interface {
	CreateOne(m *model.HKPengiriman) (int64, error)
	UpdateOneByID(id int64, m *model.HKPengiriman) (int64, error)
	GetOneByID(id int64) (*model.HKPengiriman, int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(filter *request.QueryParameter) ([]*model.HKPengiriman, int64, error)
}

type usecases struct {
	repo repositories.Repositories
}

// NewUsecases ...
func NewUsecases() Usecases {
	return &usecases{
		repo: repositories.NewRepositories(),
	}
}

func (v *usecases) CreateOne(m *model.HKPengiriman) (int64, error) {
	return v.repo.CreateOne(m)
}

func (v *usecases) UpdateOneByID(id int64, m *model.HKPengiriman) (int64, error) {
	if id == 0 {
		return -1, errors.New("id cannot be 0")
	}

	return v.repo.UpdateOneByID(id, m)
}

func (v *usecases) GetOneByID(id int64) (*model.HKPengiriman, int64, error) {
	if id == 0 {
		return nil, -1, errors.New("id cannot be 0")
	}

	return v.repo.GetOneByID(id)
}

func (v *usecases) DeleteOneByID(id int64) (int64, error) {
	if id == 0 {
		return -1, errors.New("id cannot be 0")
	}

	return v.repo.DeleteOneByID(id)
}

func (v *usecases) GetAll(filter *request.QueryParameter) ([]*model.HKPengiriman, int64, error) {
	return v.repo.GetAll(filter)
}
