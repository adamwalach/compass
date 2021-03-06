// Code generated by counterfeiter. DO NOT EDIT.
package controllersfakes

import (
	"context"
	"sync"

	"github.com/kyma-incubator/compass/components/operations-controller/controllers"
	directora "github.com/kyma-incubator/compass/components/operations-controller/internal/director"
	"github.com/kyma-incubator/compass/components/system-broker/pkg/director"
)

type FakeDirectorClient struct {
	FetchApplicationStub        func(context.Context, string) (*director.ApplicationOutput, error)
	fetchApplicationMutex       sync.RWMutex
	fetchApplicationArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	fetchApplicationReturns struct {
		result1 *director.ApplicationOutput
		result2 error
	}
	fetchApplicationReturnsOnCall map[int]struct {
		result1 *director.ApplicationOutput
		result2 error
	}
	UpdateOperationStub        func(context.Context, *directora.Request) error
	updateOperationMutex       sync.RWMutex
	updateOperationArgsForCall []struct {
		arg1 context.Context
		arg2 *directora.Request
	}
	updateOperationReturns struct {
		result1 error
	}
	updateOperationReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirectorClient) FetchApplication(arg1 context.Context, arg2 string) (*director.ApplicationOutput, error) {
	fake.fetchApplicationMutex.Lock()
	ret, specificReturn := fake.fetchApplicationReturnsOnCall[len(fake.fetchApplicationArgsForCall)]
	fake.fetchApplicationArgsForCall = append(fake.fetchApplicationArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("FetchApplication", []interface{}{arg1, arg2})
	fake.fetchApplicationMutex.Unlock()
	if fake.FetchApplicationStub != nil {
		return fake.FetchApplicationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDirectorClient) FetchApplicationCallCount() int {
	fake.fetchApplicationMutex.RLock()
	defer fake.fetchApplicationMutex.RUnlock()
	return len(fake.fetchApplicationArgsForCall)
}

func (fake *FakeDirectorClient) FetchApplicationCalls(stub func(context.Context, string) (*director.ApplicationOutput, error)) {
	fake.fetchApplicationMutex.Lock()
	defer fake.fetchApplicationMutex.Unlock()
	fake.FetchApplicationStub = stub
}

func (fake *FakeDirectorClient) FetchApplicationArgsForCall(i int) (context.Context, string) {
	fake.fetchApplicationMutex.RLock()
	defer fake.fetchApplicationMutex.RUnlock()
	argsForCall := fake.fetchApplicationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDirectorClient) FetchApplicationReturns(result1 *director.ApplicationOutput, result2 error) {
	fake.fetchApplicationMutex.Lock()
	defer fake.fetchApplicationMutex.Unlock()
	fake.FetchApplicationStub = nil
	fake.fetchApplicationReturns = struct {
		result1 *director.ApplicationOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDirectorClient) FetchApplicationReturnsOnCall(i int, result1 *director.ApplicationOutput, result2 error) {
	fake.fetchApplicationMutex.Lock()
	defer fake.fetchApplicationMutex.Unlock()
	fake.FetchApplicationStub = nil
	if fake.fetchApplicationReturnsOnCall == nil {
		fake.fetchApplicationReturnsOnCall = make(map[int]struct {
			result1 *director.ApplicationOutput
			result2 error
		})
	}
	fake.fetchApplicationReturnsOnCall[i] = struct {
		result1 *director.ApplicationOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDirectorClient) UpdateOperation(arg1 context.Context, arg2 *directora.Request) error {
	fake.updateOperationMutex.Lock()
	ret, specificReturn := fake.updateOperationReturnsOnCall[len(fake.updateOperationArgsForCall)]
	fake.updateOperationArgsForCall = append(fake.updateOperationArgsForCall, struct {
		arg1 context.Context
		arg2 *directora.Request
	}{arg1, arg2})
	fake.recordInvocation("UpdateOperation", []interface{}{arg1, arg2})
	fake.updateOperationMutex.Unlock()
	if fake.UpdateOperationStub != nil {
		return fake.UpdateOperationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateOperationReturns
	return fakeReturns.result1
}

func (fake *FakeDirectorClient) UpdateOperationCallCount() int {
	fake.updateOperationMutex.RLock()
	defer fake.updateOperationMutex.RUnlock()
	return len(fake.updateOperationArgsForCall)
}

func (fake *FakeDirectorClient) UpdateOperationCalls(stub func(context.Context, *directora.Request) error) {
	fake.updateOperationMutex.Lock()
	defer fake.updateOperationMutex.Unlock()
	fake.UpdateOperationStub = stub
}

func (fake *FakeDirectorClient) UpdateOperationArgsForCall(i int) (context.Context, *directora.Request) {
	fake.updateOperationMutex.RLock()
	defer fake.updateOperationMutex.RUnlock()
	argsForCall := fake.updateOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDirectorClient) UpdateOperationReturns(result1 error) {
	fake.updateOperationMutex.Lock()
	defer fake.updateOperationMutex.Unlock()
	fake.UpdateOperationStub = nil
	fake.updateOperationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirectorClient) UpdateOperationReturnsOnCall(i int, result1 error) {
	fake.updateOperationMutex.Lock()
	defer fake.updateOperationMutex.Unlock()
	fake.UpdateOperationStub = nil
	if fake.updateOperationReturnsOnCall == nil {
		fake.updateOperationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateOperationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirectorClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchApplicationMutex.RLock()
	defer fake.fetchApplicationMutex.RUnlock()
	fake.updateOperationMutex.RLock()
	defer fake.updateOperationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDirectorClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ controllers.DirectorClient = new(FakeDirectorClient)
