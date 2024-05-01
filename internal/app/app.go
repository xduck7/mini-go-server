package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xduck7/mini-go-server/internal/controller"
	"github.com/xduck7/mini-go-server/internal/middleware"
	"github.com/xduck7/mini-go-server/internal/repository"
	"github.com/xduck7/mini-go-server/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	inventionRepo                                  = repository.NewInventionRepository()
	inventionService                               = service.New(inventionRepo)
	loginService        service.LoginService       = service.NewLoginService()
	jwtService          service.JWTService         = service.NewJWTService()
	loginController     controller.LoginController = controller.NewLoginController(loginService, jwtService)
	inventionController                            = controller.New(inventionService)
)

func Run(port string) {
	middleware.SetupLogOutput()

	defer inventionRepo.CloseDB()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	templatesPath := filepath.Join(currentDir, "/internal/templates")
	server.Static("/css", templatesPath+"/css")
	server.LoadHTMLGlob(templatesPath + "/*.html")

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api/v1", middleware.AuthJWT())
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

		apiRoutes.PUT("/invention/:id", func(ctx *gin.Context) {
			err := inventionController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"error": "Data is valid"})
			}
		})

		apiRoutes.DELETE("/invention/:id", func(ctx *gin.Context) {
			err := inventionController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"error": "Data is valid"})
			}
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

	viewRoutes := server.Group("/")
	{
		viewRoutes.GET("/inventions", inventionController.ShowAll)
		viewRoutes.GET("/menu", inventionController.ShowMenu)
		viewRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	err = server.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
