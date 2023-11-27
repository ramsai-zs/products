package variants

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
	svc, ctx, mockProduct, mockVariant := initialiseTest(t)

	variant := models.Variant{Name: "lays", ProductID: "1", Details: "best in taste"}
	variantResp := models.Variant{ID: "1", Name: "lays", ProductID: "1", Details: "best in taste"}

	testcases := []struct {
		desc           string
		productErr     error
		requestPayload models.Variant
		mockCreateResp int
		output         models.Variant
		err            error
		mockCreateErr  error
		mockGetErr     error
		times          []int
		mockGetResp    models.Variant
	}{
		{"success", nil, variant, 1, variantResp, nil, nil, nil,
			[]int{1, 1, 1}, variantResp},
		{"variant Get Failure", nil, variant, 1, models.Variant{}, errors.EntityNotFound{}, nil,
			errors.EntityNotFound{}, []int{1, 1, 1}, models.Variant{}},
		{"insertion Failure", nil, variant, 0, models.Variant{}, errors.EntityNotFound{}, errors.EntityNotFound{},
			nil, []int{1, 1, 0}, models.Variant{}},
		{"empty payload", nil, models.Variant{ProductID: "1"}, 0,
			models.Variant{}, errors.MissingParam{Param: []string{"name", "details"}}, nil,
			nil, []int{1, 0, 0}, models.Variant{}},
		{"getByID failure", errors.EntityNotFound{}, variant, 0,
			models.Variant{}, errors.EntityNotFound{}, nil,
			nil, []int{1, 0, 0}, models.Variant{}},
		{"invalid productId", nil, models.Variant{ProductID: "ab"}, 0,
			models.Variant{}, errors.InvalidParam{Param: []string{"productId"}}, nil,
			nil, []int{0, 0, 0}, models.Variant{}},
		{"productId is zero", nil, models.Variant{ProductID: "0"}, 0,
			models.Variant{}, errors.InvalidParam{Param: []string{"productId"}}, nil,
			nil, []int{0, 0, 0}, models.Variant{}},
	}

	for i, tc := range testcases {
		mockProduct.EXPECT().GetByID(ctx, 1).Return(models.Product{}, tc.productErr).Times(tc.times[0])
		mockVariant.EXPECT().Create(ctx, variant).Return(tc.mockCreateResp, tc.mockCreateErr).Times(tc.times[1])
		mockVariant.EXPECT().GetByID(ctx, 1).Return(tc.mockGetResp, tc.mockGetErr).Times(tc.times[2])

		resp, err := svc.Create(ctx, tc.requestPayload)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.output, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func TestGetByIdAndProductId(t *testing.T) {
	svc, ctx, _, mockVariant := initialiseTest(t)

	variantResp := models.Variant{ID: "1", Name: "lays", ProductID: "1", Details: "best in taste"}

	testcases := []struct {
		desc    string
		resp    models.Variant
		mockErr error
		err     error
	}{
		{"success", variantResp, nil, nil},
		{"failure", models.Variant{}, errors.EntityNotFound{}, errors.EntityNotFound{}},
	}

	for i, tc := range testcases {
		mockVariant.EXPECT().GetByIdAndProductId(ctx, 1, 1).Return(tc.resp, tc.mockErr)

		resp, err := svc.GetByIdAndProductId(ctx, 1, 1)

		assert.Equal(t, tc.resp, resp, "Test[%d] failed.\n Desc : %s", i, tc.desc)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}

func Test_validateVariantBody(t *testing.T) {
	testcases := []struct {
		desc    string
		variant models.Variant
		err     error
	}{
		{"empty payload", models.Variant{}, errors.MissingParam{Param: []string{"name", "details"}}},
		{"missing name", models.Variant{Details: "variant details"}, errors.MissingParam{Param: []string{"name"}}},
		{"missing details", models.Variant{Name: "variant name"}, errors.MissingParam{Param: []string{"details"}}},
	}

	for i, tc := range testcases {
		err := validateVariantBody(tc.variant)

		assert.Equal(t, tc.err, err, "Test[%d] failed.\n Desc : %s", i, tc.desc)
	}
}
