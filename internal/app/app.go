package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xduck7/web-service/internal/controller"
	"github.com/xduck7/web-service/internal/middleware"
	"github.com/xduck7/web-service/internal/service"
	"net/http"
	"strconv"
)

var (
	inventionService    service.InventionService       = service.New()
	inventionController controller.InventionController = controller.New(inventionService)
)

func Run() {

	middleware.SetupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "200",
		})
	})

	server.POST("/invention", func(ctx *gin.Context) {
		invList, err := inventionController.Save(ctx)
		if err != nil {
			ctx.JSON(0, nil)
		}
		ctx.JSON(200, invList)
	})

	server.GET("/invention", func(ctx *gin.Context) {
		invList, err := inventionController.GetAll()
		if err != nil {
			ctx.JSON(404, nil)
		}
		ctx.JSON(200, invList)
	})

	server.GET("/invention/:id", func(ctx *gin.Context) {
		// извлекаем id из пути и конвертируем его в int
		idAddr := ctx.Param("id")
		id, err := strconv.Atoi(idAddr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
			return
		}
		inv, err := inventionController.GetById(id)
		if err != nil {
			ctx.JSON(404, nil)
		}
		ctx.JSON(200, inv)
	})

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
