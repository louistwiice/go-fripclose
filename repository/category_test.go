package repository

import (
	"testing"

	"github.com/louistwiice/go/fripclose/ent"
	"github.com/louistwiice/go/fripclose/ent/enttest"
	"github.com/louistwiice/go/fripclose/ent/migrate"
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Category(t *testing.T) {
	t.Run("Create a Category with non existent root should fail", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat := mocks.GenerateFixture().CategoryTshirt
	
		err :=repo.Create(cat)
	
		assert.NotNil(t, err)
	})

	t.Run("Create a root category should be successful", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat := mocks.GenerateFixture().CategoryRoot
	
		err :=repo.Create(cat)
	
		assert.Nil(t, err)
	})

	t.Run("Create a root category and children should be successful", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat_root := mocks.GenerateFixture().CategoryRoot
		cat_shoes := mocks.GenerateFixture().CategoryShoes
	
		err :=repo.Create(cat_root)
		assert.Nil(t, err)

		err =repo.Create(cat_shoes)
		assert.Nil(t, err)
	})

}

func Test_List_Category(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()

	repo := NewCategoryClient(client)
	cat := mocks.GenerateFixture().CategoryRoot

	resp_list, err := repo.List()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(resp_list))

	_ = repo.Create(cat)

	resp_list, err = repo.List()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp_list))
}

func Test_GetByID_Category(t *testing.T) {
	t.Run("Searching for an non existing ID should return not found", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		
		resp, err := repo.GetByID(3)
		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, resp)
	})

	t.Run("Searching an exsiting ID should passed", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat := mocks.GenerateFixture().CategoryRoot

		_ = repo.Create(cat)

		resp, err := repo.GetByID(cat.ID)
		assert.Nil(t, err)
		assert.Equal(t, cat, resp)
	})
}

func Test_UpdateTitle_Category(t *testing.T) {
	t.Run("Update an non existing Category should not passed", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat := mocks.GenerateFixture().CategoryRoot

		err := repo.UpdateTitle(cat)
		assert.NotNil(t, err)		
	})

	t.Run("Update an existing ID should passed", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat := mocks.GenerateFixture().CategoryRoot

		_ = repo.Create(cat)
		old_title := cat.Title
		cat.Title = "The man"

		err := repo.UpdateTitle(cat)
		assert.Nil(t, err)
		assert.NotEqual(t, old_title, cat.Title)
	})
}

func Test_UpdateParent_Category(t *testing.T) {
	t.Run("Update an non existing Category should not passed", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat := mocks.GenerateFixture().CategoryRoot

		err := repo.UpdateParent(cat)
		assert.NotNil(t, err)		
	})

	t.Run("Set parent of an existing category to a non existing one should not passed", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat_root := mocks.GenerateFixture().CategoryRoot
		cat_shoes := mocks.GenerateFixture().CategoryShoes

		_ = repo.Create(cat_root)
		_ = repo.Create(cat_shoes)

		cat_shoes.ParentID = 13
		err := repo.UpdateParent(cat_shoes)
		assert.NotNil(t, err)		
	})

	t.Run("Set parent of an existing category to an existing one should passed", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat_root := mocks.GenerateFixture().CategoryRoot
		cat_shoes := mocks.GenerateFixture().CategoryShoes
		cat_clothes := mocks.GenerateFixture().CategoryClothes

		_ = repo.Create(cat_root)
		_ = repo.Create(cat_shoes)
		_ = repo.Create(cat_clothes)

		cat_clothes.ParentID = cat_shoes.ID
		err := repo.UpdateParent(cat_clothes)
		assert.Nil(t, err)
	})
}

func Test_ClearParent_Category(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()

	repo := NewCategoryClient(client)
	cat_root := mocks.GenerateFixture().CategoryRoot
	cat_shoes := mocks.GenerateFixture().CategoryShoes

	_ = repo.Create(cat_root)
	_ = repo.Create(cat_shoes)

	err := repo.ClearParent(cat_shoes)
	assert.Nil(t, err)
	assert.Equal(t, 0, cat_shoes.ParentID)
}

func Test_Delete_Category(t *testing.T) {
	t.Run("Delete a non existing category should not pass", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat_root := mocks.GenerateFixture().CategoryRoot

		err := repo.Delete(cat_root)
		assert.NotNil(t, err)

	})

	t.Run("Delete root should make first children become root", func(t *testing.T) {
		opts := []enttest.Option{
			enttest.WithOptions(ent.Log(t.Log)),
			enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
		}
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
		defer client.Close()
	
		repo := NewCategoryClient(client)
		cat_root := mocks.GenerateFixture().CategoryRoot
		cat_shoes := mocks.GenerateFixture().CategoryShoes
		cat_clothes := mocks.GenerateFixture().CategoryClothes

		_ = repo.Create(cat_root)
		_ = repo.Create(cat_shoes)
		_ = repo.Create(cat_clothes)

		assert.Equal(t, 0,cat_root.ParentID)
		assert.Equal(t, cat_root.ID, cat_clothes.ParentID)
		assert.Equal(t, cat_root.ID, cat_shoes.ParentID)

		err := repo.Delete(cat_root)
		resp_cat_shoes, _ := repo.GetByID(cat_shoes.ID)
		resp_cat_clothes, _ := repo.GetByID(cat_clothes.ID)
		assert.Nil(t, err)
		assert.Equal(t, 0, resp_cat_shoes.ParentID)
		assert.Equal(t, 0, resp_cat_clothes.ParentID)
	})
}