package category

import (
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/stretchr/testify/mock"
)

type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) List() ([]*entity.Category, error) {
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

func (m *MockCategoryService) Create(data *entity.Category) error {
	args := m.Called(data)

	var r error
	if rf, ok := args.Get(0).(func() error); ok {
		r = rf()
	} else {
		r = args.Get(0).(error)
	}
	return r
}

func (m *MockCategoryService) GetByID(id int) (*entity.Category, error) {
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

func (m *MockCategoryService) UpdateTitle(data *entity.Category) error {
	args := m.Called(data)

	var r error
	if rf, ok := args.Get(0).(func() error); ok {
		r = rf()
	} else {
		r = args.Get(0).(error)
	}
	return r
}

func (m *MockCategoryService) UpdateParent(data *entity.Category) error {
	args := m.Called(data)

	var r error
	if rf, ok := args.Get(0).(func() error); ok {
		r = rf()
	} else {
		r = args.Get(0).(error)
	}
	return r
}

func (m *MockCategoryService) ClearParent(data *entity.Category) error {
	args := m.Called(data)

	var r error
	if rf, ok := args.Get(0).(func() error); ok {
		r = rf()
	} else {
		r = args.Get(0).(error)
	}
	return r
}

func (m *MockCategoryService) Delete(data *entity.Category) error {
	args := m.Called(data)

	var r error
	if rf, ok := args.Get(0).(func() error); ok {
		r = rf()
	} else {
		r = args.Get(0).(error)
	}
	return r
}