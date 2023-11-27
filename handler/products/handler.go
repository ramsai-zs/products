package products

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"products/models"
	"products/service"
)

type handler struct {
	service service.Products
}

func New(svc service.Products) handler {
	return handler{service: svc}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var p models.Product

	err := ctx.Bind(&p)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.service.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	pId, err := strconv.Atoi(id)
	if err != nil || pId <= 0 {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.service.GetByID(ctx, pId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	params := ctx.Params()

	resp, err := h.service.GetAll(ctx, params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
