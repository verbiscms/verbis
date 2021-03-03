// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"

	params "github.com/ainsleyclark/verbis/api/helpers/params"
)

// PostsRepository is an autogenerated mock type for the PostsRepository type
type PostsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: p
func (_m *PostsRepository) Create(p *domain.PostCreate) (domain.PostDatum, error) {
	ret := _m.Called(p)

	var r0 domain.PostDatum
	if rf, ok := ret.Get(0).(func(*domain.PostCreate) domain.PostDatum); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(domain.PostDatum)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.PostCreate) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *PostsRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: id
func (_m *PostsRepository) Exists(id int) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ExistsBySlug provides a mock function with given fields: slug
func (_m *PostsRepository) ExistsBySlug(slug string) bool {
	ret := _m.Called(slug)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: meta, layout, resource, status
func (_m *PostsRepository) Get(meta params.Params, layout bool, resource string, status string) (domain.PostData, int, error) {
	ret := _m.Called(meta, layout, resource, status)

	var r0 domain.PostData
	if rf, ok := ret.Get(0).(func(params.Params, bool, string, string) domain.PostData); ok {
		r0 = rf(meta, layout, resource, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.PostData)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(params.Params, bool, string, string) int); ok {
		r1 = rf(meta, layout, resource, status)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(params.Params, bool, string, string) error); ok {
		r2 = rf(meta, layout, resource, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id, layout
func (_m *PostsRepository) GetByID(id int, layout bool) (domain.PostDatum, error) {
	ret := _m.Called(id, layout)

	var r0 domain.PostDatum
	if rf, ok := ret.Get(0).(func(int, bool) domain.PostDatum); ok {
		r0 = rf(id, layout)
	} else {
		r0 = ret.Get(0).(domain.PostDatum)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, bool) error); ok {
		r1 = rf(id, layout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySlug provides a mock function with given fields: slug
func (_m *PostsRepository) GetBySlug(slug string) (domain.PostDatum, error) {
	ret := _m.Called(slug)

	var r0 domain.PostDatum
	if rf, ok := ret.Get(0).(func(string) domain.PostDatum); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(domain.PostDatum)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Total provides a mock function with given fields:
func (_m *PostsRepository) Total() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: p
func (_m *PostsRepository) Update(p *domain.PostCreate) (domain.PostDatum, error) {
	ret := _m.Called(p)

	var r0 domain.PostDatum
	if rf, ok := ret.Get(0).(func(*domain.PostCreate) domain.PostDatum); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(domain.PostDatum)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.PostCreate) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
