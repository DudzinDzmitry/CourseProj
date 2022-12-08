package service

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/repository"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
)

// Service s
type Service struct {
	jwtKey []byte
	repo   repository.Repo
}

func NewServiceConnect(pool repository.Repo, jwtKey []byte) *Service {
	return &Service{repo: pool, jwtKey: jwtKey}
}

func (se *Service) CreateAccount(ctx context.Context, newAccountPtr *user.AccountIfo) error {
	return se.repo.CreateAccount(ctx, newAccountPtr)
}

func (se *Service) DeleteAccount(ctx context.Context, id string) error {
	return se.repo.DeleteAccount(ctx, id)
}

func (se *Service) LogIn(ctx context.Context, id, password string) error {
	return se.repo.LogIn(ctx, id, password)
}

func (se *Service) LogOut(ctx context.Context, id, password string) error {
	return se.repo.LogOut(ctx, id)
}

func (se *Service) UpdateAccount(ctx context.Context, id string, user *user.AccountIfo) error {
	return se.repo.UpdateAccount(ctx, id, user)
}
