// Code generated by mockery v1.0.0. DO NOT EDIT.

package extractor

import mock "github.com/stretchr/testify/mock"

// MockExtractor is an autogenerated mock type for the Extractor type
type MockExtractor struct {
	mock.Mock
}

// ExtractMRIIDs provides a mock function with given fields: startingPoint, matcher
func (_m *MockExtractor) ExtractMRIIDs(startingPoint string, matcher string) ([]int, error) {
	ret := _m.Called(startingPoint, matcher)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string, string) []int); ok {
		r0 = rf(startingPoint, matcher)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(startingPoint, matcher)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
