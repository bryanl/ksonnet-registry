// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import repository "github.com/bryanl/ksonnet-registry/repository"

// PackageRepository is an autogenerated mock type for the PackageRepository type
type PackageRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ns, pkg
func (_m *PackageRepository) Create(ns string, pkg string) (repository.Package, error) {
	ret := _m.Called(ns, pkg)

	var r0 repository.Package
	if rf, ok := ret.Get(0).(func(string, string) repository.Package); ok {
		r0 = rf(ns, pkg)
	} else {
		r0 = ret.Get(0).(repository.Package)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(ns, pkg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ns
func (_m *PackageRepository) List(ns string) ([]repository.Package, error) {
	ret := _m.Called(ns)

	var r0 []repository.Package
	if rf, ok := ret.Get(0).(func(string) []repository.Package); ok {
		r0 = rf(ns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Package)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Retrieve provides a mock function with given fields: ns, pkg
func (_m *PackageRepository) Retrieve(ns string, pkg string) (repository.Package, error) {
	ret := _m.Called(ns, pkg)

	var r0 repository.Package
	if rf, ok := ret.Get(0).(func(string, string) repository.Package); ok {
		r0 = rf(ns, pkg)
	} else {
		r0 = ret.Get(0).(repository.Package)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(ns, pkg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
