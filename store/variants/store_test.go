package variants

import (
	"context"
	"database/sql"
	"gofr.dev/pkg/errors"
	"products/models"
	"testing"

	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/gofr"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	testCreate              = "INSERT INTO varaints(product_id,variant_name,variant_details) VALUES(?,?,?)"
	testGetById             = "SELECT id,product_id,variant_name,variant_details FROM varaints where id=?"
	testGetByProductId      = "SELECT id,variant_name,variant_details FROM varaints where product_id=?"
	testGetByIdAndProductId = "SELECT id,product_id,variant_name,variant_details FROM varaints where id=? AND product_id=?"
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

	str := New()

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
		{"success", 1, mock.ExpectExec(testCreate).WithArgs("1", "test", "test-product").
			WillReturnResult(sqlmock.NewResult(1, 1)), nil},
		{"failure - error while inserting", 0,
			mock.ExpectExec(testCreate).WithArgs("1", "test", "test-product").WillReturnError(errors.EntityAlreadyExists{}),
			errors.DB{Err: errors.EntityAlreadyExists{}}},
		{"failure - last insert id 0", 0, mock.ExpectExec(testCreate).WithArgs("1", "test", "test-product").
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows)), errors.DB{Err: sql.ErrNoRows}},
	}

	for i, tc := range tests {
		res, err := s.Create(ctx, models.Variant{ID: "1", ProductID: "1", Name: "test", Details: "test-product"})

		assert.Equal(t, tc.id, res, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestStore_GetByID(t *testing.T) {
	s, mock, ctx := initializeDataStore(t)

	tests := []struct {
		desc    string
		rows    *sqlmock.Rows
		resp    models.Variant
		mockErr error
		err     error
	}{
		{"success", sqlmock.NewRows([]string{"id", "product_id", "name", "details"}).AddRow(1, "1", "name",
			"details"), models.Variant{ID: "1", ProductID: "1", Name: "name", Details: "details"}, nil, nil},
		{"failure", sqlmock.NewRows([]string{"*", "id"}).AddRow(1, 2), models.Variant{},
			sql.ErrConnDone, errors.DB{Err: sql.ErrConnDone}},
		{"failure no rows exist", sqlmock.NewRows([]string{"id"}).AddRow(2), models.Variant{},
			sql.ErrNoRows, errors.EntityNotFound{Entity: "variants", ID: "1"}},
	}

	for i, tc := range tests {
		mock.ExpectQuery(testGetById).WillReturnRows(tc.rows).WillReturnError(tc.mockErr)

		resp, err := s.GetByID(ctx, 1)

		assert.Equal(t, tc.resp, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestStore_GetByIdAndProductId(t *testing.T) {
	s, mock, ctx := initializeDataStore(t)

	tests := []struct {
		desc    string
		rows    *sqlmock.Rows
		resp    models.Variant
		mockErr error
		err     error
	}{
		{"success", sqlmock.NewRows([]string{"id", "product_id", "name", "details"}).AddRow(1, "1", "name",
			"details"), models.Variant{ID: "1", ProductID: "1", Name: "name", Details: "details"}, nil, nil},
		{"failure", sqlmock.NewRows([]string{"*", "id"}).AddRow(1, 2), models.Variant{},
			sql.ErrConnDone, errors.DB{Err: sql.ErrConnDone}},
		{"failure no rows exist", sqlmock.NewRows([]string{"id"}).AddRow(2), models.Variant{},
			sql.ErrNoRows, errors.EntityNotFound{Entity: "variants", ID: "id :1,productId:1"}},
	}

	for i, tc := range tests {
		mock.ExpectQuery(testGetByIdAndProductId).WillReturnRows(tc.rows).WillReturnError(tc.mockErr)

		resp, err := s.GetByIdAndProductId(ctx, 1, 1)

		assert.Equal(t, tc.resp, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func Test_getQuery(t *testing.T) {
	tests := []struct {
		desc       string
		ids        []string
		queryParam []interface{}
		query      string
	}{
		{"no id", []string{}, []interface{}{}, "SELECT id,product_id,variant_name,variant_details FROM varaints where product_id in ()"},
		{"single id", []string{"1"}, []interface{}{"1"}, "SELECT id,product_id,variant_name,variant_details FROM varaints where product_id in (?)"},
		{"multiple ids", []string{"1", "3"}, []interface{}{"1", "3"}, "SELECT id,product_id,variant_name,variant_details FROM varaints where product_id in (?,?)"},
	}

	for i, tc := range tests {
		qp, query := getQuery(tc.ids)

		assert.Equal(t, tc.query, query, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.queryParam, qp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}
