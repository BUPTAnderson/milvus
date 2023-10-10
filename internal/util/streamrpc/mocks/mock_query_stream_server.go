// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// MockQueryStreamServer is an autogenerated mock type for the QueryNode_QueryStreamServer type
type MockQueryStreamServer struct {
	mock.Mock
}

type MockQueryStreamServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQueryStreamServer) EXPECT() *MockQueryStreamServer_Expecter {
	return &MockQueryStreamServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockQueryStreamServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockQueryStreamServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockQueryStreamServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockQueryStreamServer_Expecter) Context() *MockQueryStreamServer_Context_Call {
	return &MockQueryStreamServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockQueryStreamServer_Context_Call) Run(run func()) *MockQueryStreamServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockQueryStreamServer_Context_Call) Return(_a0 context.Context) *MockQueryStreamServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryStreamServer_Context_Call) RunAndReturn(run func() context.Context) *MockQueryStreamServer_Context_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockQueryStreamServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQueryStreamServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockQueryStreamServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockQueryStreamServer_Expecter) RecvMsg(m interface{}) *MockQueryStreamServer_RecvMsg_Call {
	return &MockQueryStreamServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockQueryStreamServer_RecvMsg_Call) Run(run func(m interface{})) *MockQueryStreamServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockQueryStreamServer_RecvMsg_Call) Return(_a0 error) *MockQueryStreamServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryStreamServer_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockQueryStreamServer_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockQueryStreamServer) Send(_a0 *internalpb.RetrieveResults) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internalpb.RetrieveResults) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQueryStreamServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockQueryStreamServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *internalpb.RetrieveResults
func (_e *MockQueryStreamServer_Expecter) Send(_a0 interface{}) *MockQueryStreamServer_Send_Call {
	return &MockQueryStreamServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockQueryStreamServer_Send_Call) Run(run func(_a0 *internalpb.RetrieveResults)) *MockQueryStreamServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*internalpb.RetrieveResults))
	})
	return _c
}

func (_c *MockQueryStreamServer_Send_Call) Return(_a0 error) *MockQueryStreamServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryStreamServer_Send_Call) RunAndReturn(run func(*internalpb.RetrieveResults) error) *MockQueryStreamServer_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockQueryStreamServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQueryStreamServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockQueryStreamServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockQueryStreamServer_Expecter) SendHeader(_a0 interface{}) *MockQueryStreamServer_SendHeader_Call {
	return &MockQueryStreamServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockQueryStreamServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockQueryStreamServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockQueryStreamServer_SendHeader_Call) Return(_a0 error) *MockQueryStreamServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryStreamServer_SendHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockQueryStreamServer_SendHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockQueryStreamServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQueryStreamServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockQueryStreamServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockQueryStreamServer_Expecter) SendMsg(m interface{}) *MockQueryStreamServer_SendMsg_Call {
	return &MockQueryStreamServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockQueryStreamServer_SendMsg_Call) Run(run func(m interface{})) *MockQueryStreamServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockQueryStreamServer_SendMsg_Call) Return(_a0 error) *MockQueryStreamServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryStreamServer_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockQueryStreamServer_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockQueryStreamServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQueryStreamServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockQueryStreamServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockQueryStreamServer_Expecter) SetHeader(_a0 interface{}) *MockQueryStreamServer_SetHeader_Call {
	return &MockQueryStreamServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockQueryStreamServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockQueryStreamServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockQueryStreamServer_SetHeader_Call) Return(_a0 error) *MockQueryStreamServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryStreamServer_SetHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockQueryStreamServer_SetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockQueryStreamServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockQueryStreamServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockQueryStreamServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockQueryStreamServer_Expecter) SetTrailer(_a0 interface{}) *MockQueryStreamServer_SetTrailer_Call {
	return &MockQueryStreamServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockQueryStreamServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockQueryStreamServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockQueryStreamServer_SetTrailer_Call) Return() *MockQueryStreamServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockQueryStreamServer_SetTrailer_Call) RunAndReturn(run func(metadata.MD)) *MockQueryStreamServer_SetTrailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQueryStreamServer creates a new instance of MockQueryStreamServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueryStreamServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueryStreamServer {
	mock := &MockQueryStreamServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
