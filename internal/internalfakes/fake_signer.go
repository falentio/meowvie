// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"meowvie/internal"
	"sync"
)

type FakeSigner struct {
	CompareStub        func(string, string) error
	compareMutex       sync.RWMutex
	compareArgsForCall []struct {
		arg1 string
		arg2 string
	}
	compareReturns struct {
		result1 error
	}
	compareReturnsOnCall map[int]struct {
		result1 error
	}
	SignStub        func(string) (string, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		arg1 string
	}
	signReturns struct {
		result1 string
		result2 error
	}
	signReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSigner) Compare(arg1 string, arg2 string) error {
	fake.compareMutex.Lock()
	ret, specificReturn := fake.compareReturnsOnCall[len(fake.compareArgsForCall)]
	fake.compareArgsForCall = append(fake.compareArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CompareStub
	fakeReturns := fake.compareReturns
	fake.recordInvocation("Compare", []interface{}{arg1, arg2})
	fake.compareMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSigner) CompareCallCount() int {
	fake.compareMutex.RLock()
	defer fake.compareMutex.RUnlock()
	return len(fake.compareArgsForCall)
}

func (fake *FakeSigner) CompareCalls(stub func(string, string) error) {
	fake.compareMutex.Lock()
	defer fake.compareMutex.Unlock()
	fake.CompareStub = stub
}

func (fake *FakeSigner) CompareArgsForCall(i int) (string, string) {
	fake.compareMutex.RLock()
	defer fake.compareMutex.RUnlock()
	argsForCall := fake.compareArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSigner) CompareReturns(result1 error) {
	fake.compareMutex.Lock()
	defer fake.compareMutex.Unlock()
	fake.CompareStub = nil
	fake.compareReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSigner) CompareReturnsOnCall(i int, result1 error) {
	fake.compareMutex.Lock()
	defer fake.compareMutex.Unlock()
	fake.CompareStub = nil
	if fake.compareReturnsOnCall == nil {
		fake.compareReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.compareReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSigner) Sign(arg1 string) (string, error) {
	fake.signMutex.Lock()
	ret, specificReturn := fake.signReturnsOnCall[len(fake.signArgsForCall)]
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SignStub
	fakeReturns := fake.signReturns
	fake.recordInvocation("Sign", []interface{}{arg1})
	fake.signMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSigner) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *FakeSigner) SignCalls(stub func(string) (string, error)) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = stub
}

func (fake *FakeSigner) SignArgsForCall(i int) string {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	argsForCall := fake.signArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSigner) SignReturns(result1 string, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSigner) SignReturnsOnCall(i int, result1 string, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	if fake.signReturnsOnCall == nil {
		fake.signReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.signReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSigner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.compareMutex.RLock()
	defer fake.compareMutex.RUnlock()
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSigner) recordInvocation(key string, args []interface{}) {
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

var _ internal.Signer = new(FakeSigner)
