// Code generated by mockery v2.32.0. DO NOT EDIT.

package synchronizer

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/hoaleee/go-ethereum/core/types"
)

// poolMock is an autogenerated mock type for the poolInterface type
type poolMock struct {
	mock.Mock
}

// DeleteReorgedTransactions provides a mock function with given fields: ctx, txs
func (_m *poolMock) DeleteReorgedTransactions(ctx context.Context, txs []*types.Transaction) error {
	ret := _m.Called(ctx, txs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*types.Transaction) error); ok {
		r0 = rf(ctx, txs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreTx provides a mock function with given fields: ctx, tx, ip, isWIP
func (_m *poolMock) StoreTx(ctx context.Context, tx types.Transaction, ip string, isWIP bool) error {
	ret := _m.Called(ctx, tx, ip, isWIP)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Transaction, string, bool) error); ok {
		r0 = rf(ctx, tx, ip, isWIP)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newPoolMock creates a new instance of poolMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newPoolMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *poolMock {
	mock := &poolMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
