package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xduck7/web-service/internal/entity"
	"github.com/xduck7/web-service/internal/service"
)

type InventionController interface {
	Save(ctx *gin.Context) (entity.Invention, error)
	GetAll() ([]entity.Invention, error)
	GetById(idx int) (entity.Invention, error)
}

type controller struct {
	service service.InventionService
}

func New(service service.InventionService) InventionController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) (entity.Invention, error) {
	var invention entity.Invention
	ctx.BindJSON(&invention)
	c.service.Add(invention)
	return invention, nil
}

func (c *controller) GetAll() ([]entity.Invention, error) {
	return c.service.GetAll()
}

func (c *controller) GetById(idx int) (entity.Invention, error) {
	return c.service.GetById(idx)
}
