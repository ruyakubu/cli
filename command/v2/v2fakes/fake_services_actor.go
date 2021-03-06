// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeServicesActor struct {
	GetServiceInstancesSummaryBySpaceStub        func(spaceGUID string) ([]v2action.ServiceInstanceSummary, v2action.Warnings, error)
	getServiceInstancesSummaryBySpaceMutex       sync.RWMutex
	getServiceInstancesSummaryBySpaceArgsForCall []struct {
		spaceGUID string
	}
	getServiceInstancesSummaryBySpaceReturns struct {
		result1 []v2action.ServiceInstanceSummary
		result2 v2action.Warnings
		result3 error
	}
	getServiceInstancesSummaryBySpaceReturnsOnCall map[int]struct {
		result1 []v2action.ServiceInstanceSummary
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServicesActor) GetServiceInstancesSummaryBySpace(spaceGUID string) ([]v2action.ServiceInstanceSummary, v2action.Warnings, error) {
	fake.getServiceInstancesSummaryBySpaceMutex.Lock()
	ret, specificReturn := fake.getServiceInstancesSummaryBySpaceReturnsOnCall[len(fake.getServiceInstancesSummaryBySpaceArgsForCall)]
	fake.getServiceInstancesSummaryBySpaceArgsForCall = append(fake.getServiceInstancesSummaryBySpaceArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("GetServiceInstancesSummaryBySpace", []interface{}{spaceGUID})
	fake.getServiceInstancesSummaryBySpaceMutex.Unlock()
	if fake.GetServiceInstancesSummaryBySpaceStub != nil {
		return fake.GetServiceInstancesSummaryBySpaceStub(spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceInstancesSummaryBySpaceReturns.result1, fake.getServiceInstancesSummaryBySpaceReturns.result2, fake.getServiceInstancesSummaryBySpaceReturns.result3
}

func (fake *FakeServicesActor) GetServiceInstancesSummaryBySpaceCallCount() int {
	fake.getServiceInstancesSummaryBySpaceMutex.RLock()
	defer fake.getServiceInstancesSummaryBySpaceMutex.RUnlock()
	return len(fake.getServiceInstancesSummaryBySpaceArgsForCall)
}

func (fake *FakeServicesActor) GetServiceInstancesSummaryBySpaceArgsForCall(i int) string {
	fake.getServiceInstancesSummaryBySpaceMutex.RLock()
	defer fake.getServiceInstancesSummaryBySpaceMutex.RUnlock()
	return fake.getServiceInstancesSummaryBySpaceArgsForCall[i].spaceGUID
}

func (fake *FakeServicesActor) GetServiceInstancesSummaryBySpaceReturns(result1 []v2action.ServiceInstanceSummary, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstancesSummaryBySpaceStub = nil
	fake.getServiceInstancesSummaryBySpaceReturns = struct {
		result1 []v2action.ServiceInstanceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesActor) GetServiceInstancesSummaryBySpaceReturnsOnCall(i int, result1 []v2action.ServiceInstanceSummary, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstancesSummaryBySpaceStub = nil
	if fake.getServiceInstancesSummaryBySpaceReturnsOnCall == nil {
		fake.getServiceInstancesSummaryBySpaceReturnsOnCall = make(map[int]struct {
			result1 []v2action.ServiceInstanceSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceInstancesSummaryBySpaceReturnsOnCall[i] = struct {
		result1 []v2action.ServiceInstanceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServiceInstancesSummaryBySpaceMutex.RLock()
	defer fake.getServiceInstancesSummaryBySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServicesActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.ServicesActor = new(FakeServicesActor)
