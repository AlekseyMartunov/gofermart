package loginhandlers

import (
	"context"

	"AlekseyMartunov/internal/users"
)

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type tokenManager interface {
	CreateToken(uuid string) (string, error)
}

type UserService interface {
	Create(ctx context.Context, user users.User) (string, error)
	CheckUser(ctx context.Context, user users.User) (string, error)
}

type LoginHandler struct {
	logger       logger
	userService  UserService
	tokenManager tokenManager
}

func NewLoginHandler(l logger, us UserService, tk tokenManager) *LoginHandler {
	return &LoginHandler{
		logger:       l,
		userService:  us,
		tokenManager: tk,
	}
}

type userDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (dto *userDTO) toEntity() users.User {
	return users.User{
		Login:    dto.Login,
		Password: dto.Password,
	}
}
