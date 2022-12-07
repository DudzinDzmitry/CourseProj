package repository

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/PhCatalog/catalog"
)

type Repository interface {
	CreateMedicine(ctx context.Context, p *catalog.Medicine) (string, error)
	GetMedicineByID(ctx context.Context, idPerson string) (*catalog.Medicine, error)
	GetAllMedicine(ctx context.Context) ([]*catalog.Medicine, error)
	DeleteMedicine(ctx context.Context, id string) error
	ChangeMedicine(ctx context.Context, id string, per *catalog.Medicine) error
}
