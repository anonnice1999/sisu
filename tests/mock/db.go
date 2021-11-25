// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/sisu-network/sisu/x/tss/types"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDatabase) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaseMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabase)(nil).Close))
}

// CreateKeygen mocks base method.
func (m *MockDatabase) CreateKeygen(chain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKeygen", chain)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateKeygen indicates an expected call of CreateKeygen.
func (mr *MockDatabaseMockRecorder) CreateKeygen(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKeygen", reflect.TypeOf((*MockDatabase)(nil).CreateKeygen), chain)
}

// GetContractFromAddress mocks base method.
func (m *MockDatabase) GetContractFromAddress(chain, address string) *types.ContractEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractFromAddress", chain, address)
	ret0, _ := ret[0].(*types.ContractEntity)
	return ret0
}

// GetContractFromAddress indicates an expected call of GetContractFromAddress.
func (mr *MockDatabaseMockRecorder) GetContractFromAddress(chain, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractFromAddress", reflect.TypeOf((*MockDatabase)(nil).GetContractFromAddress), chain, address)
}

// GetContractFromHash mocks base method.
func (m *MockDatabase) GetContractFromHash(chain, hash string) *types.ContractEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractFromHash", chain, hash)
	ret0, _ := ret[0].(*types.ContractEntity)
	return ret0
}

// GetContractFromHash indicates an expected call of GetContractFromHash.
func (mr *MockDatabaseMockRecorder) GetContractFromHash(chain, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractFromHash", reflect.TypeOf((*MockDatabase)(nil).GetContractFromHash), chain, hash)
}

// GetKeygenStatus mocks base method.
func (m *MockDatabase) GetKeygenStatus(chain string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeygenStatus", chain)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeygenStatus indicates an expected call of GetKeygenStatus.
func (mr *MockDatabaseMockRecorder) GetKeygenStatus(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeygenStatus", reflect.TypeOf((*MockDatabase)(nil).GetKeygenStatus), chain)
}

// GetPendingDeployContracts mocks base method.
func (m *MockDatabase) GetPendingDeployContracts(chain string) []*types.ContractEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingDeployContracts", chain)
	ret0, _ := ret[0].([]*types.ContractEntity)
	return ret0
}

// GetPendingDeployContracts indicates an expected call of GetPendingDeployContracts.
func (mr *MockDatabaseMockRecorder) GetPendingDeployContracts(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingDeployContracts", reflect.TypeOf((*MockDatabase)(nil).GetPendingDeployContracts), chain)
}

// GetPubKey mocks base method.
func (m *MockDatabase) GetPubKey(chain string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKey", chain)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetPubKey indicates an expected call of GetPubKey.
func (mr *MockDatabaseMockRecorder) GetPubKey(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKey", reflect.TypeOf((*MockDatabase)(nil).GetPubKey), chain)
}

// GetTxOutWithHash mocks base method.
func (m *MockDatabase) GetTxOutWithHash(chain, hash string, isHashWithSig bool) *types.TxOutEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxOutWithHash", chain, hash, isHashWithSig)
	ret0, _ := ret[0].(*types.TxOutEntity)
	return ret0
}

// GetTxOutWithHash indicates an expected call of GetTxOutWithHash.
func (mr *MockDatabaseMockRecorder) GetTxOutWithHash(chain, hash, isHashWithSig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxOutWithHash", reflect.TypeOf((*MockDatabase)(nil).GetTxOutWithHash), chain, hash, isHashWithSig)
}

// Init mocks base method.
func (m *MockDatabase) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockDatabaseMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDatabase)(nil).Init))
}

// InsertContracts mocks base method.
func (m *MockDatabase) InsertContracts(contracts []*types.ContractEntity) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertContracts", contracts)
}

// InsertContracts indicates an expected call of InsertContracts.
func (mr *MockDatabaseMockRecorder) InsertContracts(contracts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertContracts", reflect.TypeOf((*MockDatabase)(nil).InsertContracts), contracts)
}

// InsertMempoolTxHash mocks base method.
func (m *MockDatabase) InsertMempoolTxHash(hash string, blockHeight int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertMempoolTxHash", hash, blockHeight)
}

// InsertMempoolTxHash indicates an expected call of InsertMempoolTxHash.
func (mr *MockDatabaseMockRecorder) InsertMempoolTxHash(hash, blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMempoolTxHash", reflect.TypeOf((*MockDatabase)(nil).InsertMempoolTxHash), hash, blockHeight)
}

// InsertTxOuts mocks base method.
func (m *MockDatabase) InsertTxOuts(txs []*types.TxOutEntity) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertTxOuts", txs)
}

// InsertTxOuts indicates an expected call of InsertTxOuts.
func (mr *MockDatabaseMockRecorder) InsertTxOuts(txs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTxOuts", reflect.TypeOf((*MockDatabase)(nil).InsertTxOuts), txs)
}

