package services

import (
	"context"
	"crud/dao"
	"crud/database/entity"
)

func (sv ServicesCRUD) AddAnimal(ctx context.Context, d *dao.CreateAnimalReq) (*dao.CreateAnimalRes, error) {
	ent := &entity.AnimalEntity{
		Species:  d.Species,
	}
	_, err := sv.Repo.InsertAnimal(ctx, ent)
	if err != nil {
		return &dao.CreateAnimalRes{}, err
	}
	return &dao.CreateAnimalRes{}, err
}
