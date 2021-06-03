package services

import (
	"context"
	"crud/dao"
)

func (sv ServicesCRUD) DeleteAnimal(ctx context.Context, d *dao.DeleteAnimalReq) error {
	err := sv.Repo.DeleteAnimal(ctx, d.Species)
	if err.Err() != nil {
		if err.Err().Error() == "mongo: no documents in result" {
			return err.Err()
		}
	}
	return err.Err()
}
