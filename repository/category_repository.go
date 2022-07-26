package repository

import "go-lang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
