package variants

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"products/models"
	"products/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
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
func initializeMocks(t *testing.T) (handler, *service.MockVariants) {
	ctrl := gomock.NewController(t)
	mockVariant := service.NewMockVariants(ctrl)

	h := New(mockVariant)

	return h, mockVariant
}

func TestCreate(t *testing.T) {
	h, mock := initializeMocks(t)

	variant := models.Variant{
		Name:    "test",
		Details: "details of the above variant",
	}

	variantResp := models.Variant{
		ID:        "1",
		ProductID: "2",
		Name:      "test",
		Details:   "details of the above variant",
	}

	body, err := json.Marshal(variant)
	if err != nil {
		t.Errorf("error in marshalling the body:%v", err)
	}

	testcases := []struct {
		desc     string
		body     []byte
		mockResp models.Variant
		output   interface{}
		err      error
		times    int
	}{
		{"success", body, variantResp, variantResp, nil, 1},
		{"failure", body, models.Variant{}, nil, errors.EntityNotFound{}, 1},
		{"invalidBody", []byte(``), models.Variant{}, nil,
			errors.InvalidParam{Param: []string{"body"}}, 0},
	}

	for i, tc := range testcases {
		ctx := initializeContext(http.MethodPost, "/product", bytes.NewReader(tc.body))
		ctx.SetPathParams(map[string]string{"pid": "2"})

		mock.EXPECT().Create(ctx, models.Variant{ProductID: "2", Name: "test",
			Details: "details of the above variant"}).Return(variantResp, tc.err).Times(tc.times)

		resp, err := h.Create(ctx)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.output, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestGetByID(t *testing.T) {
	h, mock := initializeMocks(t)

	variantResp := models.Variant{
		ID:        "1",
		ProductID: "2",
		Name:      "test",
		Details:   "details of the above variant",
	}

	testcases := []struct {
		desc      string
		variantID string
		productID string
		mockResp  models.Variant
		output    interface{}
		err       error
		times     int
	}{
		{"success", "1", "1", variantResp, variantResp, nil, 1},
		{"failure", "1", "1", models.Variant{}, nil,
			errors.EntityNotFound{}, 1},
		{"invalid productId", "1", "abc", models.Variant{}, nil,
			errors.InvalidParam{Param: []string{"pid"}}, 0},
		{"invalid id", "abc", "1", models.Variant{}, nil,
			errors.InvalidParam{Param: []string{"id"}}, 0},
	}

	for i, tc := range testcases {
		ctx := initializeContext(http.MethodGet, "/products/1/variant/1", nil)
		ctx.SetPathParams(map[string]string{"pid": tc.productID, "id": tc.variantID})

		mock.EXPECT().GetByIdAndProductId(ctx, 1, 1).Return(tc.mockResp, tc.err).Times(tc.times)

		resp, err := h.GetByID(ctx)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.output, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}

}

// Test_validateID : tests the positive and negative cases of method validateID
func Test_validateID(t *testing.T) {
	tests := []struct {
		desc  string
		val   string
		param string
		res   int
		err   error
	}{
		{"success case", "10", "id", 10, nil},
		{"invalid param", "o", "id", 0, errors.InvalidParam{Param: []string{"id"}}},
		{"negative param", "-10", "id", 0, errors.InvalidParam{Param: []string{"id"}}},
	}

	for i, tc := range tests {
		res, err := validateID(tc.val, tc.param)

		assert.Equal(t, tc.res, res, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}
