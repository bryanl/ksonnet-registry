// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import repository "github.com/bryanl/ksonnet-registry/repository"

// NamespaceRepository is an autogenerated mock type for the NamespaceRepository type
type NamespaceRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ns
func (_m *NamespaceRepository) Create(ns string) (repository.Namespace, error) {
	ret := _m.Called(ns)

	var r0 repository.Namespace
	if rf, ok := ret.Get(0).(func(string) repository.Namespace); ok {
		r0 = rf(ns)
	} else {
		r0 = ret.Get(0).(repository.Namespace)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *NamespaceRepository) List() ([]repository.Namespace, error) {
	ret := _m.Called()

	var r0 []repository.Namespace
	if rf, ok := ret.Get(0).(func() []repository.Namespace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Namespace)
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
