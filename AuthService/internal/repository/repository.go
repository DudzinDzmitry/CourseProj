package repository

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
)

type Repo interface {
	CreateUser(ctx context.Context, p *user.User) error
	UpdateUser(ctx context.Context, id string, p *user.User) error
	DeleteUser(ctx context.Context, id string) error
}
