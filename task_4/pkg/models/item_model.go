package models

type Item struct {
	ID          string  `bson:"_id" json:"id"`
	Name        string  `bson:"name" json:"name" validate:"required,min=3,max=50"`
	Description string  `bson:"description" json:"description" validate:"required,min=10,max=200"`
	Price       float64 `bson:"price" json:"price" validate:"required,number,gt=0"`
}
