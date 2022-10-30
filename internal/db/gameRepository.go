package db

import (
	"context"
	"github.com/d-chosen-one/tgb/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct {
	collection *mongo.Collection
}

func New() GameRepository {
	collection := ClientMongo.Database("tgb").Collection("game")
	instance := new(repo)
	instance.collection = collection
	return instance
}

func (r repo) GetAllGames() ([]*model.Game, error) {
	findOptions := options.Find()

	cur, err := r.collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	var games []*model.Game
	if err = cur.All(context.TODO(), &games); err != nil {
		return nil, err
	}
	return games, nil
}

func (r repo) SaveGame(game model.Game) error {
	_, err := r.collection.InsertOne(context.TODO(), game)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) LoadGame(id string) (*model.Game, error) {
	//TODO implement me
	panic("implement me")
}
