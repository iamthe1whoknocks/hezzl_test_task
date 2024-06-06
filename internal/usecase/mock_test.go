// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/interfaces.go
//
// Generated by this command:
//
//	mockgen -source ./internal/usecase/interfaces.go -package usecase_test
//
// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	models "github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockItem is a mock of Item interface.
type MockItem struct {
	ctrl     *gomock.Controller
	recorder *MockItemMockRecorder
}

// MockItemMockRecorder is the mock recorder for MockItem.
type MockItemMockRecorder struct {
	mock *MockItem
}

// NewMockItem creates a new mock instance.
func NewMockItem(ctrl *gomock.Controller) *MockItem {
	mock := &MockItem{ctrl: ctrl}
	mock.recorder = &MockItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItem) EXPECT() *MockItemMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockItem) Delete(arg0 context.Context, arg1, arg2 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockItemMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItem)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockItem) Get(arg0 context.Context) ([]models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockItemMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockItem)(nil).Get), arg0)
}

// GetCache mocks base method.
func (m *MockItem) GetCache(ctx context.Context, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", ctx, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache.
func (mr *MockItemMockRecorder) GetCache(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockItem)(nil).GetCache), ctx, key)
}

// InvalidateCache mocks base method.
func (m *MockItem) InvalidateCache(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateCache", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateCache indicates an expected call of InvalidateCache.
func (mr *MockItemMockRecorder) InvalidateCache(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateCache", reflect.TypeOf((*MockItem)(nil).InvalidateCache), ctx, key)
}

// Save mocks base method.
func (m *MockItem) Save(arg0 context.Context, arg1 *models.Item) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockItemMockRecorder) Save(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockItem)(nil).Save), arg0, arg1)
}

// SetCache mocks base method.
func (m *MockItem) SetCache(ctx context.Context, key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCache", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCache indicates an expected call of SetCache.
func (mr *MockItemMockRecorder) SetCache(ctx, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCache", reflect.TypeOf((*MockItem)(nil).SetCache), ctx, key, value)
}

// Update mocks base method.
func (m *MockItem) Update(arg0 context.Context, arg1 *models.Item) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockItemMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItem)(nil).Update), arg0, arg1)
}

// MockItemsRepo is a mock of ItemsRepo interface.
type MockItemsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockItemsRepoMockRecorder
}

// MockItemsRepoMockRecorder is the mock recorder for MockItemsRepo.
type MockItemsRepoMockRecorder struct {
	mock *MockItemsRepo
}

// NewMockItemsRepo creates a new mock instance.
func NewMockItemsRepo(ctrl *gomock.Controller) *MockItemsRepo {
	mock := &MockItemsRepo{ctrl: ctrl}
	mock.recorder = &MockItemsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemsRepo) EXPECT() *MockItemsRepoMockRecorder {
	return m.recorder
}

// DeleteItem mocks base method.
func (m *MockItemsRepo) DeleteItem(arg0 context.Context, arg1, arg2 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockItemsRepoMockRecorder) DeleteItem(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockItemsRepo)(nil).DeleteItem), arg0, arg1, arg2)
}

// GetItems mocks base method.
func (m *MockItemsRepo) GetItems(arg0 context.Context) ([]models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", arg0)
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockItemsRepoMockRecorder) GetItems(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockItemsRepo)(nil).GetItems), arg0)
}

// SaveItem mocks base method.
func (m *MockItemsRepo) SaveItem(arg0 context.Context, arg1 *models.Item) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveItem", arg0, arg1)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveItem indicates an expected call of SaveItem.
func (mr *MockItemsRepoMockRecorder) SaveItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveItem", reflect.TypeOf((*MockItemsRepo)(nil).SaveItem), arg0, arg1)
}

// UpdateItem mocks base method.
func (m *MockItemsRepo) UpdateItem(arg0 context.Context, arg1 *models.Item) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", arg0, arg1)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockItemsRepoMockRecorder) UpdateItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockItemsRepo)(nil).UpdateItem), arg0, arg1)
}

// MockCacher is a mock of Cacher interface.
type MockCacher struct {
	ctrl     *gomock.Controller
	recorder *MockCacherMockRecorder
}

// MockCacherMockRecorder is the mock recorder for MockCacher.
type MockCacherMockRecorder struct {
	mock *MockCacher
}

// NewMockCacher creates a new mock instance.
func NewMockCacher(ctrl *gomock.Controller) *MockCacher {
	mock := &MockCacher{ctrl: ctrl}
	mock.recorder = &MockCacherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacher) EXPECT() *MockCacherMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCacher) Get(ctx context.Context, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacherMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacher)(nil).Get), ctx, key)
}

// Invalidate mocks base method.
func (m *MockCacher) Invalidate(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidate", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate.
func (mr *MockCacherMockRecorder) Invalidate(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockCacher)(nil).Invalidate), ctx, key)
}

// Set mocks base method.
func (m *MockCacher) Set(ctx context.Context, key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacherMockRecorder) Set(ctx, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacher)(nil).Set), ctx, key, value)
}

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// GetSubject mocks base method.
func (m *MockBroker) GetSubject() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubject")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSubject indicates an expected call of GetSubject.
func (mr *MockBrokerMockRecorder) GetSubject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubject", reflect.TypeOf((*MockBroker)(nil).GetSubject))
}

// Publish mocks base method.
func (m *MockBroker) Publish(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockBrokerMockRecorder) Publish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockBroker)(nil).Publish), arg0, arg1)
}

// Subscriber mocks base method.
func (m *MockBroker) Subscriber() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscriber")
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscriber indicates an expected call of Subscriber.
func (mr *MockBrokerMockRecorder) Subscriber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscriber", reflect.TypeOf((*MockBroker)(nil).Subscriber))
}