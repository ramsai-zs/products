package products

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"products/models"
	"products/store"
)

func initialiseTest(t *testing.T) (service, *gofr.Context, *store.MockProducts, *store.MockVariants) {
	ctrl := gomock.NewController(t)
	product := store.NewMockProducts(ctrl)
	variant := store.NewMockVariants(ctrl)

	s := New(product, variant)

	ctx := gofr.NewContext(nil, nil, gofr.New())

	return s, ctx, product, variant
}

func Test_Create(t *testing.T) {
	svc, ctx, mockProduct, _ := initialiseTest(t)

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

	testcases := []struct {
		desc           string
		requestPayload models.Product
		mockCreateResp int
		output         models.Product
		err            error
		mockCreateErr  error
		mockGetErr     error
		times          []int
		mockGetResp    models.Product
	}{
		{"success", product, 1, productResp, nil, nil, nil,
			[]int{1, 1}, productResp},
		{"Get failure", product, 1, models.Product{}, errors.EntityNotFound{}, nil,
			errors.EntityNotFound{}, []int{1, 1}, productResp},
		{"create failure", product, 0, models.Product{}, errors.EntityNotFound{}, errors.EntityNotFound{},
			nil, []int{1, 0}, models.Product{}},
		{"missing payload", models.Product{}, 0, models.Product{}, errors.MissingParam{Param: []string{"name", "brandName", "details", "imageUrl"}},
			nil, nil, []int{0, 0}, models.Product{}},
	}

	for i, tc := range testcases {
		mockProduct.EXPECT().Create(ctx, product).Return(tc.mockCreateResp, tc.mockCreateErr).Times(tc.times[0])
		mockProduct.EXPECT().GetByID(ctx, 1).Return(productResp, tc.mockGetErr).Times(tc.times[1])

		resp, err := svc.Create(ctx, tc.requestPayload)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.output, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestGetByIdAndProductId(t *testing.T) {
	svc, ctx, mockProduct, mockVariant := initialiseTest(t)

	variantResp := []models.Variant{{ID: "1", Name: "lays", ProductID: "1", Details: "best in taste"}}
	productResp := models.Product{ID: "1", Name: "productName", BrandName: "brandName", Details: "details",
		ImageUrl: "imageUrl"}

	testcases := []struct {
		desc           string
		variantResp    []models.Variant
		productResp    models.Product
		mockVariantErr error
		mockProductErr error
		resp           models.Response
		times          []int
		err            error
	}{
		{"success", variantResp, productResp, nil, nil,
			models.Response{Variant: variantResp, Product: productResp}, []int{1, 1}, nil},
		{"product failure", variantResp, models.Product{}, nil, errors.EntityNotFound{},
			models.Response{}, []int{1, 1}, errors.EntityNotFound{}},
		{"variant failure", []models.Variant{}, models.Product{}, errors.EntityNotFound{}, nil,
			models.Response{}, []int{1, 0}, errors.EntityNotFound{}},
	}

	for i, tc := range testcases {
		mockVariant.EXPECT().GetByProductId(ctx, 1).Return(tc.variantResp, tc.mockVariantErr).Times(tc.times[0])
		mockProduct.EXPECT().GetByID(ctx, 1).Return(tc.productResp, tc.mockProductErr).Times(tc.times[1])

		resp, err := svc.GetByID(ctx, 1)

		assert.Equal(t, tc.resp, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestGetAll(t *testing.T) {
	svc, ctx, mockProduct, _ := initialiseTest(t)

	resp := []models.Products{{
		ID:        "1",
		Name:      "productName",
		BrandName: "brandName",
		Details:   "details",
		ImageURL:  "img.png",
		Variants: []models.Variant{{
			ID:        "1",
			ProductID: "1",
			Name:      "variantName",
			Details:   "variantDetails",
		}},
	}}

	testcases := []struct {
		desc    string
		filters models.Filters
		resp    []models.Products
		mockErr error
		err     error
	}{
		{"success", models.Filters{ProductID: 1, ProductName: "test", VariantID: 1}, resp, nil, nil},
		{"failure", models.Filters{ProductID: 1, ProductName: "test", VariantID: 1}, nil, errors.EntityNotFound{}, errors.EntityNotFound{}},
	}

	for i, tc := range testcases {
		mockProduct.EXPECT().GetAll(ctx, tc.filters).Return(tc.resp, tc.mockErr)

		resp, err := svc.GetAll(ctx, map[string]string{"productId": "1", "variantId": "1", "productName": "test"})

		assert.Equal(t, tc.resp, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func Test_validateProductBody(t *testing.T) {
	tests := []struct {
		desc    string
		product models.Product
		err     error
	}{
		{"missing payload", models.Product{}, errors.MissingParam{Param: []string{"name", "brandName", "details", "imageUrl"}}},
		{"no missing field", models.Product{Name: "name", BrandName: "brand_name", Details: "details", ImageUrl: "image_url"},
			nil},
		{"missing name", models.Product{BrandName: "brand_name", Details: "details", ImageUrl: "image_url"},
			errors.MissingParam{Param: []string{"name"}}},
		{"missing brand name", models.Product{Name: "name", Details: "details", ImageUrl: "image_url"},
			errors.MissingParam{Param: []string{"brandName"}}},
		{"missing details", models.Product{Name: "name", BrandName: "brand_name", ImageUrl: "image_url"},
			errors.MissingParam{Param: []string{"details"}}},
		{"missing imageURL", models.Product{Name: "name", BrandName: "brand_name", Details: "details"},
			errors.MissingParam{Param: []string{"imageUrl"}}},
	}

	for i, tc := range tests {
		err := validateProductBody(tc.product)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestValidateFilters(t *testing.T) {
	testcases := []struct {
		desc    string
		params  map[string]string
		filters models.Filters
		err     error
	}{
		{"all params exist", map[string]string{"productId": "1", "variantId": "1", "productName": "test"},
			models.Filters{ProductID: 1, ProductName: "test", VariantID: 1}, nil},
		{"missing productId", map[string]string{"variantId": "1", "productName": "test"},
			models.Filters{ProductName: "test", VariantID: 1}, nil},
		{"missing variantId", map[string]string{"productId": "1", "productName": "test"},
			models.Filters{ProductID: 1, ProductName: "test"}, nil},
		{"invalid productId", map[string]string{"productId": "1@1", "variantId": "1", "productName": "test"},
			models.Filters{}, errors.InvalidParam{Param: []string{"productId"}}},
		{"invalid variantId", map[string]string{"productId": "1", "variantId": "1@", "productName": "test"},
			models.Filters{}, errors.InvalidParam{Param: []string{"variantId"}}},
	}

	for i, tc := range testcases {
		filter, err := validateFilters(tc.params)

		assert.Equal(t, tc.filters, filter, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}
