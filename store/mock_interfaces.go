// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package store is a generated GoMock package.
package store

import (
	models "products/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

// MockProducts is a mock of Products interface.
type MockProducts struct {
	ctrl     *gomock.Controller
	recorder *MockProductsMockRecorder
}

// MockProductsMockRecorder is the mock recorder for MockProducts.
type MockProductsMockRecorder struct {
	mock *MockProducts
}

// NewMockProducts creates a new mock instance.
func NewMockProducts(ctrl *gomock.Controller) *MockProducts {
	mock := &MockProducts{ctrl: ctrl}
	mock.recorder = &MockProductsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducts) EXPECT() *MockProductsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProducts) Create(ctx *gofr.Context, p models.Product) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductsMockRecorder) Create(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProducts)(nil).Create), ctx, p)
}

// GetAll mocks base method.
func (m *MockProducts) GetAll(ctx *gofr.Context, filters models.Filters) ([]models.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, filters)
	ret0, _ := ret[0].([]models.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductsMockRecorder) GetAll(ctx, filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProducts)(nil).GetAll), ctx, filters)
}

// GetByID mocks base method.
func (m *MockProducts) GetByID(ctx *gofr.Context, id int) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProductsMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProducts)(nil).GetByID), ctx, id)
}

// MockVariants is a mock of Variants interface.
type MockVariants struct {
	ctrl     *gomock.Controller
	recorder *MockVariantsMockRecorder
}

// MockVariantsMockRecorder is the mock recorder for MockVariants.
type MockVariantsMockRecorder struct {
	mock *MockVariants
}

// NewMockVariants creates a new mock instance.
func NewMockVariants(ctrl *gomock.Controller) *MockVariants {
	mock := &MockVariants{ctrl: ctrl}
	mock.recorder = &MockVariantsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVariants) EXPECT() *MockVariantsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVariants) Create(ctx *gofr.Context, v models.Variant) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, v)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVariantsMockRecorder) Create(ctx, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVariants)(nil).Create), ctx, v)
}

// GetByID mocks base method.
func (m *MockVariants) GetByID(ctx *gofr.Context, id int) (models.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockVariantsMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockVariants)(nil).GetByID), ctx, id)
}

// GetByIdAndProductId mocks base method.
func (m *MockVariants) GetByIdAndProductId(ctx *gofr.Context, id, pID int) (models.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdAndProductId", ctx, id, pID)
	ret0, _ := ret[0].(models.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdAndProductId indicates an expected call of GetByIdAndProductId.
func (mr *MockVariantsMockRecorder) GetByIdAndProductId(ctx, id, pID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdAndProductId", reflect.TypeOf((*MockVariants)(nil).GetByIdAndProductId), ctx, id, pID)
}

// GetByMultipleProductId mocks base method.
func (m *MockVariants) GetByMultipleProductId(ctx *gofr.Context, pIDs []string) ([]models.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMultipleProductId", ctx, pIDs)
	ret0, _ := ret[0].([]models.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMultipleProductId indicates an expected call of GetByMultipleProductId.
func (mr *MockVariantsMockRecorder) GetByMultipleProductId(ctx, pIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMultipleProductId", reflect.TypeOf((*MockVariants)(nil).GetByMultipleProductId), ctx, pIDs)
}

// GetByProductId mocks base method.
func (m *MockVariants) GetByProductId(ctx *gofr.Context, pId int) ([]models.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByProductId", ctx, pId)
	ret0, _ := ret[0].([]models.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByProductId indicates an expected call of GetByProductId.
func (mr *MockVariantsMockRecorder) GetByProductId(ctx, pId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByProductId", reflect.TypeOf((*MockVariants)(nil).GetByProductId), ctx, pId)
}
