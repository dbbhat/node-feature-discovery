// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

// Re-generate by running 'make mock'

package source

import mock "github.com/stretchr/testify/mock"

// MockFeatureSource is an autogenerated mock type for the FeatureSource type
type MockFeatureSource struct {
	mock.Mock
}

// Discover provides a mock function with given fields:
func (_m *MockFeatureSource) Discover() (Features, error) {
	ret := _m.Called()

	var r0 Features
	if rf, ok := ret.Get(0).(func() Features); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Features)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfig provides a mock function with given fields:
func (_m *MockFeatureSource) GetConfig() Config {
	ret := _m.Called()

	var r0 Config
	if rf, ok := ret.Get(0).(func() Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Config)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *MockFeatureSource) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewConfig provides a mock function with given fields:
func (_m *MockFeatureSource) NewConfig() Config {
	ret := _m.Called()

	var r0 Config
	if rf, ok := ret.Get(0).(func() Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Config)
		}
	}

	return r0
}

// SetConfig provides a mock function with given fields: _a0
func (_m *MockFeatureSource) SetConfig(_a0 Config) {
	_m.Called(_a0)
}
