package repository

import (
	"context"
	"crud/database/datasources"
	"crud/database/entity"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Collection *mongo.Collection
}

type IRepository interface {
	FindAllAnimal(ctx context.Context) (*[]entity.AnimalEntity, error)
	InsertAnimal(ctx context.Context, ent *entity.AnimalEntity) (*mongo.InsertOneResult, error)
	DeleteAnimal(ctx context.Context, Species string) *mongo.SingleResult
	EditAnimal(ctx context.Context, ent *entity.AnimalEntity) (*mongo.UpdateResult, error)
}

func NewRepository(ds *datasources.MongoDB) IRepository {
	return &Repository{
		Collection: ds.MongoDB.Database("crudtest").Collection("animals"),
	}
}

func (repo Repository) FindAllAnimal(ctx context.Context) (*[]entity.AnimalEntity, error) {
	option := options.Find()
	cur, err := repo.Collection.Find(ctx, bson.M{}, option)
	if err != nil {
		return &[]entity.AnimalEntity{}, err
	}
	defer cur.Close(ctx)
	var ent []entity.AnimalEntity
	err = cur.All(ctx, &ent)
	return &ent, err
}

func (repo Repository) InsertAnimal(ctx context.Context, ent *entity.AnimalEntity) (*mongo.InsertOneResult, error) {
	return repo.Collection.InsertOne(ctx, &ent)
}

func (repo Repository) DeleteAnimal(ctx context.Context, Species string) *mongo.SingleResult {
	filter := bson.M{"species": Species}
	return repo.Collection.FindOneAndDelete(ctx, filter)
}

func (repo Repository) EditAnimal(ctx context.Context, ent *entity.AnimalEntity) (*mongo.UpdateResult, error) {
	fmt.Println("Entity ID : ", ent.ID)
	fmt.Println("Species : ", ent.Species)
	return repo.Collection.UpdateOne(ctx, bson.M{
		"_id": ent.ID,
	},
		bson.M{"$set": bson.M{
			"species": ent.Species,
		},
		},
	)
}
