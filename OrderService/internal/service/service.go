package service

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/OrderService/internal/order"
	"github.com/DudzinDzmitry/CourseProj/OrderService/internal/repository"
)

type Service struct {
	jwtKey []byte
	rps    repository.Repository
}

// NewService create new service connection
func NewService(pool repository.Repository, jwtKey []byte) *Service {
	return &Service{rps: pool, jwtKey: jwtKey}
}

func (se *Service) GetItem(ctx context.Context, id string) (*order.Item, error) {
	return se.rps.GetItemByID(ctx, id)
}

func (se *Service) GetAllItems(ctx context.Context) ([]*order.Item, error) {
	return se.rps.GetAllItems(ctx)
}

// DeleteItem _
func (se *Service) DeleteItem(ctx context.Context, id string) error {
	return se.rps.DeleteItem(ctx, id)
}

// UpdateItem _
func (se *Service) UpdateItem(ctx context.Context, id string, user *order.Item) error {
	return se.rps.UpdateItem(ctx, id, user)
}

// CreateItem create new medicine, add him to db
func (se *Service) CreateItem(ctx context.Context, m *order.Item) (string, error) {
	return se.rps.CreateItem(ctx, m)
}
