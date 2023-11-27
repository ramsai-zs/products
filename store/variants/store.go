package variants

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"products/models"
)

type store struct{}

func New() store {
	return store{}
}

func (s store) Create(ctx *gofr.Context, v models.Variant) (int, error) {
	var id int64

	res, err := ctx.DB().ExecContext(ctx, create, v.ProductID, v.Name, v.Details)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	if id, err = res.LastInsertId(); err != nil {
		return 0, errors.DB{Err: err}
	}

	return int(id), nil
}

func (s store) GetByID(ctx *gofr.Context, id int) (models.Variant, error) {
	var v models.Variant

	res := ctx.DB().QueryRowContext(ctx, getById, id)

	err := res.Scan(&v.ID, &v.ProductID, &v.Name, &v.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Variant{}, errors.EntityNotFound{Entity: "variants", ID: strconv.Itoa(id)}
		}
		return models.Variant{}, errors.DB{Err: err}
	}

	return v, nil
}

func (s store) GetByProductId(ctx *gofr.Context, pId int) ([]models.Variant, error) {
	var variant []models.Variant

	rows, err := ctx.DB().QueryContext(ctx, getByProductId, pId)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	for rows.Next() {
		var v models.Variant

		err := rows.Scan(&v.ID, &v.Name, &v.Details)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		variant = append(variant, v)
	}

	return variant, nil
}

func (s store) GetByIdAndProductId(ctx *gofr.Context, id, pID int) (models.Variant, error) {
	var v models.Variant

	res := ctx.DB().QueryRowContext(ctx, getByIdAndProductId, id, pID)

	err := res.Scan(&v.ID, &v.ProductID, &v.Name, &v.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Variant{}, errors.EntityNotFound{Entity: "variants", ID: fmt.Sprintf("id :%v,productId:%v", id, pID)}
		}
		return models.Variant{}, errors.DB{Err: err}
	}

	return v, nil
}

func (s store) GetByMultipleProductId(ctx *gofr.Context, pIDs []string) ([]models.Variant, error) {
	var variant []models.Variant

	query := getByProductId

	qp, query := getQuery(pIDs)

	rows, err := ctx.DB().QueryContext(ctx, query, qp...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var v models.Variant

		err := rows.Scan(&v.ID, &v.ProductID, &v.Name, &v.Details)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		variant = append(variant, v)
	}

	return variant, nil
}

func getQuery(pIDs []string) (queryParam []interface{}, query string) {
	queryParam = make([]interface{}, 0)

	query = "SELECT id,product_id,variant_name,variant_details FROM varaints where product_id in ("

	ID := make([]string, 0)

	for i := range pIDs {
		ID = append(ID, "?")
		queryParam = append(queryParam, pIDs[i])
	}

	clause := strings.Join(ID, ",")

	query += clause + ")"

	return queryParam, query
}
