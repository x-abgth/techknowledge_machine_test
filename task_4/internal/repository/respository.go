package repository

import (
	"context"
	"log"
	"time"

	"example.com/machine_test/task_4/pkg/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Respository struct {
	Client *mongo.Client
}

func (rep *Respository) InsertItemRepository(item models.Item) error {
	collection := rep.Client.Database("items").Collection("items")

	_, err := collection.InsertOne(context.TODO(), models.Item{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	})

	if err != nil {
		log.Println("Error inserting", err)
		return err
	}

	return nil
}

func (rep *Respository) GetOneItemRepository(id string) (models.Item, error) {
	var item models.Item

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := rep.Client.Database("items").Collection("items")

	docID, err := primitive.ObjectIDFromHex(uuid.MustParse(id).String())
	if err != nil {
		return item, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (rep *Respository) GetAllItemsRespository() ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := rep.Client.Database("items").Collection("items")

	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []models.Item

	for cursor.Next(ctx) {
		var item models.Item

		err := cursor.Decode(&item)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (rep *Respository) UpdateOneItemRepository(item models.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := rep.Client.Database("items").Collection("items")

	docID, err := primitive.ObjectIDFromHex(string(uuid.MustParse(item.ID).String()))
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{
				"$set", bson.D{
					{"name", item.Name},
					{"description", item.Description},
					{"price", item.Price},
				},
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (rep *Respository) DeleteOneItemRespository(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := rep.Client.Database("items").Collection("items")

	docID, err := primitive.ObjectIDFromHex(string(uuid.MustParse(id).String()))
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		return err
	}

	return nil
}
