package service

import (
	"gofr.dev/pkg/gofr"

	"products/models"
)

type Products interface {
	Create(ctx *gofr.Context, product models.Product) (models.Product, error)
	GetByID(ctx *gofr.Context, id int) (models.Response, error)
	GetAll(ctx *gofr.Context, params map[string]string) ([]models.Products, error)
}

type Variants interface {
	Create(ctx *gofr.Context, variant models.Variant) (models.Variant, error)
	GetByIdAndProductId(ctx *gofr.Context, variantId, productId int) (models.Variant, error)
}
