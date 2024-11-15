// Code generated by mockery v2.46.0. DO NOT EDIT.

package db

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// GetAccounts provides a mock function with given fields: ctx
func (_m *MockQuerier) GetAccounts(ctx context.Context) ([]Account, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAccounts")
	}

	var r0 []Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Account, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Account); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccounts'
type MockQuerier_GetAccounts_Call struct {
	*mock.Call
}

// GetAccounts is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerier_Expecter) GetAccounts(ctx interface{}) *MockQuerier_GetAccounts_Call {
	return &MockQuerier_GetAccounts_Call{Call: _e.mock.On("GetAccounts", ctx)}
}

func (_c *MockQuerier_GetAccounts_Call) Run(run func(ctx context.Context)) *MockQuerier_GetAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerier_GetAccounts_Call) Return(_a0 []Account, _a1 error) *MockQuerier_GetAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetAccounts_Call) RunAndReturn(run func(context.Context) ([]Account, error)) *MockQuerier_GetAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAccount provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, UpdateAccountParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerier_UpdateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAccount'
type MockQuerier_UpdateAccount_Call struct {
	*mock.Call
}

// UpdateAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - arg UpdateAccountParams
func (_e *MockQuerier_Expecter) UpdateAccount(ctx interface{}, arg interface{}) *MockQuerier_UpdateAccount_Call {
	return &MockQuerier_UpdateAccount_Call{Call: _e.mock.On("UpdateAccount", ctx, arg)}
}

func (_c *MockQuerier_UpdateAccount_Call) Run(run func(ctx context.Context, arg UpdateAccountParams)) *MockQuerier_UpdateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(UpdateAccountParams))
	})
	return _c
}

func (_c *MockQuerier_UpdateAccount_Call) Return(_a0 error) *MockQuerier_UpdateAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerier_UpdateAccount_Call) RunAndReturn(run func(context.Context, UpdateAccountParams) error) *MockQuerier_UpdateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
