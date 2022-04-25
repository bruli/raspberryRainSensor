// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package app_test

import (
	"context"
	"sync"

	"github.com/bruli/raspberryRainSensor/internal/app"
	"github.com/bruli/raspberryRainSensor/internal/domain/rain"
)

// Ensure, that RainSensorMock does implement app.RainSensor.
// If this is not the case, regenerate this file with moq.
var _ app.RainSensor = &RainSensorMock{}

// RainSensorMock is a mock implementation of app.RainSensor.
//
// 	func TestSomethingThatUsesRainSensor(t *testing.T) {
//
// 		// make and configure a mocked app.RainSensor
// 		mockedRainSensor := &RainSensorMock{
// 			ReadFunc: func(ctx context.Context) (rain.Rain, error) {
// 				panic("mock out the Read method")
// 			},
// 		}
//
// 		// use mockedRainSensor in code that requires app.RainSensor
// 		// and then make assertions.
//
// 	}
type RainSensorMock struct {
	// ReadFunc mocks the Read method.
	ReadFunc func(ctx context.Context) (rain.Rain, error)

	// calls tracks calls to the methods.
	calls struct {
		// Read holds details about calls to the Read method.
		Read []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockRead sync.RWMutex
}

// Read calls ReadFunc.
func (mock *RainSensorMock) Read(ctx context.Context) (rain.Rain, error) {
	if mock.ReadFunc == nil {
		panic("RainSensorMock.ReadFunc: method is nil but RainSensor.Read was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	mock.lockRead.Unlock()
	return mock.ReadFunc(ctx)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedRainSensor.ReadCalls())
func (mock *RainSensorMock) ReadCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockRead.RLock()
	calls = mock.calls.Read
	mock.lockRead.RUnlock()
	return calls
}
