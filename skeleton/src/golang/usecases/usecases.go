package usecases

import (
	"errors"
	"{{ toDelimeted .ProjectName 45 }}/helpers/request"
	"{{ toDelimeted .ProjectName 45 }}/model"
	"{{ toDelimeted .ProjectName 45 }}/repositories"
)

// Usecases ...
type Usecases interface {
	CreateOne(m *model.{{ toCamel .ProjectName }}) (int64, error)
	UpdateOneByID(id int64, m *model.{{ toCamel .ProjectName }}) (int64, error)
	GetOneByID(id int64) (*model.{{ toCamel .ProjectName }}, int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(filter *request.QueryParameter) ([]*model.{{ toCamel .ProjectName }}, int64, error)
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

func (v *usecases) CreateOne(m *model.{{ toCamel .ProjectName }}) (int64, error) {
	return v.repo.CreateOne(m)
}

func (v *usecases) UpdateOneByID(id int64, m *model.{{ toCamel .ProjectName }}) (int64, error) {
	if id == 0 {
		return -1, errors.New("id cannot be 0")
	}

	return v.repo.UpdateOneByID(id, m)
}

func (v *usecases) GetOneByID(id int64) (*model.{{ toCamel .ProjectName }}, int64, error) {
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

func (v *usecases) GetAll(filter *request.QueryParameter) ([]*model.{{ toCamel .ProjectName }}, int64, error) {
	return v.repo.GetAll(filter)
}
