// Code generated by counterfeiter. DO NOT EDIT.
package credsfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/creds"
	flags "github.com/jessevdk/go-flags"
)

type FakeManagerFactory struct {
	AddConfigStub        func(*flags.Group) creds.Manager
	addConfigMutex       sync.RWMutex
	addConfigArgsForCall []struct {
		arg1 *flags.Group
	}
	addConfigReturns struct {
		result1 creds.Manager
	}
	addConfigReturnsOnCall map[int]struct {
		result1 creds.Manager
	}
	NewInstanceStub        func(interface{}) (creds.Manager, error)
	newInstanceMutex       sync.RWMutex
	newInstanceArgsForCall []struct {
		arg1 interface{}
	}
	newInstanceReturns struct {
		result1 creds.Manager
		result2 error
	}
	newInstanceReturnsOnCall map[int]struct {
		result1 creds.Manager
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManagerFactory) AddConfig(arg1 *flags.Group) creds.Manager {
	fake.addConfigMutex.Lock()
	ret, specificReturn := fake.addConfigReturnsOnCall[len(fake.addConfigArgsForCall)]
	fake.addConfigArgsForCall = append(fake.addConfigArgsForCall, struct {
		arg1 *flags.Group
	}{arg1})
	fake.recordInvocation("AddConfig", []interface{}{arg1})
	fake.addConfigMutex.Unlock()
	if fake.AddConfigStub != nil {
		return fake.AddConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addConfigReturns
	return fakeReturns.result1
}

func (fake *FakeManagerFactory) AddConfigCallCount() int {
	fake.addConfigMutex.RLock()
	defer fake.addConfigMutex.RUnlock()
	return len(fake.addConfigArgsForCall)
}

func (fake *FakeManagerFactory) AddConfigCalls(stub func(*flags.Group) creds.Manager) {
	fake.addConfigMutex.Lock()
	defer fake.addConfigMutex.Unlock()
	fake.AddConfigStub = stub
}

func (fake *FakeManagerFactory) AddConfigArgsForCall(i int) *flags.Group {
	fake.addConfigMutex.RLock()
	defer fake.addConfigMutex.RUnlock()
	argsForCall := fake.addConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManagerFactory) AddConfigReturns(result1 creds.Manager) {
	fake.addConfigMutex.Lock()
	defer fake.addConfigMutex.Unlock()
	fake.AddConfigStub = nil
	fake.addConfigReturns = struct {
		result1 creds.Manager
	}{result1}
}

func (fake *FakeManagerFactory) AddConfigReturnsOnCall(i int, result1 creds.Manager) {
	fake.addConfigMutex.Lock()
	defer fake.addConfigMutex.Unlock()
	fake.AddConfigStub = nil
	if fake.addConfigReturnsOnCall == nil {
		fake.addConfigReturnsOnCall = make(map[int]struct {
			result1 creds.Manager
		})
	}
	fake.addConfigReturnsOnCall[i] = struct {
		result1 creds.Manager
	}{result1}
}

func (fake *FakeManagerFactory) NewInstance(arg1 interface{}) (creds.Manager, error) {
	fake.newInstanceMutex.Lock()
	ret, specificReturn := fake.newInstanceReturnsOnCall[len(fake.newInstanceArgsForCall)]
	fake.newInstanceArgsForCall = append(fake.newInstanceArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("NewInstance", []interface{}{arg1})
	fake.newInstanceMutex.Unlock()
	if fake.NewInstanceStub != nil {
		return fake.NewInstanceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newInstanceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManagerFactory) NewInstanceCallCount() int {
	fake.newInstanceMutex.RLock()
	defer fake.newInstanceMutex.RUnlock()
	return len(fake.newInstanceArgsForCall)
}

func (fake *FakeManagerFactory) NewInstanceCalls(stub func(interface{}) (creds.Manager, error)) {
	fake.newInstanceMutex.Lock()
	defer fake.newInstanceMutex.Unlock()
	fake.NewInstanceStub = stub
}

func (fake *FakeManagerFactory) NewInstanceArgsForCall(i int) interface{} {
	fake.newInstanceMutex.RLock()
	defer fake.newInstanceMutex.RUnlock()
	argsForCall := fake.newInstanceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManagerFactory) NewInstanceReturns(result1 creds.Manager, result2 error) {
	fake.newInstanceMutex.Lock()
	defer fake.newInstanceMutex.Unlock()
	fake.NewInstanceStub = nil
	fake.newInstanceReturns = struct {
		result1 creds.Manager
		result2 error
	}{result1, result2}
}

func (fake *FakeManagerFactory) NewInstanceReturnsOnCall(i int, result1 creds.Manager, result2 error) {
	fake.newInstanceMutex.Lock()
	defer fake.newInstanceMutex.Unlock()
	fake.NewInstanceStub = nil
	if fake.newInstanceReturnsOnCall == nil {
		fake.newInstanceReturnsOnCall = make(map[int]struct {
			result1 creds.Manager
			result2 error
		})
	}
	fake.newInstanceReturnsOnCall[i] = struct {
		result1 creds.Manager
		result2 error
	}{result1, result2}
}

func (fake *FakeManagerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addConfigMutex.RLock()
	defer fake.addConfigMutex.RUnlock()
	fake.newInstanceMutex.RLock()
	defer fake.newInstanceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManagerFactory) recordInvocation(key string, args []interface{}) {
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

var _ creds.ManagerFactory = new(FakeManagerFactory)
