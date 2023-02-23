package usecase

import (
	"example.com/machine_test/task_4/internal/repository"
	"example.com/machine_test/task_4/pkg/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

var rep repository.Respository

func NewClientInit(client *mongo.Client) {
	rep = repository.Respository{
		Client: client,
	}
}

func GetAllItemsUsecase() ([]models.Item, error) {
	return rep.GetAllItemsRespository()
}

func GetOneItemUsecase(id string) (models.Item, error) {
	return rep.GetOneItemRepository(id)
}

func InsertItemUsecase(item models.Item) error {
	item.ID = uuid.New().String()

	return rep.InsertItemRepository(item)
}

func UpdateItemUsecase(id string, item models.Item) error {
	item.ID = id
	return rep.UpdateOneItemRepository(item)
}

func DeleteItemUsecase(id string) error {
	return rep.DeleteOneItemRespository(id)
}
