package variants

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"products/models"
	"products/store"
)

type service struct {
	productStore store.Products
	store        store.Variants
}

func New(ps store.Products, s store.Variants) service {
	return service{productStore: ps, store: s}
}

func (s service) Create(ctx *gofr.Context, variant models.Variant) (models.Variant, error) {
	pID, err := strconv.Atoi(variant.ProductID)
	if err != nil || pID <= 0 {
		return models.Variant{}, errors.InvalidParam{Param: []string{"productId"}}
	}

	_, err = s.productStore.GetByID(ctx, pID)
	if err != nil {
		return models.Variant{}, err
	}

	err = validateVariantBody(variant)
	if err != nil {
		return models.Variant{}, err
	}

	id, err := s.store.Create(ctx, variant)
	if err != nil {
		return models.Variant{}, err
	}

	resp, err := s.store.GetByID(ctx, id)
	if err != nil {
		return models.Variant{}, err
	}

	return resp, nil
}

func (s service) GetByIdAndProductId(ctx *gofr.Context, variantId, productId int) (models.Variant, error) {
	resp, err := s.store.GetByIdAndProductId(ctx, variantId, productId)
	if err != nil {
		return models.Variant{}, err
	}

	return resp, nil
}

func validateVariantBody(v models.Variant) error {
	var missing []string

	if v.Name == "" {
		missing = append(missing, "name")
	}

	if v.Details == "" {
		missing = append(missing, "details")
	}

	if len(missing) > 0 {
		return errors.MissingParam{Param: missing}
	}

	return nil
}
