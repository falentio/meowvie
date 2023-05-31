// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"meowvie/internal"
	"sync"

	"github.com/rs/xid"
)

type FakeDownloadUrlRepo struct {
	CreateStub        func(*internal.DownloadUrl) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *internal.DownloadUrl
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	CreateBatchStub        func([]*internal.DownloadUrl) error
	createBatchMutex       sync.RWMutex
	createBatchArgsForCall []struct {
		arg1 []*internal.DownloadUrl
	}
	createBatchReturns struct {
		result1 error
	}
	createBatchReturnsOnCall map[int]struct {
		result1 error
	}
	FindStub        func(xid.ID) (*internal.DownloadUrl, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 xid.ID
	}
	findReturns struct {
		result1 *internal.DownloadUrl
		result2 error
	}
	findReturnsOnCall map[int]struct {
		result1 *internal.DownloadUrl
		result2 error
	}
	FindByMovieIDStub        func(xid.ID) ([]*internal.DownloadUrl, error)
	findByMovieIDMutex       sync.RWMutex
	findByMovieIDArgsForCall []struct {
		arg1 xid.ID
	}
	findByMovieIDReturns struct {
		result1 []*internal.DownloadUrl
		result2 error
	}
	findByMovieIDReturnsOnCall map[int]struct {
		result1 []*internal.DownloadUrl
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDownloadUrlRepo) Create(arg1 *internal.DownloadUrl) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *internal.DownloadUrl
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownloadUrlRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeDownloadUrlRepo) CreateCalls(stub func(*internal.DownloadUrl) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeDownloadUrlRepo) CreateArgsForCall(i int) *internal.DownloadUrl {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownloadUrlRepo) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloadUrlRepo) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloadUrlRepo) CreateBatch(arg1 []*internal.DownloadUrl) error {
	var arg1Copy []*internal.DownloadUrl
	if arg1 != nil {
		arg1Copy = make([]*internal.DownloadUrl, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createBatchMutex.Lock()
	ret, specificReturn := fake.createBatchReturnsOnCall[len(fake.createBatchArgsForCall)]
	fake.createBatchArgsForCall = append(fake.createBatchArgsForCall, struct {
		arg1 []*internal.DownloadUrl
	}{arg1Copy})
	stub := fake.CreateBatchStub
	fakeReturns := fake.createBatchReturns
	fake.recordInvocation("CreateBatch", []interface{}{arg1Copy})
	fake.createBatchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownloadUrlRepo) CreateBatchCallCount() int {
	fake.createBatchMutex.RLock()
	defer fake.createBatchMutex.RUnlock()
	return len(fake.createBatchArgsForCall)
}

func (fake *FakeDownloadUrlRepo) CreateBatchCalls(stub func([]*internal.DownloadUrl) error) {
	fake.createBatchMutex.Lock()
	defer fake.createBatchMutex.Unlock()
	fake.CreateBatchStub = stub
}

func (fake *FakeDownloadUrlRepo) CreateBatchArgsForCall(i int) []*internal.DownloadUrl {
	fake.createBatchMutex.RLock()
	defer fake.createBatchMutex.RUnlock()
	argsForCall := fake.createBatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownloadUrlRepo) CreateBatchReturns(result1 error) {
	fake.createBatchMutex.Lock()
	defer fake.createBatchMutex.Unlock()
	fake.CreateBatchStub = nil
	fake.createBatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloadUrlRepo) CreateBatchReturnsOnCall(i int, result1 error) {
	fake.createBatchMutex.Lock()
	defer fake.createBatchMutex.Unlock()
	fake.CreateBatchStub = nil
	if fake.createBatchReturnsOnCall == nil {
		fake.createBatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createBatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloadUrlRepo) Find(arg1 xid.ID) (*internal.DownloadUrl, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 xid.ID
	}{arg1})
	stub := fake.FindStub
	fakeReturns := fake.findReturns
	fake.recordInvocation("Find", []interface{}{arg1})
	fake.findMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDownloadUrlRepo) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeDownloadUrlRepo) FindCalls(stub func(xid.ID) (*internal.DownloadUrl, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeDownloadUrlRepo) FindArgsForCall(i int) xid.ID {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownloadUrlRepo) FindReturns(result1 *internal.DownloadUrl, result2 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 *internal.DownloadUrl
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloadUrlRepo) FindReturnsOnCall(i int, result1 *internal.DownloadUrl, result2 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 *internal.DownloadUrl
			result2 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 *internal.DownloadUrl
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloadUrlRepo) FindByMovieID(arg1 xid.ID) ([]*internal.DownloadUrl, error) {
	fake.findByMovieIDMutex.Lock()
	ret, specificReturn := fake.findByMovieIDReturnsOnCall[len(fake.findByMovieIDArgsForCall)]
	fake.findByMovieIDArgsForCall = append(fake.findByMovieIDArgsForCall, struct {
		arg1 xid.ID
	}{arg1})
	stub := fake.FindByMovieIDStub
	fakeReturns := fake.findByMovieIDReturns
	fake.recordInvocation("FindByMovieID", []interface{}{arg1})
	fake.findByMovieIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDownloadUrlRepo) FindByMovieIDCallCount() int {
	fake.findByMovieIDMutex.RLock()
	defer fake.findByMovieIDMutex.RUnlock()
	return len(fake.findByMovieIDArgsForCall)
}

func (fake *FakeDownloadUrlRepo) FindByMovieIDCalls(stub func(xid.ID) ([]*internal.DownloadUrl, error)) {
	fake.findByMovieIDMutex.Lock()
	defer fake.findByMovieIDMutex.Unlock()
	fake.FindByMovieIDStub = stub
}

func (fake *FakeDownloadUrlRepo) FindByMovieIDArgsForCall(i int) xid.ID {
	fake.findByMovieIDMutex.RLock()
	defer fake.findByMovieIDMutex.RUnlock()
	argsForCall := fake.findByMovieIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownloadUrlRepo) FindByMovieIDReturns(result1 []*internal.DownloadUrl, result2 error) {
	fake.findByMovieIDMutex.Lock()
	defer fake.findByMovieIDMutex.Unlock()
	fake.FindByMovieIDStub = nil
	fake.findByMovieIDReturns = struct {
		result1 []*internal.DownloadUrl
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloadUrlRepo) FindByMovieIDReturnsOnCall(i int, result1 []*internal.DownloadUrl, result2 error) {
	fake.findByMovieIDMutex.Lock()
	defer fake.findByMovieIDMutex.Unlock()
	fake.FindByMovieIDStub = nil
	if fake.findByMovieIDReturnsOnCall == nil {
		fake.findByMovieIDReturnsOnCall = make(map[int]struct {
			result1 []*internal.DownloadUrl
			result2 error
		})
	}
	fake.findByMovieIDReturnsOnCall[i] = struct {
		result1 []*internal.DownloadUrl
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloadUrlRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.createBatchMutex.RLock()
	defer fake.createBatchMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	fake.findByMovieIDMutex.RLock()
	defer fake.findByMovieIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDownloadUrlRepo) recordInvocation(key string, args []interface{}) {
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

var _ internal.DownloadUrlRepo = new(FakeDownloadUrlRepo)