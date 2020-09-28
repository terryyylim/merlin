// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/gojek/merlin/models"
	mock "github.com/stretchr/testify/mock"
)

// ModelEndpointsService is an autogenerated mock type for the ModelEndpointsService type
type ModelEndpointsService struct {
	mock.Mock
}

// DeployEndpoint provides a mock function with given fields: ctx, model, endpoint
func (_m *ModelEndpointsService) DeployEndpoint(ctx context.Context, model *models.Model, endpoint *models.ModelEndpoint) (*models.ModelEndpoint, error) {
	ret := _m.Called(ctx, model, endpoint)

	var r0 *models.ModelEndpoint
	if rf, ok := ret.Get(0).(func(context.Context, *models.Model, *models.ModelEndpoint) *models.ModelEndpoint); ok {
		r0 = rf(ctx, model, endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ModelEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Model, *models.ModelEndpoint) error); ok {
		r1 = rf(ctx, model, endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, id
func (_m *ModelEndpointsService) FindById(ctx context.Context, id models.Id) (*models.ModelEndpoint, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.ModelEndpoint
	if rf, ok := ret.Get(0).(func(context.Context, models.Id) *models.ModelEndpoint); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ModelEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Id) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModelEndpoints provides a mock function with given fields: ctx, modelId
func (_m *ModelEndpointsService) ListModelEndpoints(ctx context.Context, modelId models.Id) ([]*models.ModelEndpoint, error) {
	ret := _m.Called(ctx, modelId)

	var r0 []*models.ModelEndpoint
	if rf, ok := ret.Get(0).(func(context.Context, models.Id) []*models.ModelEndpoint); ok {
		r0 = rf(ctx, modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ModelEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Id) error); ok {
		r1 = rf(ctx, modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModelEndpointsInProject provides a mock function with given fields: ctx, projectId, region
func (_m *ModelEndpointsService) ListModelEndpointsInProject(ctx context.Context, projectId models.Id, region string) ([]*models.ModelEndpoint, error) {
	ret := _m.Called(ctx, projectId, region)

	var r0 []*models.ModelEndpoint
	if rf, ok := ret.Get(0).(func(context.Context, models.Id, string) []*models.ModelEndpoint); ok {
		r0 = rf(ctx, projectId, region)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ModelEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Id, string) error); ok {
		r1 = rf(ctx, projectId, region)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, endpoint
func (_m *ModelEndpointsService) Save(ctx context.Context, endpoint *models.ModelEndpoint) (*models.ModelEndpoint, error) {
	ret := _m.Called(ctx, endpoint)

	var r0 *models.ModelEndpoint
	if rf, ok := ret.Get(0).(func(context.Context, *models.ModelEndpoint) *models.ModelEndpoint); ok {
		r0 = rf(ctx, endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ModelEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.ModelEndpoint) error); ok {
		r1 = rf(ctx, endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UndeployEndpoint provides a mock function with given fields: ctx, model, endpoint
func (_m *ModelEndpointsService) UndeployEndpoint(ctx context.Context, model *models.Model, endpoint *models.ModelEndpoint) (*models.ModelEndpoint, error) {
	ret := _m.Called(ctx, model, endpoint)

	var r0 *models.ModelEndpoint
	if rf, ok := ret.Get(0).(func(context.Context, *models.Model, *models.ModelEndpoint) *models.ModelEndpoint); ok {
		r0 = rf(ctx, model, endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ModelEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Model, *models.ModelEndpoint) error); ok {
		r1 = rf(ctx, model, endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEndpoint provides a mock function with given fields: ctx, model, endpoint
func (_m *ModelEndpointsService) UpdateEndpoint(ctx context.Context, model *models.Model, endpoint *models.ModelEndpoint) (*models.ModelEndpoint, error) {
	ret := _m.Called(ctx, model, endpoint)

	var r0 *models.ModelEndpoint
	if rf, ok := ret.Get(0).(func(context.Context, *models.Model, *models.ModelEndpoint) *models.ModelEndpoint); ok {
		r0 = rf(ctx, model, endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ModelEndpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Model, *models.ModelEndpoint) error); ok {
		r1 = rf(ctx, model, endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
