// Code generated by mockery v2.0.0-alpha.14. DO NOT EDIT.

package mocks

import (
	mlp "github.com/gojek/merlin/mlp"
	mock "github.com/stretchr/testify/mock"

	models "github.com/gojek/merlin/models"
)

// ImageBuilder is an autogenerated mock type for the ImageBuilder type
type ImageBuilder struct {
	mock.Mock
}

// BuildImage provides a mock function with given fields: project, model, version
func (_m *ImageBuilder) BuildImage(project mlp.Project, model *models.Model, version *models.Version) (string, error) {
	ret := _m.Called(project, model, version)

	var r0 string
	if rf, ok := ret.Get(0).(func(mlp.Project, *models.Model, *models.Version) string); ok {
		r0 = rf(project, model, version)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mlp.Project, *models.Model, *models.Version) error); ok {
		r1 = rf(project, model, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainers provides a mock function with given fields: project, model, version
func (_m *ImageBuilder) GetContainers(project mlp.Project, model *models.Model, version *models.Version) ([]*models.Container, error) {
	ret := _m.Called(project, model, version)

	var r0 []*models.Container
	if rf, ok := ret.Get(0).(func(mlp.Project, *models.Model, *models.Version) []*models.Container); ok {
		r0 = rf(project, model, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mlp.Project, *models.Model, *models.Version) error); ok {
		r1 = rf(project, model, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
