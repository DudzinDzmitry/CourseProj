package repository

import (
	"context"
	"fmt"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// PRepository p
type PRepository struct {
	Pool *pgxpool.Pool
}

func (p *PRepository) CreateAccount(ctx context.Context, newAccountPtr *user.AccountIfo) error {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into persons(id,name,position,password) values($1,$2,$3,$4)",
		newID, &newAccountPtr.UserName, &newAccountPtr.Password)
	if err != nil {
		log.Errorf("database error with create user: %v", err)
		return err
	}
	return nil
}

func (p *PRepository) DeleteAccount(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from persons where id=$1", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("error with delete user %v", err)
		return err
	}
	return nil
}

func (p *PRepository) LogIn(ctx context.Context, id string, password string) error {
	a, err := p.Pool.Exec(ctx, "update persons set refreshToken=$1 where id=$2", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}

func (p *PRepository) LogOut(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PRepository) UpdateAccount(ctx context.Context, newAccountPtr *user.AccountIfo) error {
	a, err := p.Pool.Exec(ctx, "update persons set name=$1,position=$2 where id=$3", &newAccountPtr.UserName)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}
