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

// CreateUser add user to db
func (p *PRepository) CreateUser(ctx context.Context, person *user.Person) (string, error) {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into persons(id,name,position,password) values($1,$2,$3,$4)",
		newID, &person.Name, &person.Position, &person.Password)
	if err != nil {
		log.Errorf("database error with create user: %v", err)
		return "", err
	}
	return newID, nil
}

// DeleteUser delete user by id
func (p *PRepository) DeleteUser(ctx context.Context, id string) error {
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

// UpdateUser update parameters for user
func (p *PRepository) UpdateUser(ctx context.Context, id string, per *user.Person) error {
	a, err := p.Pool.Exec(ctx, "update persons set name=$1,position=$2 where id=$3", &per.Name, &per.Position, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}

// UpdateAuth logout, delete refresh token
func (p *PRepository) UpdateAuth(ctx context.Context, id, refreshToken string) error {
	a, err := p.Pool.Exec(ctx, "update persons set refreshToken=$1 where id=$2", refreshToken, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}
