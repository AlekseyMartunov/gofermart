package users

import (
	"context"
)

type storage interface {
	Register(ctx context.Context, login, password string) error
	CheckUserUUID(ctx context.Context, login, password string) (uuid string, err error)
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

func (us *UserService) CheckUserUUID(ctx context.Context, login, password string) (uuid string, err error) {
	return us.repo.CheckUserUUID(ctx, login, us.hash.Encode(password))
}
