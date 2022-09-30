package service_category

import (
	"errors"
	"testing"

	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/mocks"
	"github.com/louistwiice/go/fripclose/mocks/category"
	"github.com/stretchr/testify/assert"
)

func Test_List_Category(t *testing.T) {
	c := mocks.GenerateFixture().CategoryList

	repo := category.MockCategoryRepo{}
	repo.On("List").Return(c, nil)

	service := NewCategoryService(&repo)
	resp, err := service.List()

	assert.Nil(t, err)
	assert.Equal(t, 4, len(resp))
	assert.Equal(t, c, resp)
}

func Test_Create_Category(t *testing.T) {
	t.Run("Should return an error when repo returns an error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("Create", c).Return(errors.New("something happens"))
	
		service := NewCategoryService(&repo)
		err := service.Create(c)
		assert.NotNil(t, err)
	})

	t.Run("Should be ok when repo returns no error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("Create", c).Return(nil)
	
		service := NewCategoryService(&repo)
		err := service.Create(c)
		assert.Nil(t, err)
	})
}

func Test_GetByID_Category(t *testing.T) {
	t.Run("Should returns error when ID has not been found", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("GetByID", c.ID).Return(&entity.Category{} ,entity.ErrNotFound)
	
		service := NewCategoryService(&repo)
		_, err := service.GetByID(c.ID)
		assert.NotNil(t, err)
	})

	t.Run("Should be ok when when ID has been found", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("GetByID", c.ID).Return(c, nil)
	
		service := NewCategoryService(&repo)
		resp, err := service.GetByID(c.ID)
		assert.Nil(t, err)
		assert.Equal(t, c, resp)
	})
}

func Test_UpdateTitle_Category(t *testing.T) {
	t.Run("Should returns error when repo return and error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("UpdateTitle", c).Return(entity.ErrNotFound)
	
		service := NewCategoryService(&repo)
		err := service.UpdateTitle(c)
		assert.NotNil(t, err)
	})

	t.Run("Should be ok when returns no error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("UpdateTitle", c).Return(nil)
	
		service := NewCategoryService(&repo)
		err := service.UpdateTitle(c)
		assert.Nil(t, err)
	})
}

func Test_UpdateParent_Category(t *testing.T) {
	t.Run("Should returns error when repo return and error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("UpdateParent", c).Return(entity.ErrNotFound)
	
		service := NewCategoryService(&repo)
		err := service.UpdateParent(c)
		assert.NotNil(t, err)
	})

	t.Run("Should be ok when returns no error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("UpdateParent", c).Return(nil)
	
		service := NewCategoryService(&repo)
		err := service.UpdateParent(c)
		assert.Nil(t, err)
	})
}

func Test_ClearParent_Category(t *testing.T) {
	t.Run("Should returns error when repo return and error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("ClearParent", c).Return(entity.ErrNotFound)
	
		service := NewCategoryService(&repo)
		err := service.ClearParent(c)
		assert.NotNil(t, err)
	})

	t.Run("Should be ok when returns no error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("ClearParent", c).Return(nil)
	
		service := NewCategoryService(&repo)
		err := service.ClearParent(c)
		assert.Nil(t, err)
	})
}

func Test_Delete_Category(t *testing.T) {
	t.Run("Should returns error when repo return and error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("Delete", c).Return(entity.ErrNotFound)
	
		service := NewCategoryService(&repo)
		err := service.Delete(c)
		assert.NotNil(t, err)
	})

	t.Run("Should be ok when returns no error", func(t *testing.T) {
		c := mocks.GenerateFixture().CategoryRoot

		repo := category.MockCategoryRepo{}
		repo.On("Delete", c).Return(nil)
	
		service := NewCategoryService(&repo)
		err := service.Delete(c)
		assert.Nil(t, err)
	})
}