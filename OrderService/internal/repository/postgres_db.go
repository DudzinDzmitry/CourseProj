package repository

import (
	"context"
	"fmt"
	"github.com/DudzinDzmitry/CourseProj/OrderService/internal/order"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// PRepository p
type PRepository struct {
	Pool *pgxpool.Pool
}

func (p *PRepository) CreateItem(ctx context.Context, med *order.Item) (string, error) {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into medicines(id,name,count,price) values($1,$2,$3,$4)",
		newID, &med.Name, &med.Count, &med.Price)
	if err != nil {
		log.Errorf("database error with create medicine: %v", err)
		return "", err
	}
	return newID, nil
}

func (p *PRepository) GetItemByID(ctx context.Context, idMedicine string) (*order.Item, error) {
	u := order.Item{}
	err := p.Pool.QueryRow(ctx, "select id,name,count,price from medicine where id=$1", idMedicine).Scan(
		&u.Id, &u.Name, &u.Count, &u.Price)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &order.Item{}, fmt.Errorf("medicine with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return &order.Item{}, err
	}
	return &u, nil
}

func (p *PRepository) GetAllItems(ctx context.Context) ([]*order.Item, error) {
	var medicines []*order.Item
	rows, err := p.Pool.Query(ctx, "select id,name,count,price from medicines")
	if err != nil {
		log.Errorf("database error with select all medicines, %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		med := order.Item{}
		err = rows.Scan(&med.Id, &med.Name, &med.Count, &med.Price)
		if err != nil {
			log.Errorf("database error with select all medicines, %v", err)
			return nil, err
		}
		medicines = append(medicines, &med)
	}

	return medicines, nil
}

func (p *PRepository) DeleteItem(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from medicines where id=$1", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("medicine with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("medicine with this id doesnt exist: %v", err)
		}
		log.Errorf("error with delete medicine %v", err)
		return err
	}
	return nil
}

func (p *PRepository) UpdateItem(ctx context.Context, id string, med *order.Item) error {
	a, err := p.Pool.Exec(ctx, "update medicines set name=$1,count=$2,price=$3 where id=$4", &med.Name, &med.Count, &med.Price, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("medicine with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update mediciner %v", err)
		return err
	}
	return nil
}
