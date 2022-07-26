package service

import (
	"errors"
	"go-lang-unit-test/entity"
	"go-lang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
