package store

import (
	"gofr.dev/pkg/gofr"

	"products/models"
)

type Products interface {
	Create(ctx *gofr.Context, p models.Product) (int, error)
	GetByID(ctx *gofr.Context, id int) (models.Product, error)
	GetAll(ctx *gofr.Context, filters models.Filters) ([]models.Products, error)
}

type Variants interface {
	Create(ctx *gofr.Context, v models.Variant) (int, error)
	GetByID(ctx *gofr.Context, id int) (models.Variant, error)
	GetByProductId(ctx *gofr.Context, pId int) ([]models.Variant, error)
	GetByIdAndProductId(ctx *gofr.Context, id, pID int) (models.Variant, error)
	GetByMultipleProductId(ctx *gofr.Context, pIDs []string) ([]models.Variant, error)
}
