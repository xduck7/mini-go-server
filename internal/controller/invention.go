package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xduck7/web-service/internal/entity"
	"github.com/xduck7/web-service/internal/service"
	"github.com/xduck7/web-service/internal/validators"
)

var validate *validator.Validate

type InventionController interface {
	Add(ctx *gin.Context) error
	GetAll() ([]entity.Invention, error)
	GetById(idx int) (entity.Invention, error)
}

type controller struct {
	service service.InventionService
}

func New(service service.InventionService) InventionController {
	validate = validator.New()
	validate.RegisterValidation("is-good", validators.ValidateGoodTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) Add(ctx *gin.Context) error {
	var invention entity.Invention
	err := ctx.ShouldBindJSON(&invention)
	if err != nil {
		return err
	}
	err = validate.Struct(invention)
	if err != nil {
		return err
	}
	c.service.Add(invention)
	return nil
}

func (c *controller) GetAll() ([]entity.Invention, error) {
	return c.service.GetAll()
}

func (c *controller) GetById(idx int) (entity.Invention, error) {
	return c.service.GetById(idx)
}
