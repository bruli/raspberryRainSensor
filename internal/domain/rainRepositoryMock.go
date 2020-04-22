// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package domain

import (
	"sync"
)

var (
	lockRainRepositoryMockRead sync.RWMutex
)

// Ensure, that RainRepositoryMock does implement RainRepository.
// If this is not the case, regenerate this file with moq.
var _ RainRepository = &RainRepositoryMock{}

// RainRepositoryMock is a mock implementation of RainRepository.
//
//     func TestSomethingThatUsesRainRepository(t *testing.T) {
//
//         // make and configure a mocked RainRepository
//         mockedRainRepository := &RainRepositoryMock{
//             ReadFunc: func() (uint16, error) {
// 	               panic("mock out the Read method")
//             },
//         }
//
//         // use mockedRainRepository in code that requires RainRepository
//         // and then make assertions.
//
//     }
type RainRepositoryMock struct {
	// ReadFunc mocks the Read method.
	ReadFunc func() (uint16, error)

	// calls tracks calls to the methods.
	calls struct {
		// Read holds details about calls to the Read method.
		Read []struct {
		}
	}
}

// Read calls ReadFunc.
func (mock *RainRepositoryMock) Read() (uint16, error) {
	if mock.ReadFunc == nil {
		panic("RainRepositoryMock.ReadFunc: method is nil but RainRepository.Read was just called")
	}
	callInfo := struct {
	}{}
	lockRainRepositoryMockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	lockRainRepositoryMockRead.Unlock()
	return mock.ReadFunc()
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedRainRepository.ReadCalls())
func (mock *RainRepositoryMock) ReadCalls() []struct {
} {
	var calls []struct {
	}
	lockRainRepositoryMockRead.RLock()
	calls = mock.calls.Read
	lockRainRepositoryMockRead.RUnlock()
	return calls
}
