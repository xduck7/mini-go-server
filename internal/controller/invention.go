package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xduck7/mini-go-server/internal/entity"
	"github.com/xduck7/mini-go-server/internal/service"
	"github.com/xduck7/mini-go-server/internal/validators"
	"net/http"
)

var validate *validator.Validate

type InventionController interface {
	Add(ctx *gin.Context) error
	GetAll() ([]entity.Invention, error)
	GetById(idx int) (entity.Invention, error)
	ShowAll(ctx *gin.Context)
	ShowMenu(ctx *gin.Context)
}

type controller struct {
	service service.InventionService
}

func New(service service.InventionService) InventionController {
	validate = validator.New()
	if err := validate.RegisterValidation("is_ok", validators.ValidateGoodTitle); err != nil {
		fmt.Println(err)
		return nil
	}
	return &controller{
		service: service,
	}
}

// @Summary Controller
// @Router /api/v1/invention

func (c *controller) Add(ctx *gin.Context) error {
	var invention entity.Invention
	invention.Date.Format("02-01-2006 15:04:05")
	err := ctx.ShouldBindJSON(&invention)
	if err != nil {
		return err
	}
	err = validate.Struct(invention)
	if err != nil {
		return err
	}
	_, err = c.service.Add(invention)
	if err != nil {
		return err
	}
	return nil
}

func (c *controller) GetAll() ([]entity.Invention, error) {
	return c.service.GetAll()
}

func (c *controller) GetById(idx int) (entity.Invention, error) {
	return c.service.GetById(idx)
}

func (c *controller) ShowAll(ctx *gin.Context) {
	inventions, _ := c.service.GetAll()
	data := gin.H{
		"title":      "Inventions page",
		"inventions": inventions,
	}
	ctx.HTML(http.StatusOK, "view.html", data)
}

func (c *controller) ShowMenu(ctx *gin.Context) {
	data := gin.H{
		"title": "Inventions Menu",
	}
	ctx.HTML(http.StatusOK, "menu.html", data)
}
