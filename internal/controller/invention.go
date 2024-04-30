package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/xduck7/mini-go-server/docs"
	"github.com/xduck7/mini-go-server/internal/entity"
	"github.com/xduck7/mini-go-server/internal/service"
	"github.com/xduck7/mini-go-server/internal/validators"
	"net/http"
	"strconv"
)

var validate *validator.Validate

// InventionController представляет интерфейс для управления изобретениями.
type InventionController interface {
	Add(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

// @Summary Add
// @Description Add a new invention with the input payload
// @Tags api/v1
// @Accept  json
// @Produce  json
// @Param invention body entity.Invention true "Invention entity"
// @Success 200 {object} string "Data is valid"
// @Router /invention [post]
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

// @Summary Get
// @Description Get all inventions
// @Tags api/v1
// @Produce  json
// @Success 200 {array} entity.Invention
// @Router /invention [get]
func (c *controller) GetAll() ([]entity.Invention, error) {
	return c.service.GetAll()
}

// @Summary Get
// @Description Get an invention by id
// @Tags api/v1
// @Produce  json
// @Param id path int true "Invention ID"
// @Success 200 {object} entity.Invention
// @Router /invention/{id} [get]
func (c *controller) GetById(idx int) (entity.Invention, error) {
	return c.service.GetById(idx)
}

// @Summary Update
// @Description Update an existing invention with the input payload
// @Tags api/v1
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the invention to update"
// @Param invention body entity.Invention true "Invention entity"
// @Success 200 {object} string "Invention updated successfully"
// @Router /invention/{id} [put]
func (c *controller) Update(ctx *gin.Context) error {
	var invention entity.Invention
	invention.Date.Format("02-01-2006 15:04:05")
	err := ctx.ShouldBindJSON(&invention)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	invention.ID = strconv.FormatUint(id, 10)
	err = validate.Struct(invention)
	if err != nil {
		return err
	}
	c.service.Update(invention)
	return nil
}

// @Summary Delete
// @Description Delete an existing invention by ID
// @Tags api/v1
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the invention to delete"
// @Success 200 {object} string "Invention deleted successfully"
// @Router /invention/{id} [delete]
func (c *controller) Delete(ctx *gin.Context) error {
	var invention entity.Invention
	invention.Date.Format("02-01-2006 15:04:05")
	err := ctx.ShouldBindJSON(&invention)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	invention.ID = strconv.FormatUint(id, 10)
	err = validate.Struct(invention)
	if err != nil {
		return err
	}
	c.service.Delete(invention)
	return nil
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
