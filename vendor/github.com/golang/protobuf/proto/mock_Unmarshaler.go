// Code generated by mockery v1.0.0. DO NOT EDIT.
package proto

import mock "github.com/stretchr/testify/mock"

// MockUnmarshaler is an autogenerated mock type for the Unmarshaler type
type MockUnmarshaler struct {
	mock.Mock
}

// Unmarshal provides a mock function with given fields: _a0
func (_m *MockUnmarshaler) Unmarshal(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
