package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xduck7/mini-go-server/internal/controller"
	"github.com/xduck7/mini-go-server/internal/middleware"
	"github.com/xduck7/mini-go-server/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	inventionService    = service.New()
	inventionController = controller.New(inventionService)
)

func Run() {
	middleware.SetupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	templatesPath := filepath.Join(currentDir, "/internal/templates")
	server.Static("/css", templatesPath+"/css")
	server.LoadHTMLGlob(templatesPath + "/*.html")

	apiRoutes := server.Group("/api/v1")
	{
		apiRoutes.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "200",
			})
		})

		apiRoutes.POST("/invention", func(ctx *gin.Context) {
			err := inventionController.Add(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"error": "Data is valid"})
			}
		})

		apiRoutes.GET("/invention", func(ctx *gin.Context) {
			invList, err := inventionController.GetAll()
			if err != nil {
				ctx.JSON(404, nil)
			}
			ctx.JSON(200, invList)
		})

		apiRoutes.GET("/invention/:id", func(ctx *gin.Context) {
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
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/inventions", inventionController.ShowAll)
		viewRoutes.GET("/menu", inventionController.ShowMenu)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = server.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
