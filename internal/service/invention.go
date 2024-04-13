package service

import (
	"github.com/xduck7/web-service/internal/entity"
)

type InventionService interface {
	Add(entity.Invention) (entity.Invention, error) //сохраняем новое изобретение
	GetAll() ([]entity.Invention, error)            // возвращаем все
	GetById(idx int) (entity.Invention, error)      //возвращаем изобретение по id
}

type inventionService struct {
	inventionsList []entity.Invention
}

func New() InventionService {
	return &inventionService{}
}

func (service *inventionService) Add(invention entity.Invention) (entity.Invention, error) {
	service.inventionsList = append(service.inventionsList, invention)
	return invention, nil
}

func (service *inventionService) GetAll() ([]entity.Invention, error) {
	return service.inventionsList, nil
}

// todo: fix
func (service *inventionService) GetById(idx int) (entity.Invention, error) {
	lst := service.inventionsList
	for i := 0; i < len(lst); i++ {
		if lst[i].ID == string(idx) {
			return lst[i], nil
		}
	}
	return entity.Invention{}, nil
}
