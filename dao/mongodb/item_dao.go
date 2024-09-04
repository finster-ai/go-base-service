package mongodb

import (
	"context"
	"time"

	"go-base-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_ITEMS = "items"

//TODO: PUEDE QUE EN EL FUTURO SE NECESITE UNA COLECCION ASI, O TENGA QUE USAR 2 TABLAS DESDE EL MISMO ItemDAO
//const COLLECTION_USERID_BY_ITEM_ID = "items"

type ItemDAO struct {
	collection *mongo.Collection
}

func NewItemDAO() *ItemDAO {
	collection := GetCollection(COLLECTION_ITEMS)
	return &ItemDAO{collection: collection}
}

func (dao *ItemDAO) Insert(item model.Item) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := dao.collection.InsertOne(ctx, item)
	return result, err
}

func (dao *ItemDAO) FindByID(id string) (*model.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	var item model.Item
	err := dao.collection.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (dao *ItemDAO) DeleteByID(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	result, err := dao.collection.DeleteOne(ctx, filter)
	return result, err
}
