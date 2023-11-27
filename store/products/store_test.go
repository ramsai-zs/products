package products

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"products/models"
	stores "products/store"
	"testing"
)

const (
	testCreate  = "INSERT into products(name,brand_name,details,image_url) VALUES(?,?,?,?)"
	testGetById = "SELECT id,name,brand_name,details,image_url FROM products WHERE id = ?"
	testGet     = "SELECT id,name,brand_name,details,image_url FROM products"
)

// initializeVariantStore initializes mock,db and returns instance of mysql,sql and gofr.
func initializeDataStore(t *testing.T) (store, sqlmock.Sqlmock, *gofr.Context) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Mock is not initialized. Error: %v", err)
	}

	ctx := gofr.NewContext(nil, nil, gofr.New())
	ctx.Context = context.Background()
	ctx.DataStore = datastore.DataStore{ORM: db}

	str := New(&stores.MockVariants{})

	return str, mock, ctx
}

func TestStore_Create(t *testing.T) {
	s, mock, ctx := initializeDataStore(t)

	tests := []struct {
		desc string
		id   int
		mock *sqlmock.ExpectedExec
		err  error
	}{
		{"success", 1, mock.ExpectExec(testCreate).WithArgs("name", "tesla", "details", "image.png").
			WillReturnResult(sqlmock.NewResult(1, 1)), nil},
		{"failure - error while inserting", 0,
			mock.ExpectExec(testCreate).WithArgs("name", "tesla", "details", "image.png").WillReturnError(errors.EntityAlreadyExists{}),
			errors.DB{Err: errors.EntityAlreadyExists{}}},
		{"failure - last insert id 0", 0, mock.ExpectExec(testCreate).WithArgs("name", "tesla", "details", "image.png").
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows)), errors.DB{Err: sql.ErrNoRows}},
	}

	for i, tc := range tests {
		res, err := s.Create(ctx, models.Product{Name: "name", BrandName: "tesla",
			Details: "details", ImageUrl: "image.png"})

		assert.Equal(t, tc.id, res, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestStore_GetByID(t *testing.T) {
	s, mock, ctx := initializeDataStore(t)

	tests := []struct {
		desc    string
		rows    *sqlmock.Rows
		resp    models.Product
		mockErr error
		err     error
	}{
		{"success", sqlmock.NewRows([]string{"id", "name", "brand_name", "details", "image_url"}).
			AddRow(1, "Lays", "pepsico", "tasteBetter", "image.png"), models.Product{
			ID: "1", Name: "Lays", BrandName: "pepsico", Details: "tasteBetter", ImageUrl: "image.png",
		}, nil, nil},
		{"failure", sqlmock.NewRows([]string{"*", "id"}).AddRow(1, 2), models.Product{},
			sql.ErrConnDone, errors.DB{Err: sql.ErrConnDone}},
		{"failure no rows exist", sqlmock.NewRows([]string{"id"}).AddRow(2), models.Product{},
			sql.ErrNoRows, errors.EntityNotFound{Entity: "products", ID: "1"}},
	}

	for i, tc := range tests {
		mock.ExpectQuery(testGetById).WillReturnRows(tc.rows).WillReturnError(tc.mockErr)

		resp, err := s.GetByID(ctx, 1)

		assert.Equal(t, tc.resp, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}
