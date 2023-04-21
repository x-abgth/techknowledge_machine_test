package repository

import (
	"errors"

	"example.com/machine_test/task_4/pkg/models"
)

type Respository struct {
	Items []models.Item
}

func (rep *Respository) InsertItemRepository(item models.Item) error {
	rep.Items = append(rep.Items, item)

	return nil
}

func (rep *Respository) GetOneItemRepository(id string) (models.Item, error) {
	var item models.Item

	for _, val := range rep.Items {
		if val.ID == id {
			return item, nil
		}
	}

	return item, errors.New("Couldn't find the data")
}

func (rep *Respository) GetAllItemsRespository() ([]models.Item, error) {
	return rep.Items, nil
}

func (rep *Respository) UpdateOneItemRepository(item models.Item) error {
	for _, val := range rep.Items {
		if val.ID == item.ID {
			val.Name = item.Name
			val.Description = item.Description
			val.Price = item.Price

			return nil
		}
	}

	return errors.New("Couldn't find the data")
}

func (rep *Respository) DeleteOneItemRespository(id string) error {
	for i, val := range rep.Items {
		if val.ID == id {
			rep.Items = append(rep.Items[:i], rep.Items[i+1:]...)
			return nil
		}
	}

	return errors.New("Couldn't find the data")
}
