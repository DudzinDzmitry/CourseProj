package repository

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/OrderService/internal/order"
)

type Repository interface {
	CreateItem(ctx context.Context, p *order.Item) (string, error)
	GetItemByID(ctx context.Context, idPerson string) (*order.Item, error)
	GetAllItems(ctx context.Context) ([]*order.Item, error)
	DeleteItem(ctx context.Context, id string) error
	UpdateItem(ctx context.Context, id string, per *order.Item) error
}
