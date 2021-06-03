package services

import (
	"context"
	"crud/dao"
	"crud/database/repository"
)

type ServicesCRUD struct {
	Repo repository.IRepository
}

type IServicesCRUD interface {
	GetAllAnimal(ctx context.Context) ([]*dao.GetAnimalRes, error)
	AddAnimal(ctx context.Context, d *dao.CreateAnimalReq) (*dao.CreateAnimalRes, error)
	DeleteAnimal(ctx context.Context, d *dao.DeleteAnimalReq) error
	EditAnimal(ctx context.Context, d *dao.GetAnimalRes) error
}

func NewServiceCRUD(repo repository.IRepository) IServicesCRUD {
	return &ServicesCRUD{
		Repo: repo,
	}
}
