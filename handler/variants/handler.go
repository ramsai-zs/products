package variants

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"products/models"
	"products/service"
)

type handler struct {
	service service.Variants
}

func New(svc service.Variants) handler {
	return handler{service: svc}
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	pID := ctx.PathParam("pid")
	vID := ctx.PathParam("id")

	productId, err := validateID(pID, "pid")
	if err != nil {
		return nil, err
	}

	variantId, err := validateID(vID, "id")
	if err != nil {
		return nil, err
	}

	resp, err := h.service.GetByIdAndProductId(ctx, variantId, productId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	pID := ctx.PathParam("pid")

	var variant models.Variant

	err := ctx.Bind(&variant)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	variant.ProductID = pID

	resp, err := h.service.Create(ctx, variant)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// validateID validates the given param and returns param as int,
// if error returns error with the given paramName
func validateID(param, paramName string) (int, error) {
	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		return 0, errors.InvalidParam{Param: []string{paramName}}
	}

	return id, nil
}
