package repository

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
)

type Repo interface {
	CreateAccount(ctx context.Context, newAccountPtr *user.AccountIfo) error
	DeleteAccount(ctx context.Context, id string) error
	LogIn(ctx context.Context, id string, password string) error
	LogOut(ctx context.Context, id string) error
	UpdateAccount(ctx context.Context, newAccountPtr *user.AccountIfo) error
}