// IsChainKeyAddress mocks base method.
func (m *MockDatabase) IsChainKeyAddress(chain, address string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsChainKeyAddress", chain, address)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsChainKeyAddress indicates an expected call of IsChainKeyAddress.
func (mr *MockDatabaseMockRecorder) IsChainKeyAddress(chain, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsChainKeyAddress", reflect.TypeOf((*MockDatabase)(nil).IsChainKeyAddress), chain, address)
}

// IsContractDeployTx mocks base method.
func (m *MockDatabase) IsContractDeployTx(chain, hashWithoutSig string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsContractDeployTx", chain, hashWithoutSig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsContractDeployTx indicates an expected call of IsContractDeployTx.
func (mr *MockDatabaseMockRecorder) IsContractDeployTx(chain, hashWithoutSig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsContractDeployTx", reflect.TypeOf((*MockDatabase)(nil).IsContractDeployTx), chain, hashWithoutSig)
}

// IsKeyExisted mocks base method.
func (m *MockDatabase) IsKeyExisted(chain string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsKeyExisted", chain)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsKeyExisted indicates an expected call of IsKeyExisted.
func (mr *MockDatabaseMockRecorder) IsKeyExisted(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsKeyExisted", reflect.TypeOf((*MockDatabase)(nil).IsKeyExisted), chain)
}

// MempoolTxExisted mocks base method.
func (m *MockDatabase) MempoolTxExisted(hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MempoolTxExisted", hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MempoolTxExisted indicates an expected call of MempoolTxExisted.
func (mr *MockDatabaseMockRecorder) MempoolTxExisted(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MempoolTxExisted", reflect.TypeOf((*MockDatabase)(nil).MempoolTxExisted), hash)
}

// MempoolTxExistedRange mocks base method.
func (m *MockDatabase) MempoolTxExistedRange(hash string, minBlock, maxBlock int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MempoolTxExistedRange", hash, minBlock, maxBlock)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MempoolTxExistedRange indicates an expected call of MempoolTxExistedRange.
func (mr *MockDatabaseMockRecorder) MempoolTxExistedRange(hash, minBlock, maxBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MempoolTxExistedRange", reflect.TypeOf((*MockDatabase)(nil).MempoolTxExistedRange), hash, minBlock, maxBlock)
}

// UpdateContractAddress mocks base method.
func (m *MockDatabase) UpdateContractAddress(chain, hash, address string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateContractAddress", chain, hash, address)
}

// UpdateContractAddress indicates an expected call of UpdateContractAddress.
func (mr *MockDatabaseMockRecorder) UpdateContractAddress(chain, hash, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContractAddress", reflect.TypeOf((*MockDatabase)(nil).UpdateContractAddress), chain, hash, address)
}

// UpdateContractDeployTx mocks base method.
func (m *MockDatabase) UpdateContractDeployTx(chain, id, txHash string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateContractDeployTx", chain, id, txHash)
}

// UpdateContractDeployTx indicates an expected call of UpdateContractDeployTx.
func (mr *MockDatabaseMockRecorder) UpdateContractDeployTx(chain, id, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContractDeployTx", reflect.TypeOf((*MockDatabase)(nil).UpdateContractDeployTx), chain, id, txHash)
}

// UpdateContractsStatus mocks base method.
func (m *MockDatabase) UpdateContractsStatus(contracts []*types.ContractEntity, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContractsStatus", contracts, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContractsStatus indicates an expected call of UpdateContractsStatus.
func (mr *MockDatabaseMockRecorder) UpdateContractsStatus(contracts, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContractsStatus", reflect.TypeOf((*MockDatabase)(nil).UpdateContractsStatus), contracts, status)
}

// UpdateKeygenAddress mocks base method.
func (m *MockDatabase) UpdateKeygenAddress(chain, address string, pubKey []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateKeygenAddress", chain, address, pubKey)
}

// UpdateKeygenAddress indicates an expected call of UpdateKeygenAddress.
func (mr *MockDatabaseMockRecorder) UpdateKeygenAddress(chain, address, pubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKeygenAddress", reflect.TypeOf((*MockDatabase)(nil).UpdateKeygenAddress), chain, address, pubKey)
}

// UpdateKeygenStatus mocks base method.
func (m *MockDatabase) UpdateKeygenStatus(chain, status string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateKeygenStatus", chain, status)
}

// UpdateKeygenStatus indicates an expected call of UpdateKeygenStatus.
func (mr *MockDatabaseMockRecorder) UpdateKeygenStatus(chain, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKeygenStatus", reflect.TypeOf((*MockDatabase)(nil).UpdateKeygenStatus), chain, status)
}

// UpdateTxOutSig mocks base method.
func (m *MockDatabase) UpdateTxOutSig(chain, hashWithoutSign, hashWithSig string, sig []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTxOutSig", chain, hashWithoutSign, hashWithSig, sig)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTxOutSig indicates an expected call of UpdateTxOutSig.
func (mr *MockDatabaseMockRecorder) UpdateTxOutSig(chain, hashWithoutSign, hashWithSig, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTxOutSig", reflect.TypeOf((*MockDatabase)(nil).UpdateTxOutSig), chain, hashWithoutSign, hashWithSig, sig)
}

// UpdateTxOutStatus mocks base method.
func (m *MockDatabase) UpdateTxOutStatus(chain, hashWithoutSig string, status types.TxOutStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTxOutStatus", chain, hashWithoutSig, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTxOutStatus indicates an expected call of UpdateTxOutStatus.
func (mr *MockDatabaseMockRecorder) UpdateTxOutStatus(chain, hashWithoutSig, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTxOutStatus", reflect.TypeOf((*MockDatabase)(nil).UpdateTxOutStatus), chain, hashWithoutSig, status)
}
