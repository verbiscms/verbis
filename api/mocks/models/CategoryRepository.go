// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"

	params "github.com/ainsleyclark/verbis/api/helpers/params"
)

// CategoryRepository is an autogenerated mock type for the CategoryRepository type
type CategoryRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *CategoryRepository) Create(c *domain.Category) (domain.Category, error) {
	ret := _m.Called(c)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(*domain.Category) domain.Category); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Category) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *CategoryRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePostCategories provides a mock function with given fields: id
func (_m *CategoryRepository) DeletePostCategories(id int) error {
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
func (_m *CategoryRepository) Exists(id int) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: meta, resource
func (_m *CategoryRepository) Get(meta params.Params, resource string) (domain.Categories, int, error) {
	ret := _m.Called(meta, resource)

	var r0 domain.Categories
	if rf, ok := ret.Get(0).(func(params.Params, string) domain.Categories); ok {
		r0 = rf(meta, resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Categories)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(params.Params, string) int); ok {
		r1 = rf(meta, resource)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(params.Params, string) error); ok {
		r2 = rf(meta, resource)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id
func (_m *CategoryRepository) GetByID(id int) (domain.Category, error) {
	ret := _m.Called(id)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(int) domain.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *CategoryRepository) GetByName(name string) (domain.Category, error) {
	ret := _m.Called(name)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(string) domain.Category); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPost provides a mock function with given fields: pageID
func (_m *CategoryRepository) GetByPost(pageID int) (*domain.Category, error) {
	ret := _m.Called(pageID)

	var r0 *domain.Category
	if rf, ok := ret.Get(0).(func(int) *domain.Category); ok {
		r0 = rf(pageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(pageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySlug provides a mock function with given fields: slug
func (_m *CategoryRepository) GetBySlug(slug string) (domain.Category, error) {
	ret := _m.Called(slug)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(string) domain.Category); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParent provides a mock function with given fields: id
func (_m *CategoryRepository) GetParent(id int) (domain.Category, error) {
	ret := _m.Called(id)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(int) domain.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertPostCategory provides a mock function with given fields: postID, categoryID
func (_m *CategoryRepository) InsertPostCategory(postID int, categoryID *int) error {
	ret := _m.Called(postID, categoryID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *int) error); ok {
		r0 = rf(postID, categoryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: c
func (_m *CategoryRepository) Update(c *domain.Category) (domain.Category, error) {
	ret := _m.Called(c)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(*domain.Category) domain.Category); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Category) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
