package products

import (
	"database/sql"
	stores "products/store"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"products/models"
)

type store struct {
	variantStore stores.Variants
}

func New(vs stores.Variants) store {
	return store{variantStore: vs}
}

func (s store) Create(ctx *gofr.Context, p models.Product) (int, error) {
	var id int64

	res, err := ctx.DB().ExecContext(ctx, create, p.Name, p.BrandName, p.Details, p.ImageUrl)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	if id, err = res.LastInsertId(); err != nil {
		return 0, errors.DB{Err: err}
	}

	return int(id), nil
}

func (s store) GetByID(ctx *gofr.Context, id int) (models.Product, error) {
	var p models.Product

	res := ctx.DB().QueryRowContext(ctx, getById, id)

	err := res.Scan(&p.ID, &p.Name, &p.BrandName, &p.Details, &p.ImageUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
		}
		return models.Product{}, errors.DB{Err: err}
	}

	return p, nil
}

func (s store) GetAll(ctx *gofr.Context, filters models.Filters) ([]models.Products, error) {
	whereClause, values := generateWhereClause(filters)

	query := "SELECT id, name, brand_name, details, image_url FROM products"

	var products []models.Products

	q := query + whereClause

	rows, err := ctx.DB().QueryContext(ctx, q, values...)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var (
			p   models.Products
			pid int
		)

		err = rows.Scan(&p.ID, &p.Name, &p.BrandName, &p.Details, &p.ImageURL)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		pid, err = strconv.Atoi(p.ID)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		if filters.VariantID < 1 {
			p.Variants, err = s.variantStore.GetByProductId(ctx, pid)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, errors.DB{Err: err}
				}
			}
		} else {
			variant, err := s.variantStore.GetByIdAndProductId(ctx, filters.VariantID, pid)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, errors.DB{Err: err}
				}
			}

			p.Variants = []models.Variant{variant}
		}

		products = append(products, p)
	}

	return products, nil
}

func generateWhereClause(filters models.Filters) (clause string, values []interface{}) {
	clause = ""

	if filters.ProductID != 0 && filters.ProductName != "" {
		clause = " WHERE id=? AND name=?"
		values = append(values, filters.ProductID, filters.ProductName)
	} else if filters.ProductID != 0 {
		clause = " WHERE id=?"
		values = append(values, filters.ProductID)
	} else if filters.ProductName != "" {
		clause = " WHERE name=?"
		values = append(values, filters.ProductName)
	}

	return clause, values
}
