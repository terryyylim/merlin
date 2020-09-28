// Code generated by mockery v2.0.0-alpha.14. DO NOT EDIT.

package mocks

import (
	context "context"

	mlp "github.com/gojek/merlin/mlp"
	mock "github.com/stretchr/testify/mock"
)

// APIClient is an autogenerated mock type for the APIClient type
type APIClient struct {
	mock.Mock
}

// CreateProject provides a mock function with given fields: ctx, project
func (_m *APIClient) CreateProject(ctx context.Context, project mlp.Project) (mlp.Project, error) {
	ret := _m.Called(ctx, project)

	var r0 mlp.Project
	if rf, ok := ret.Get(0).(func(context.Context, mlp.Project) mlp.Project); ok {
		r0 = rf(ctx, project)
	} else {
		r0 = ret.Get(0).(mlp.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mlp.Project) error); ok {
		r1 = rf(ctx, project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSecret provides a mock function with given fields: ctx, projectID, secret
func (_m *APIClient) CreateSecret(ctx context.Context, projectID int32, secret mlp.Secret) (mlp.Secret, error) {
	ret := _m.Called(ctx, projectID, secret)

	var r0 mlp.Secret
	if rf, ok := ret.Get(0).(func(context.Context, int32, mlp.Secret) mlp.Secret); ok {
		r0 = rf(ctx, projectID, secret)
	} else {
		r0 = ret.Get(0).(mlp.Secret)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32, mlp.Secret) error); ok {
		r1 = rf(ctx, projectID, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSecret provides a mock function with given fields: ctx, secretID, projectID
func (_m *APIClient) DeleteSecret(ctx context.Context, secretID int32, projectID int32) error {
	ret := _m.Called(ctx, secretID, projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, secretID, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPlainSecretByIDandProjectID provides a mock function with given fields: ctx, secretID, projectID
func (_m *APIClient) GetPlainSecretByIDandProjectID(ctx context.Context, secretID int32, projectID int32) (mlp.Secret, error) {
	ret := _m.Called(ctx, secretID, projectID)

	var r0 mlp.Secret
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) mlp.Secret); ok {
		r0 = rf(ctx, secretID, projectID)
	} else {
		r0 = ret.Get(0).(mlp.Secret)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32, int32) error); ok {
		r1 = rf(ctx, secretID, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlainSecretByNameAndProjectID provides a mock function with given fields: ctx, secretName, projectID
func (_m *APIClient) GetPlainSecretByNameAndProjectID(ctx context.Context, secretName string, projectID int32) (mlp.Secret, error) {
	ret := _m.Called(ctx, secretName, projectID)

	var r0 mlp.Secret
	if rf, ok := ret.Get(0).(func(context.Context, string, int32) mlp.Secret); ok {
		r0 = rf(ctx, secretName, projectID)
	} else {
		r0 = ret.Get(0).(mlp.Secret)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int32) error); ok {
		r1 = rf(ctx, secretName, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectByID provides a mock function with given fields: ctx, projectID
func (_m *APIClient) GetProjectByID(ctx context.Context, projectID int32) (mlp.Project, error) {
	ret := _m.Called(ctx, projectID)

	var r0 mlp.Project
	if rf, ok := ret.Get(0).(func(context.Context, int32) mlp.Project); ok {
		r0 = rf(ctx, projectID)
	} else {
		r0 = ret.Get(0).(mlp.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectByName provides a mock function with given fields: ctx, projectName
func (_m *APIClient) GetProjectByName(ctx context.Context, projectName string) (mlp.Project, error) {
	ret := _m.Called(ctx, projectName)

	var r0 mlp.Project
	if rf, ok := ret.Get(0).(func(context.Context, string) mlp.Project); ok {
		r0 = rf(ctx, projectName)
	} else {
		r0 = ret.Get(0).(mlp.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecretByIDandProjectID provides a mock function with given fields: ctx, secretID, projectID
func (_m *APIClient) GetSecretByIDandProjectID(ctx context.Context, secretID int32, projectID int32) (mlp.Secret, error) {
	ret := _m.Called(ctx, secretID, projectID)

	var r0 mlp.Secret
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) mlp.Secret); ok {
		r0 = rf(ctx, secretID, projectID)
	} else {
		r0 = ret.Get(0).(mlp.Secret)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32, int32) error); ok {
		r1 = rf(ctx, secretID, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecretByNameAndProjectID provides a mock function with given fields: ctx, secretName, projectID
func (_m *APIClient) GetSecretByNameAndProjectID(ctx context.Context, secretName string, projectID int32) (mlp.Secret, error) {
	ret := _m.Called(ctx, secretName, projectID)

	var r0 mlp.Secret
	if rf, ok := ret.Get(0).(func(context.Context, string, int32) mlp.Secret); ok {
		r0 = rf(ctx, secretName, projectID)
	} else {
		r0 = ret.Get(0).(mlp.Secret)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int32) error); ok {
		r1 = rf(ctx, secretName, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: ctx, projectName
func (_m *APIClient) ListProjects(ctx context.Context, projectName string) (mlp.Projects, error) {
	ret := _m.Called(ctx, projectName)

	var r0 mlp.Projects
	if rf, ok := ret.Get(0).(func(context.Context, string) mlp.Projects); ok {
		r0 = rf(ctx, projectName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mlp.Projects)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecrets provides a mock function with given fields: ctx, projectID
func (_m *APIClient) ListSecrets(ctx context.Context, projectID int32) (mlp.Secrets, error) {
	ret := _m.Called(ctx, projectID)

	var r0 mlp.Secrets
	if rf, ok := ret.Get(0).(func(context.Context, int32) mlp.Secrets); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mlp.Secrets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProject provides a mock function with given fields: ctx, project
func (_m *APIClient) UpdateProject(ctx context.Context, project mlp.Project) (mlp.Project, error) {
	ret := _m.Called(ctx, project)

	var r0 mlp.Project
	if rf, ok := ret.Get(0).(func(context.Context, mlp.Project) mlp.Project); ok {
		r0 = rf(ctx, project)
	} else {
		r0 = ret.Get(0).(mlp.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mlp.Project) error); ok {
		r1 = rf(ctx, project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecret provides a mock function with given fields: ctx, projectID, secret
func (_m *APIClient) UpdateSecret(ctx context.Context, projectID int32, secret mlp.Secret) (mlp.Secret, error) {
	ret := _m.Called(ctx, projectID, secret)

	var r0 mlp.Secret
	if rf, ok := ret.Get(0).(func(context.Context, int32, mlp.Secret) mlp.Secret); ok {
		r0 = rf(ctx, projectID, secret)
	} else {
		r0 = ret.Get(0).(mlp.Secret)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32, mlp.Secret) error); ok {
		r1 = rf(ctx, projectID, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
