// Code generated by mockery v1.0.0. DO NOT EDIT.
package student

import mock "github.com/stretchr/testify/mock"

// mockSchoolStudent is an autogenerated mock type for the schoolStudent type
type mockSchoolStudent struct {
	mock.Mock
}

// getName provides a mock function with given fields:
func (_m *mockSchoolStudent) getName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getSubjects provides a mock function with given fields:
func (_m *mockSchoolStudent) getSubjects() []*Subject {
	ret := _m.Called()

	var r0 []*Subject
	if rf, ok := ret.Get(0).(func() []*Subject); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Subject)
		}
	}

	return r0
}

// getTotalScore provides a mock function with given fields:
func (_m *mockSchoolStudent) getTotalScore() (float32, float32) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 float32
	if rf, ok := ret.Get(1).(func() float32); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float32)
	}

	return r0, r1
}
