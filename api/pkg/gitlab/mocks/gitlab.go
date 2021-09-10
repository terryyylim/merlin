// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	gitlab "github.com/gojek/merlin/pkg/gitlab"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateFile provides a mock function with given fields: opt
func (_m *Client) CreateFile(opt gitlab.CreateFileOptions) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(gitlab.CreateFileOptions) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFileContent provides a mock function with given fields: opt
func (_m *Client) GetFileContent(opt gitlab.GetFileContentOptions) (string, error) {
	ret := _m.Called(opt)

	var r0 string
	if rf, ok := ret.Get(0).(func(gitlab.GetFileContentOptions) string); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.GetFileContentOptions) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFile provides a mock function with given fields: opt
func (_m *Client) UpdateFile(opt gitlab.UpdateFileOptions) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(gitlab.UpdateFileOptions) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}