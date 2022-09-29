package category

import (
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepo struct {
	mock.Mock
}

func (m *MockCategoryRepo) List() ([]*entity.Category, error) {
	args := m.Called()

	var r0 []*entity.Category
	if rf, ok := args.Get(0).(func() []*entity.Category); ok {
		r0 = rf()
	} else {
		r0 = args.Get(0).([]*entity.Category)
	}

	r1 := args.Error(1)

	return r0, r1
}

func (m *MockCategoryRepo) Create(data *entity.Category) error {
	args := m.Called(data)

	return args.Error(0)
}

func (m *MockCategoryRepo) GetByID(id int) (*entity.Category, error) {
	args := m.Called(id)

	var r0 *entity.Category
	if rf, ok := args.Get(0).(func() *entity.Category); ok {
		r0 = rf()
	} else {
		r0 = args.Get(0).(*entity.Category)
	}

	r1 := args.Error(1)

	return r0, r1
}

func (m *MockCategoryRepo) UpdateTitle(data *entity.Category) error {
	args := m.Called(data)

	return args.Error(0)
}

func (m *MockCategoryRepo) UpdateParent(data *entity.Category) error {
	args := m.Called(data)

	return args.Error(0)
}

func (m *MockCategoryRepo) ClearParent(data *entity.Category) error {
	args := m.Called(data)

	return args.Error(0)
}

func (m *MockCategoryRepo) Delete(data *entity.Category) error {
	args := m.Called(data)

	return args.Error(0)
}