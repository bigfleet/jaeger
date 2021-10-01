// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	metrics "github.com/uber/jaeger-lib/metrics"
	zap "go.uber.org/zap"

	storage "github.com/jaegertracing/jaeger/storage"
	dependencystore "github.com/jaegertracing/jaeger/storage/dependencystore"
	spanstore "github.com/jaegertracing/jaeger/storage/spanstore"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// CreateDependencyReader provides a mock function with given fields:
func (_m *Factory) CreateDependencyReader() (dependencystore.Reader, error) {
	ret := _m.Called()

	var r0 dependencystore.Reader
	if rf, ok := ret.Get(0).(func() dependencystore.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dependencystore.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSpanReader provides a mock function with given fields:
func (_m *Factory) CreateSpanReader() (spanstore.Reader, error) {
	ret := _m.Called()

	var r0 spanstore.Reader
	if rf, ok := ret.Get(0).(func() spanstore.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spanstore.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSpanWriter provides a mock function with given fields:
func (_m *Factory) CreateSpanWriter() (spanstore.Writer, error) {
	ret := _m.Called()

	var r0 spanstore.Writer
	if rf, ok := ret.Get(0).(func() spanstore.Writer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spanstore.Writer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: metricsFactory, logger
func (_m *Factory) Initialize(metricsFactory metrics.Factory, logger *zap.Logger) error {
	ret := _m.Called(metricsFactory, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(metrics.Factory, *zap.Logger) error); ok {
		r0 = rf(metricsFactory, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ storage.Factory = (*Factory)(nil)