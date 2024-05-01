package service

const (
	authorizedUsername = "xduck7"
	authorizedPassword = "admin"
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
		authorizedUsername: authorizedUsername,
		authorizedPassword: authorizedPassword,
	}
}

func (s *loginService) Login(username string, password string) bool {
	return s.authorizedUsername == username && s.authorizedPassword == password
}
