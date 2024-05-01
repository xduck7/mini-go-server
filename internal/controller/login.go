package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xduck7/mini-go-server/internal/entity"
	"github.com/xduck7/mini-go-server/internal/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func NewLoginController(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var auth entity.Auth
	err := ctx.ShouldBind(&auth)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.loginService.Login(auth.Username, auth.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(auth.Username, true)
	}
	return ""
}
