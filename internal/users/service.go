package users

import (
	"context"
)

type storage interface {
	Create(ctx context.Context, login, password string) (uuid string, err error)
	CheckUser(ctx context.Context, login, password string) (uuid string, err error)
	CheckUserUUID(ctx context.Context, userUUID string) error
	GetIDByUUID(ctx context.Context, uuid string) (int, error)
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

func (us *UserService) Create(ctx context.Context, login, password string) (uuid string, err error) {
	return us.repo.Create(ctx, login, us.hash.Encode(password))
}

func (us *UserService) CheckUser(ctx context.Context, login, password string) (uuid string, err error) {
	return us.repo.CheckUser(ctx, login, us.hash.Encode(password))
}

func (us *UserService) CheckUserUUID(ctx context.Context, userUUID string) error {
	return us.repo.CheckUserUUID(ctx, userUUID)
}

func (us *UserService) GetIDByUUID(ctx context.Context, uuid string) (int, error) {
	return us.repo.GetIDByUUID(ctx, uuid)
}
