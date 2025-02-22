// Code generated by mockery v2.50.0. DO NOT EDIT.

package testhelpers

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockResponseWriter is an autogenerated mock type for the ResponseWriter type
type MockResponseWriter struct {
	mock.Mock
}

type MockResponseWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResponseWriter) EXPECT() *MockResponseWriter_Expecter {
	return &MockResponseWriter_Expecter{mock: &_m.Mock}
}

// Header provides a mock function with no fields
func (_m *MockResponseWriter) Header() http.Header {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 http.Header
	if rf, ok := ret.Get(0).(func() http.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Header)
		}
	}

	return r0
}

// MockResponseWriter_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockResponseWriter_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockResponseWriter_Expecter) Header() *MockResponseWriter_Header_Call {
	return &MockResponseWriter_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *MockResponseWriter_Header_Call) Run(run func()) *MockResponseWriter_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResponseWriter_Header_Call) Return(_a0 http.Header) *MockResponseWriter_Header_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockResponseWriter_Header_Call) RunAndReturn(run func() http.Header) *MockResponseWriter_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: _a0
func (_m *MockResponseWriter) Write(_a0 []byte) (int, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResponseWriter_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockResponseWriter_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - _a0 []byte
func (_e *MockResponseWriter_Expecter) Write(_a0 interface{}) *MockResponseWriter_Write_Call {
	return &MockResponseWriter_Write_Call{Call: _e.mock.On("Write", _a0)}
}

func (_c *MockResponseWriter_Write_Call) Run(run func(_a0 []byte)) *MockResponseWriter_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockResponseWriter_Write_Call) Return(_a0 int, _a1 error) *MockResponseWriter_Write_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResponseWriter_Write_Call) RunAndReturn(run func([]byte) (int, error)) *MockResponseWriter_Write_Call {
	_c.Call.Return(run)
	return _c
}

// WriteHeader provides a mock function with given fields: statusCode
func (_m *MockResponseWriter) WriteHeader(statusCode int) {
	_m.Called(statusCode)
}

// MockResponseWriter_WriteHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteHeader'
type MockResponseWriter_WriteHeader_Call struct {
	*mock.Call
}

// WriteHeader is a helper method to define mock.On call
//   - statusCode int
func (_e *MockResponseWriter_Expecter) WriteHeader(statusCode interface{}) *MockResponseWriter_WriteHeader_Call {
	return &MockResponseWriter_WriteHeader_Call{Call: _e.mock.On("WriteHeader", statusCode)}
}

func (_c *MockResponseWriter_WriteHeader_Call) Run(run func(statusCode int)) *MockResponseWriter_WriteHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockResponseWriter_WriteHeader_Call) Return() *MockResponseWriter_WriteHeader_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockResponseWriter_WriteHeader_Call) RunAndReturn(run func(int)) *MockResponseWriter_WriteHeader_Call {
	_c.Run(run)
	return _c
}

// NewMockResponseWriter creates a new instance of MockResponseWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResponseWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResponseWriter {
	mock := &MockResponseWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
