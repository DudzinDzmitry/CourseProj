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

// CreateUser _
func (se *Service) CreateUser(ctx context.Context, user *user.User) error {
	return se.repo.CreateUser(ctx, user)
}

// UpdateUser _
func (se *Service) UpdateUser(ctx context.Context, id string, user *user.User) error {
	return se.repo.UpdateUser(ctx, id, user)
}

// DeleteUser _
func (se *Service) DeleteUser(ctx context.Context, id string) error {
	return se.repo.DeleteUser(ctx, id)
}
