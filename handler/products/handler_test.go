package products

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"products/models"
	"testing"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"products/service"
)

// initializeContext takes in http method, request path, body and, initializes a gofr context
func initializeContext(method, path string, body io.Reader) (ctx *gofr.Context) {
	req := httptest.NewRequest(method, path, body)
	r := request.NewHTTPRequest(req)

	ctx = gofr.NewContext(nil, r, gofr.New())
	ctx.Context = context.Background()

	return ctx
}

// initializeMocks initializes handler and mock service variables
func initializeMocks(t *testing.T) (handler, *service.MockProducts) {
	ctrl := gomock.NewController(t)
	mockProduct := service.NewMockProducts(ctrl)

	h := New(mockProduct)

	return h, mockProduct
}

func TestCreate(t *testing.T) {
	h, mock := initializeMocks(t)

	product := models.Product{
		Name:      "lays",
		BrandName: "pepsi",
		Details:   "best in taste",
		ImageUrl:  "lays.img",
	}

	productResp := models.Product{
		ID:        "1",
		Name:      "lays",
		BrandName: "pepsi",
		Details:   "best in taste",
		ImageUrl:  "lays.img",
	}

	body, err := json.Marshal(product)
	if err != nil {
		t.Errorf("error in marshalling the body:%v", err)
	}

	testcases := []struct {
		desc     string
		body     []byte
		mockResp models.Product
		output   interface{}
		err      error
		times    int
	}{
		{"success", body, productResp, productResp, nil, 1},
		{"failure", body, models.Product{}, nil, errors.EntityNotFound{}, 1},
		{"invalidBody", []byte(``), models.Product{}, nil,
			errors.InvalidParam{Param: []string{"body"}}, 0},
	}

	for i, tc := range testcases {
		ctx := initializeContext(http.MethodPost, "/product", bytes.NewReader(tc.body))

		mock.EXPECT().Create(ctx, product).Return(productResp, tc.err).Times(tc.times)

		resp, err := h.Create(ctx)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.output, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func Test_GetByID(t *testing.T) {
	h, mock := initializeMocks(t)

	productResp := models.Response{
		Product: models.Product{
			ID:        "1",
			Name:      "lays",
			BrandName: "pepsi",
			Details:   "best in taste",
			ImageUrl:  "lays.img",
		},
		Variant: []models.Variant{
			{
				ID:      "1",
				Name:    "fanta",
				Details: "taste better",
			},
		},
	}

	testcases := []struct {
		desc      string
		pathParam string
		mockResp  models.Response
		output    interface{}
		err       error
		times     int
	}{
		{"success", "1", productResp, productResp, nil, 1},
		{"failure", "1", models.Response{}, nil, errors.EntityNotFound{}, 1},
		{"invalid id", "abc", models.Response{}, nil,
			errors.InvalidParam{Param: []string{"id"}}, 0},
		{"negative id", "-1", models.Response{}, nil,
			errors.InvalidParam{Param: []string{"id"}}, 0},
	}

	for i, tc := range testcases {
		ctx := initializeContext(http.MethodGet, "/product/{id}", nil)
		ctx.SetPathParams(map[string]string{"id": tc.pathParam})

		mock.EXPECT().GetByID(ctx, 1).Return(tc.mockResp, tc.err).Times(tc.times)

		resp, err := h.GetByID(ctx)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.output, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func Test_GetAll(t *testing.T) {
	h, mock := initializeMocks(t)

	variant := []models.Variant{
		{
			ID:      "1",
			Name:    "fanta",
			Details: "taste better",
		},
	}

	productResp := []models.Products{{ID: "1", Name: "lays", BrandName: "pepsi",
		Details: "best in taste", ImageURL: "lays.img", Variants: variant}}

	testcases := []struct {
		desc     string
		mockResp []models.Products
		output   interface{}
		err      error
		times    int
	}{
		{"success", productResp, productResp, nil, 1},
		{"failure", []models.Products{{}}, nil, errors.EntityNotFound{}, 1},
	}

	for i, tc := range testcases {
		ctx := initializeContext(http.MethodGet, "/product", nil)

		mock.EXPECT().GetAll(ctx, map[string]string{}).Return(tc.mockResp, tc.err).Times(tc.times)

		resp, err := h.GetAll(ctx)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.output, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}
