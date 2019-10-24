// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// SystemAuthService is an autogenerated mock type for the SystemAuthService type
type SystemAuthService struct {
	mock.Mock
}

// DeleteByIDForObject provides a mock function with given fields: ctx, objectType, authID
func (_m *SystemAuthService) DeleteByIDForObject(ctx context.Context, objectType model.SystemAuthReferenceObjectType, authID string) error {
	ret := _m.Called(ctx, objectType, authID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SystemAuthReferenceObjectType, string) error); ok {
		r0 = rf(ctx, objectType, authID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByIDForObject provides a mock function with given fields: ctx, objectType, authID
func (_m *SystemAuthService) GetByIDForObject(ctx context.Context, objectType model.SystemAuthReferenceObjectType, authID string) (*model.SystemAuth, error) {
	ret := _m.Called(ctx, objectType, authID)

	var r0 *model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, model.SystemAuthReferenceObjectType, string) *model.SystemAuth); ok {
		r0 = rf(ctx, objectType, authID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SystemAuthReferenceObjectType, string) error); ok {
		r1 = rf(ctx, objectType, authID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGlobal provides a mock function with given fields: ctx, id
func (_m *SystemAuthService) GetGlobal(ctx context.Context, id string) (*model.SystemAuth, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.SystemAuth); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
