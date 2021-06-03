package services

import (
	"context"
	"crud/dao"
)

func (sv ServicesCRUD) GetAllAnimal(ctx context.Context) ([]*dao.GetAnimalRes, error) {
	Res := make([]*dao.GetAnimalRes, 0)
	res, err := sv.Repo.FindAllAnimal(ctx)
	if err != nil {
		return []*dao.GetAnimalRes{}, err
	}
	for _, r := range *res {
		data := &dao.GetAnimalRes{
			ID:       r.ID.Hex(),
			Species:  r.Species,
		}
		Res = append(Res, data)
	}
	return Res, err
}
