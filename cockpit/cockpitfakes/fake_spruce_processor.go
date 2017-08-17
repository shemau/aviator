// Code generated by counterfeiter. DO NOT EDIT.
package cockpitfakes

import (
	"sync"

	"github.com/JulzDiverse/aviator/cockpit"
)

type FakeSpruceProcessor struct {
	ProcessStub        func([]cockpit.Spruce) ([]byte, error)
	processMutex       sync.RWMutex
	processArgsForCall []struct {
		arg1 []cockpit.Spruce
	}
	processReturns struct {
		result1 []byte
		result2 error
	}
	processReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpruceProcessor) Process(arg1 []cockpit.Spruce) ([]byte, error) {
	var arg1Copy []cockpit.Spruce
	if arg1 != nil {
		arg1Copy = make([]cockpit.Spruce, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.processMutex.Lock()
	ret, specificReturn := fake.processReturnsOnCall[len(fake.processArgsForCall)]
	fake.processArgsForCall = append(fake.processArgsForCall, struct {
		arg1 []cockpit.Spruce
	}{arg1Copy})
	fake.recordInvocation("Process", []interface{}{arg1Copy})
	fake.processMutex.Unlock()
	if fake.ProcessStub != nil {
		return fake.ProcessStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.processReturns.result1, fake.processReturns.result2
}

func (fake *FakeSpruceProcessor) ProcessCallCount() int {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return len(fake.processArgsForCall)
}

func (fake *FakeSpruceProcessor) ProcessArgsForCall(i int) []cockpit.Spruce {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return fake.processArgsForCall[i].arg1
}

func (fake *FakeSpruceProcessor) ProcessReturns(result1 []byte, result2 error) {
	fake.ProcessStub = nil
	fake.processReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeSpruceProcessor) ProcessReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.ProcessStub = nil
	if fake.processReturnsOnCall == nil {
		fake.processReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.processReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeSpruceProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpruceProcessor) recordInvocation(key string, args []interface{}) {
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

var _ cockpit.SpruceProcessor = new(FakeSpruceProcessor)