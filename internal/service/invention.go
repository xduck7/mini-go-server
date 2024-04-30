package service

import (
	"github.com/xduck7/mini-go-server/internal/entity"
	"github.com/xduck7/mini-go-server/internal/repository"
)

type InventionService interface {
	Add(entity.Invention) (entity.Invention, error)
	GetAll() ([]entity.Invention, error)
	GetById(idx int) (entity.Invention, error)
	Update(invention entity.Invention)
	Delete(invention entity.Invention)
}

type inventionService struct {
	inventionRepository repository.InventionRepository
}

func New(repo repository.InventionRepository) InventionService {
	return &inventionService{
		inventionRepository: repo,
	}
}

func (service *inventionService) Add(invention entity.Invention) (entity.Invention, error) {
	service.inventionRepository.Save(invention)
	return invention, nil
}

func (service *inventionService) GetAll() ([]entity.Invention, error) {
	return service.inventionRepository.FindAll(), nil
}

func (service *inventionService) GetById(idx int) (entity.Invention, error) {
	return service.inventionRepository.FindById(idx), nil
}

func (service *inventionService) Update(invention entity.Invention) {
	service.inventionRepository.Update(invention)
}
func (service *inventionService) Delete(invention entity.Invention) {
	service.inventionRepository.Delete(invention)
}
