// Code generated by MockGen. DO NOT EDIT.
// Source: simple_bank/api (interfaces: Handler)

// Package mockapi is a generated GoMock package.
package mockapi

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	ctrl "simple_bank/api/controller"
	models "simple_bank/db/models"
	repository "simple_bank/db/repository"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method
func (m *MockHandler) AddAccountBalance(arg0 context.Context, arg1 repository.AddAccountBalanceParams) (models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance
func (mr *MockHandlerMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockHandler)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method
func (m *MockHandler) CreateAccount(arg0 context.Context, arg1 repository.CreateAccountParams) (models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockHandlerMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockHandler)(nil).CreateAccount), arg0, arg1)
}

// CreateEntry mocks base method
func (m *MockHandler) CreateEntry(arg0 context.Context, arg1 repository.CreateEntryParams) (models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry
func (mr *MockHandlerMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockHandler)(nil).CreateEntry), arg0, arg1)
}

// CreateTransfer mocks base method
func (m *MockHandler) CreateTransfer(arg0 context.Context, arg1 repository.CreateTransferParams) (models.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(models.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer
func (mr *MockHandlerMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockHandler)(nil).CreateTransfer), arg0, arg1)
}

// DeleteAccount mocks base method
func (m *MockHandler) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockHandlerMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockHandler)(nil).DeleteAccount), arg0, arg1)
}

// DeleteEntry mocks base method
func (m *MockHandler) DeleteEntry(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntry indicates an expected call of DeleteEntry
func (mr *MockHandlerMockRecorder) DeleteEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockHandler)(nil).DeleteEntry), arg0, arg1)
}

// DeleteTransfer mocks base method
func (m *MockHandler) DeleteTransfer(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransfer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransfer indicates an expected call of DeleteTransfer
func (mr *MockHandlerMockRecorder) DeleteTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransfer", reflect.TypeOf((*MockHandler)(nil).DeleteTransfer), arg0, arg1)
}

// GetAccount mocks base method
func (m *MockHandler) GetAccount(arg0 context.Context, arg1 int64) (models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockHandlerMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockHandler)(nil).GetAccount), arg0, arg1)
}

// GetEntry mocks base method
func (m *MockHandler) GetEntry(arg0 context.Context, arg1 int64) (models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry
func (mr *MockHandlerMockRecorder) GetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockHandler)(nil).GetEntry), arg0, arg1)
}

// GetRandomId mocks base method
func (m *MockHandler) GetRandomId(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomId", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomId indicates an expected call of GetRandomId
func (mr *MockHandlerMockRecorder) GetRandomId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomId", reflect.TypeOf((*MockHandler)(nil).GetRandomId), arg0)
}

// GetTransfer mocks base method
func (m *MockHandler) GetTransfer(arg0 context.Context, arg1 int64) (models.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(models.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer
func (mr *MockHandlerMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockHandler)(nil).GetTransfer), arg0, arg1)
}

// ListAccounts mocks base method
func (m *MockHandler) ListAccounts(arg0 context.Context, arg1 repository.ListAccountsParams) ([]models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts
func (mr *MockHandlerMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockHandler)(nil).ListAccounts), arg0, arg1)
}

// ListEntry mocks base method
func (m *MockHandler) ListEntry(arg0 context.Context, arg1 repository.ListEntryParams) ([]models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntry", arg0, arg1)
	ret0, _ := ret[0].([]models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntry indicates an expected call of ListEntry
func (mr *MockHandlerMockRecorder) ListEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntry", reflect.TypeOf((*MockHandler)(nil).ListEntry), arg0, arg1)
}

// ListTransfer mocks base method
func (m *MockHandler) ListTransfer(arg0 context.Context, arg1 repository.ListTransferParams) ([]models.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfer", arg0, arg1)
	ret0, _ := ret[0].([]models.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfer indicates an expected call of ListTransfer
func (mr *MockHandlerMockRecorder) ListTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfer", reflect.TypeOf((*MockHandler)(nil).ListTransfer), arg0, arg1)
}

// ST_C_TransferTx mocks base method
func (m *MockHandler) ST_C_TransferTx(arg0 context.Context, arg1 ctrl.TransferTxParams) (ctrl.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ST_C_TransferTx", arg0, arg1)
	ret0, _ := ret[0].(ctrl.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ST_C_TransferTx indicates an expected call of ST_C_TransferTx
func (mr *MockHandlerMockRecorder) ST_C_TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ST_C_TransferTx", reflect.TypeOf((*MockHandler)(nil).ST_C_TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method
func (m *MockHandler) UpdateAccount(arg0 context.Context, arg1 repository.UpdateAccountParams) (models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockHandlerMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockHandler)(nil).UpdateAccount), arg0, arg1)
}

// UpdateEntry mocks base method
func (m *MockHandler) UpdateEntry(arg0 context.Context, arg1 repository.UpdateEntryParams) (models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntry", arg0, arg1)
	ret0, _ := ret[0].(models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntry indicates an expected call of UpdateEntry
func (mr *MockHandlerMockRecorder) UpdateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockHandler)(nil).UpdateEntry), arg0, arg1)
}

// UpdateTransfer mocks base method
func (m *MockHandler) UpdateTransfer(arg0 context.Context, arg1 repository.UpdateTransferParams) (models.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransfer", arg0, arg1)
	ret0, _ := ret[0].(models.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransfer indicates an expected call of UpdateTransfer
func (mr *MockHandlerMockRecorder) UpdateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransfer", reflect.TypeOf((*MockHandler)(nil).UpdateTransfer), arg0, arg1)
}
