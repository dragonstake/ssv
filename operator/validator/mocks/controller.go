// Code generated by MockGen. DO NOT EDIT.
// Source: ./controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
	validator "github.com/bloxapp/ssv/protocol/v2/ssv/validator"
	types "github.com/bloxapp/ssv/protocol/v2/types"
	storage "github.com/bloxapp/ssv/registry/storage"
	basedb "github.com/bloxapp/ssv/storage/basedb"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	zap "go.uber.org/zap"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// ActiveValidatorIndices mocks base method.
func (m *MockController) ActiveValidatorIndices(logger *zap.Logger) []phase0.ValidatorIndex {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveValidatorIndices", logger)
	ret0, _ := ret[0].([]phase0.ValidatorIndex)
	return ret0
}

// ActiveValidatorIndices indicates an expected call of ActiveValidatorIndices.
func (mr *MockControllerMockRecorder) ActiveValidatorIndices(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveValidatorIndices", reflect.TypeOf((*MockController)(nil).ActiveValidatorIndices), logger)
}

// GetOperatorData mocks base method.
func (m *MockController) GetOperatorData() *storage.OperatorData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorData")
	ret0, _ := ret[0].(*storage.OperatorData)
	return ret0
}

// GetOperatorData indicates an expected call of GetOperatorData.
func (mr *MockControllerMockRecorder) GetOperatorData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorData", reflect.TypeOf((*MockController)(nil).GetOperatorData))
}

// GetOperatorShares mocks base method.
func (m *MockController) GetOperatorShares() []*types.SSVShare {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorShares")
	ret0, _ := ret[0].([]*types.SSVShare)
	return ret0
}

// GetOperatorShares indicates an expected call of GetOperatorShares.
func (mr *MockControllerMockRecorder) GetOperatorShares() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorShares", reflect.TypeOf((*MockController)(nil).GetOperatorShares))
}

// GetValidator mocks base method.
func (m *MockController) GetValidator(pubKey string) (*validator.Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidator", pubKey)
	ret0, _ := ret[0].(*validator.Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidator indicates an expected call of GetValidator.
func (mr *MockControllerMockRecorder) GetValidator(pubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockController)(nil).GetValidator), pubKey)
}

// GetValidatorStats mocks base method.
func (m *MockController) GetValidatorStats() (uint64, uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorStats")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetValidatorStats indicates an expected call of GetValidatorStats.
func (mr *MockControllerMockRecorder) GetValidatorStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorStats", reflect.TypeOf((*MockController)(nil).GetValidatorStats))
}

// LiquidateCluster mocks base method.
func (m *MockController) LiquidateCluster(owner common.Address, operatorIDs []uint64, toLiquidate []*types.SSVShare) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LiquidateCluster", owner, operatorIDs, toLiquidate)
	ret0, _ := ret[0].(error)
	return ret0
}

// LiquidateCluster indicates an expected call of LiquidateCluster.
func (mr *MockControllerMockRecorder) LiquidateCluster(owner, operatorIDs, toLiquidate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LiquidateCluster", reflect.TypeOf((*MockController)(nil).LiquidateCluster), owner, operatorIDs, toLiquidate)
}

// ReactivateCluster mocks base method.
func (m *MockController) ReactivateCluster(owner common.Address, operatorIDs []uint64, toReactivate []*types.SSVShare) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReactivateCluster", owner, operatorIDs, toReactivate)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReactivateCluster indicates an expected call of ReactivateCluster.
func (mr *MockControllerMockRecorder) ReactivateCluster(owner, operatorIDs, toReactivate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReactivateCluster", reflect.TypeOf((*MockController)(nil).ReactivateCluster), owner, operatorIDs, toReactivate)
}

// StartNetworkHandlers mocks base method.
func (m *MockController) StartNetworkHandlers(logger *zap.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartNetworkHandlers", logger)
}

// StartNetworkHandlers indicates an expected call of StartNetworkHandlers.
func (mr *MockControllerMockRecorder) StartNetworkHandlers(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNetworkHandlers", reflect.TypeOf((*MockController)(nil).StartNetworkHandlers), logger)
}

// StartValidator mocks base method.
func (m *MockController) StartValidator(share *types.SSVShare) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartValidator", share)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartValidator indicates an expected call of StartValidator.
func (mr *MockControllerMockRecorder) StartValidator(share interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartValidator", reflect.TypeOf((*MockController)(nil).StartValidator), share)
}

// StartValidators mocks base method.
func (m *MockController) StartValidators(logger *zap.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartValidators", logger)
}

// StartValidators indicates an expected call of StartValidators.
func (mr *MockControllerMockRecorder) StartValidators(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartValidators", reflect.TypeOf((*MockController)(nil).StartValidators), logger)
}

// StopValidator mocks base method.
func (m *MockController) StopValidator(publicKey []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopValidator", publicKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopValidator indicates an expected call of StopValidator.
func (mr *MockControllerMockRecorder) StopValidator(publicKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopValidator", reflect.TypeOf((*MockController)(nil).StopValidator), publicKey)
}

// UpdateFeeRecipient mocks base method.
func (m *MockController) UpdateFeeRecipient(owner, recipient common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFeeRecipient", owner, recipient)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFeeRecipient indicates an expected call of UpdateFeeRecipient.
func (mr *MockControllerMockRecorder) UpdateFeeRecipient(owner, recipient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFeeRecipient", reflect.TypeOf((*MockController)(nil).UpdateFeeRecipient), owner, recipient)
}

// UpdateValidatorMetaDataLoop mocks base method.
func (m *MockController) UpdateValidatorMetaDataLoop(logger *zap.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateValidatorMetaDataLoop", logger)
}

// UpdateValidatorMetaDataLoop indicates an expected call of UpdateValidatorMetaDataLoop.
func (mr *MockControllerMockRecorder) UpdateValidatorMetaDataLoop(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorMetaDataLoop", reflect.TypeOf((*MockController)(nil).UpdateValidatorMetaDataLoop), logger)
}

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// BumpNonce mocks base method.
func (m *MockEventHandler) BumpNonce(txn basedb.Txn, owner common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BumpNonce", txn, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// BumpNonce indicates an expected call of BumpNonce.
func (mr *MockEventHandlerMockRecorder) BumpNonce(txn, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BumpNonce", reflect.TypeOf((*MockEventHandler)(nil).BumpNonce), txn, owner)
}

// GetEventData mocks base method.
func (m *MockEventHandler) GetEventData(txHash common.Hash) (*storage.EventData, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventData", txHash)
	ret0, _ := ret[0].(*storage.EventData)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEventData indicates an expected call of GetEventData.
func (mr *MockEventHandlerMockRecorder) GetEventData(txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventData", reflect.TypeOf((*MockEventHandler)(nil).GetEventData), txHash)
}

// GetNextNonce mocks base method.
func (m *MockEventHandler) GetNextNonce(txn basedb.Txn, owner common.Address) (storage.Nonce, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextNonce", txn, owner)
	ret0, _ := ret[0].(storage.Nonce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextNonce indicates an expected call of GetNextNonce.
func (mr *MockEventHandlerMockRecorder) GetNextNonce(txn, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextNonce", reflect.TypeOf((*MockEventHandler)(nil).GetNextNonce), txn, owner)
}

// SaveEventData mocks base method.
func (m *MockEventHandler) SaveEventData(txHash common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEventData", txHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveEventData indicates an expected call of SaveEventData.
func (mr *MockEventHandlerMockRecorder) SaveEventData(txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEventData", reflect.TypeOf((*MockEventHandler)(nil).SaveEventData), txHash)
}
