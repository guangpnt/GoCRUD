package services

import (
	"context"
	"crud/dao"
	"crud/database/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sv ServicesCRUD) EditAnimal(ctx context.Context, d *dao.GetAnimalRes) error {
	id, err := primitive.ObjectIDFromHex(d.ID)
	ent := &entity.AnimalEntity{
		ID:      id,
		Species: d.Species,
	}
	_, err = sv.Repo.EditAnimal(ctx, ent)
	if err != nil {
		return err
	}
	return err
}
