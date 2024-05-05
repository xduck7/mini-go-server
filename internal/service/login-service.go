package service

const (
	authorizedUsernameForJWT = "xduck7"
	authorizedPasswordForJWT = "admin"
)

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: authorizedUsernameForJWT,
		authorizedPassword: authorizedPasswordForJWT,
	}
}

func (s *loginService) Login(username string, password string) bool {
	return s.authorizedUsername == username && s.authorizedPassword == password
}
