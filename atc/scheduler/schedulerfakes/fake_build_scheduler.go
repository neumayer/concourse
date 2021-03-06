// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/scheduler"
	"github.com/concourse/concourse/atc/scheduler/algorithm"
)

type FakeBuildScheduler struct {
	ScheduleStub        func(context.Context, lager.Logger, db.Pipeline, db.Job, db.Resources, algorithm.NameToIDMap) (bool, error)
	scheduleMutex       sync.RWMutex
	scheduleArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.Pipeline
		arg4 db.Job
		arg5 db.Resources
		arg6 algorithm.NameToIDMap
	}
	scheduleReturns struct {
		result1 bool
		result2 error
	}
	scheduleReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildScheduler) Schedule(arg1 context.Context, arg2 lager.Logger, arg3 db.Pipeline, arg4 db.Job, arg5 db.Resources, arg6 algorithm.NameToIDMap) (bool, error) {
	fake.scheduleMutex.Lock()
	ret, specificReturn := fake.scheduleReturnsOnCall[len(fake.scheduleArgsForCall)]
	fake.scheduleArgsForCall = append(fake.scheduleArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.Pipeline
		arg4 db.Job
		arg5 db.Resources
		arg6 algorithm.NameToIDMap
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("Schedule", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.scheduleMutex.Unlock()
	if fake.ScheduleStub != nil {
		return fake.ScheduleStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.scheduleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildScheduler) ScheduleCallCount() int {
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	return len(fake.scheduleArgsForCall)
}

func (fake *FakeBuildScheduler) ScheduleCalls(stub func(context.Context, lager.Logger, db.Pipeline, db.Job, db.Resources, algorithm.NameToIDMap) (bool, error)) {
	fake.scheduleMutex.Lock()
	defer fake.scheduleMutex.Unlock()
	fake.ScheduleStub = stub
}

func (fake *FakeBuildScheduler) ScheduleArgsForCall(i int) (context.Context, lager.Logger, db.Pipeline, db.Job, db.Resources, algorithm.NameToIDMap) {
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	argsForCall := fake.scheduleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeBuildScheduler) ScheduleReturns(result1 bool, result2 error) {
	fake.scheduleMutex.Lock()
	defer fake.scheduleMutex.Unlock()
	fake.ScheduleStub = nil
	fake.scheduleReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildScheduler) ScheduleReturnsOnCall(i int, result1 bool, result2 error) {
	fake.scheduleMutex.Lock()
	defer fake.scheduleMutex.Unlock()
	fake.ScheduleStub = nil
	if fake.scheduleReturnsOnCall == nil {
		fake.scheduleReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.scheduleReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildScheduler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildScheduler) recordInvocation(key string, args []interface{}) {
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

var _ scheduler.BuildScheduler = new(FakeBuildScheduler)
