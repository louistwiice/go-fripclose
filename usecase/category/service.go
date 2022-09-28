package category

import (
	"github.com/louistwiice/go/fripclose/domain"
	"github.com/louistwiice/go/fripclose/entity"
)

type categoryservice struct {
	repo domain.CategoryRepository
}

func NewCategoryService(r domain.CategoryRepository) *categoryservice {
	return &categoryservice{
		repo: r,
	}
}

func (s *categoryservice) List() ([]*entity.Category, error) {
	return s.repo.List()
}

func (s *categoryservice) Create(data *entity.Category) error {
	return s.repo.Create(data)
}

func (s *categoryservice) GetByID(id int) (*entity.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryservice) UpdateTitle(data *entity.Category) error {
	return s.repo.UpdateTitle(data)
}

func (s *categoryservice) UpdateParent(data *entity.Category) error {
	return s.repo.UpdateParent(data)
}

func (s *categoryservice) ClearParent(data *entity.Category) error {
	return s.repo.ClearParent(data)
}


func (s *categoryservice) Delete(data *entity.Category) error {
	return s.repo.Delete(data)
}