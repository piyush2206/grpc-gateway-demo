// Code generated by mockery v1.0.0. DO NOT EDIT.
package http2

import mock "github.com/stretchr/testify/mock"

// mockHeadersEnder is an autogenerated mock type for the headersEnder type
type mockHeadersEnder struct {
	mock.Mock
}

// HeadersEnded provides a mock function with given fields:
func (_m *mockHeadersEnder) HeadersEnded() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
