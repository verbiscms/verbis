// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// CreateBucket provides a mock function with given fields: provider, name
func (_m *Provider) CreateBucket(provider domain.StorageProvider, name string) error {
	ret := _m.Called(provider, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.StorageProvider, string) error); ok {
		r0 = rf(provider, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Provider) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBucket provides a mock function with given fields: provider, name
func (_m *Provider) DeleteBucket(provider domain.StorageProvider, name string) error {
	ret := _m.Called(provider, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.StorageProvider, string) error); ok {
		r0 = rf(provider, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: name
func (_m *Provider) Exists(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: url
func (_m *Provider) Find(url string) ([]byte, domain.File, error) {
	ret := _m.Called(url)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 domain.File
	if rf, ok := ret.Get(1).(func(string) domain.File); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Get(1).(domain.File)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(url)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Info provides a mock function with given fields:
func (_m *Provider) Info() (domain.StorageConfiguration, error) {
	ret := _m.Called()

	var r0 domain.StorageConfiguration
	if rf, ok := ret.Get(0).(func() domain.StorageConfiguration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.StorageConfiguration)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBuckets provides a mock function with given fields: provider
func (_m *Provider) ListBuckets(provider domain.StorageProvider) (domain.Buckets, error) {
	ret := _m.Called(provider)

	var r0 domain.Buckets
	if rf, ok := ret.Get(0).(func(domain.StorageProvider) domain.Buckets); ok {
		r0 = rf(provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Buckets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.StorageProvider) error); ok {
		r1 = rf(provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields: from, to
func (_m *Provider) Migrate(from domain.StorageChange, to domain.StorageChange) (int, error) {
	ret := _m.Called(from, to)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.StorageChange, domain.StorageChange) int); ok {
		r0 = rf(from, to)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.StorageChange, domain.StorageChange) error); ok {
		r1 = rf(from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: upload
func (_m *Provider) Upload(upload domain.Upload) (domain.File, error) {
	ret := _m.Called(upload)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(domain.Upload) domain.File); ok {
		r0 = rf(upload)
	} else {
		r0 = ret.Get(0).(domain.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Upload) error); ok {
		r1 = rf(upload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
