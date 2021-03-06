// Code generated by mockery v1.0.0. DO NOT EDIT.
package http2

import mock "github.com/stretchr/testify/mock"
import tls "crypto/tls"

// mockConnectionStater is an autogenerated mock type for the connectionStater type
type mockConnectionStater struct {
	mock.Mock
}

// ConnectionState provides a mock function with given fields:
func (_m *mockConnectionStater) ConnectionState() tls.ConnectionState {
	ret := _m.Called()

	var r0 tls.ConnectionState
	if rf, ok := ret.Get(0).(func() tls.ConnectionState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(tls.ConnectionState)
	}

	return r0
}
