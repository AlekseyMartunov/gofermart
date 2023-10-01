package users

import "context"

type storage interface {
	Register(ctx context.Context, login, password string) error
}

type UserService struct {
	repo storage
}

func NewUserService(repo storage) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, login, password string) error {
	return us.repo.Register(ctx, login, password)
}
