package users

import (
	"context"
)

type storage interface {
	Register(ctx context.Context, login, password string) error
}

type hash interface {
	Encode(key string) string
}

type UserService struct {
	repo storage
	hash hash
}

func NewUserService(r storage, h hash) *UserService {
	return &UserService{
		repo: r,
		hash: h,
	}
}

func (us *UserService) Register(ctx context.Context, login, password string) error {
	return us.repo.Register(ctx, login, us.hash.Encode(password))
}
