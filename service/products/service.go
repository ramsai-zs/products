package products

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"products/models"
	"products/store"
	"strconv"
)

type service struct {
	variantStore store.Variants
	store        store.Products
}

func New(s store.Products, vs store.Variants) service {
	return service{store: s, variantStore: vs}
}

func (s service) Create(ctx *gofr.Context, product models.Product) (models.Product, error) {
	err := validateProductBody(product)
	if err != nil {
		return models.Product{}, err
	}

	id, err := s.store.Create(ctx, product)
	if err != nil {
		return models.Product{}, err
	}

	resp, err := s.store.GetByID(ctx, id)
	if err != nil {
		return models.Product{}, err
	}

	return resp, nil
}

func (s service) GetAll(ctx *gofr.Context, params map[string]string) ([]models.Products, error) {
	filters, err := validateFilters(params)
	if err != nil {
		return nil, err
	}

	resp, err := s.store.GetAll(ctx, filters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s service) GetByID(ctx *gofr.Context, id int) (models.Response, error) {
	variant, err := s.variantStore.GetByProductId(ctx, id)
	if err != nil {
		return models.Response{}, err
	}

	product, err := s.store.GetByID(ctx, id)
	if err != nil {
		return models.Response{}, err
	}

	return models.Response{Product: product, Variant: variant}, nil
}

func validateProductBody(p models.Product) error {
	var missing []string

	if p.Name == "" {
		missing = append(missing, "name")
	}

	if p.BrandName == "" {
		missing = append(missing, "brandName")
	}

	if p.Details == "" {
		missing = append(missing, "details")
	}

	if p.ImageUrl == "" {
		missing = append(missing, "imageUrl")
	}

	if len(missing) > 0 {
		return errors.MissingParam{Param: missing}
	}

	return nil
}

func validateFilters(params map[string]string) (filters models.Filters, err error) {
	pId := params["productId"]
	vId := params["variantId"]

	if pId != "" {
		filters.ProductID, err = strconv.Atoi(pId)
		if err != nil {
			return models.Filters{}, errors.InvalidParam{Param: []string{"productId"}}
		}
	}

	if vId != "" {
		filters.VariantID, err = strconv.Atoi(vId)
		if err != nil {
			return models.Filters{}, errors.InvalidParam{Param: []string{"variantId"}}
		}
	}

	filters.ProductName = params["productName"]

	return filters, nil
}
