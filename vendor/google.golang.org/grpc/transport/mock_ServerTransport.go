// Code generated by mockery v1.0.0. DO NOT EDIT.
package transport

import context "context"
import metadata "google.golang.org/grpc/metadata"
import mock "github.com/stretchr/testify/mock"
import net "net"
import status "google.golang.org/grpc/status"

// MockServerTransport is an autogenerated mock type for the ServerTransport type
type MockServerTransport struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockServerTransport) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Drain provides a mock function with given fields:
func (_m *MockServerTransport) Drain() {
	_m.Called()
}

// HandleStreams provides a mock function with given fields: _a0, _a1
func (_m *MockServerTransport) HandleStreams(_a0 func(*Stream), _a1 func(context.Context, string) context.Context) {
	_m.Called(_a0, _a1)
}

// IncrMsgRecv provides a mock function with given fields:
func (_m *MockServerTransport) IncrMsgRecv() {
	_m.Called()
}

// IncrMsgSent provides a mock function with given fields:
func (_m *MockServerTransport) IncrMsgSent() {
	_m.Called()
}

// RemoteAddr provides a mock function with given fields:
func (_m *MockServerTransport) RemoteAddr() net.Addr {
	ret := _m.Called()

	var r0 net.Addr
	if rf, ok := ret.Get(0).(func() net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Addr)
		}
	}

	return r0
}

// Write provides a mock function with given fields: s, hdr, data, opts
func (_m *MockServerTransport) Write(s *Stream, hdr []byte, data []byte, opts *Options) error {
	ret := _m.Called(s, hdr, data, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Stream, []byte, []byte, *Options) error); ok {
		r0 = rf(s, hdr, data, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteHeader provides a mock function with given fields: s, md
func (_m *MockServerTransport) WriteHeader(s *Stream, md metadata.MD) error {
	ret := _m.Called(s, md)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Stream, metadata.MD) error); ok {
		r0 = rf(s, md)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteStatus provides a mock function with given fields: s, st
func (_m *MockServerTransport) WriteStatus(s *Stream, st *status.Status) error {
	ret := _m.Called(s, st)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Stream, *status.Status) error); ok {
		r0 = rf(s, st)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
