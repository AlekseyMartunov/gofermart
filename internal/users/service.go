package users

import (
	"context"
)

type storage interface {
	Create(ctx context.Context, login, password string) (uuid string, err error)
	CheckUser(ctx context.Context, login, password string) (uuid string, err error)
	GetIDByUUID(ctx context.Context, uuid string) (int, error)
	Balance(ctx context.Context, userID int) (User, error)
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

func (us *UserService) Create(ctx context.Context, user User) (uuid string, err error) {
	return us.repo.Create(ctx, user.Login, us.hash.Encode(user.Password))
}

func (us *UserService) CheckUser(ctx context.Context, user User) (uuid string, err error) {
	return us.repo.CheckUser(ctx, user.Login, us.hash.Encode(user.Password))
}

func (us *UserService) GetIDByUUID(ctx context.Context, uuid string) (int, error) {
	return us.repo.GetIDByUUID(ctx, uuid)
}

func (us *UserService) Balance(ctx context.Context, userID int) (User, error) {
	return us.repo.Balance(ctx, userID)
}
