// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import tenantmapping "github.com/kyma-incubator/compass/components/director/internal/tenantmapping"

// TenantAndScopesForUserProvider is an autogenerated mock type for the TenantAndScopesForUserProvider type
type TenantAndScopesForUserProvider struct {
	mock.Mock
}

// GetTenantAndScopes provides a mock function with given fields: reqData, authID
func (_m *TenantAndScopesForUserProvider) GetTenantAndScopes(reqData tenantmapping.ReqData, authID string) (string, string, error) {
	ret := _m.Called(reqData, authID)

	var r0 string
	if rf, ok := ret.Get(0).(func(tenantmapping.ReqData, string) string); ok {
		r0 = rf(reqData, authID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(tenantmapping.ReqData, string) string); ok {
		r1 = rf(reqData, authID)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(tenantmapping.ReqData, string) error); ok {
		r2 = rf(reqData, authID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
