package repository_category

import (
	"context"

	"github.com/louistwiice/go/fripclose/ent"
	"github.com/louistwiice/go/fripclose/ent/category"
	"github.com/louistwiice/go/fripclose/entity"
)


type CategoryClient struct {
	client *ent.Client
}

func NewCategoryClient(client *ent.Client) *CategoryClient {
	return &CategoryClient{
		client: client,
	}
}

func (c * CategoryClient) List() ([]*entity.Category, error) {
	var data []*entity.Category
	ctx := context.Background()

	err := c.client.Category.Query().
		Select(category.FieldID, category.FieldTitle, category.FieldParentID).
		Scan(ctx, &data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *CategoryClient) Create(data *entity.Category) error {
	ctx := context.Background()
	var err error
	var resp *ent.Category

	if data.ParentID != 0 {
		resp, err = c.client.Category.Create().
		SetTitle(data.Title).
		SetParentID(data.ParentID).
		Save(ctx)
	} else {
		resp, err = c.client.Category.Create().
		SetTitle(data.Title).
		Save(ctx)
	}

	if err != nil {
		return err
	}

	data.ID = resp.ID
	data.ParentID = resp.ParentID
	return nil
}

func (c *CategoryClient) GetByID(id int) (*entity.Category, error) {
	var data entity.Category
	ctx := context.Background()

	resp := c.client.Category.Query().
		Where(category.ID(id)).
		AllX(ctx)
	
	if len(resp) > 0 {
		data.ID = resp[0].ID
		data.Title = resp[0].Title
		data.ParentID = resp[0].ParentID
	} else {
		return nil, entity.ErrNotFound
	}

	return &data, nil
}

func (c *CategoryClient) UpdateTitle(data *entity.Category) error {
	ctx := context.Background()

	resp, err := c.client.Category.UpdateOneID(data.ID).
		SetTitle(data.Title).
		Save(ctx)

	if err != nil {
		return err
	}

	data.ID = resp.ID
	data.ParentID = resp.ParentID

	return nil
}

func (c *CategoryClient) UpdateParent(data *entity.Category) error {
	ctx := context.Background()

	resp, err := c.client.Category.UpdateOneID(data.ID).
		SetParentID(data.ParentID).
		Save(ctx)

	if err != nil {
		return err
	}

	data.ID = resp.ID
	data.ParentID = resp.ParentID

	return nil
}

func (c *CategoryClient) ClearParent(data *entity.Category) error {
	ctx := context.Background()

	resp, err := c.client.Category.UpdateOneID(data.ID).
		ClearParentID().
		Save(ctx)

	if err != nil {
		return err
	}

	data.ID = resp.ID
	data.ParentID = resp.ParentID

	return nil
}

func (c *CategoryClient) Delete(data *entity.Category) error {
	ctx := context.Background()
	var err error

	if data.ParentID == 0 {
		// If the data to delete has no parent it means, we have to clear parents children
		_, err = c.client.Category.Update().
		ClearParentID().
		Where(category.ParentID(data.ID)).
		Save(ctx)
	} else {
		_, err = c.client.Category.Update().
		SetParentID(data.ParentID).
		Where(category.ParentID(data.ID)).
		Save(ctx)
	}
	if err != nil {
		return err
	}

	err = c.client.Category.DeleteOneID(data.ID).Exec(ctx)
	return err
}