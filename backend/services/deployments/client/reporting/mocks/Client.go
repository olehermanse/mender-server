// Copyright 2022 Northern.tech AS
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	    http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.

// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mendersoftware/mender-server/services/deployments/model"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CheckHealth provides a mock function with given fields: ctx
func (_m *Client) CheckHealth(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: ctx, tenantId, searchParams
func (_m *Client) Search(ctx context.Context, tenantId string, searchParams model.SearchParams) ([]model.InvDevice, int, error) {
	ret := _m.Called(ctx, tenantId, searchParams)

	var r0 []model.InvDevice
	if rf, ok := ret.Get(0).(func(context.Context, string, model.SearchParams) []model.InvDevice); ok {
		r0 = rf(ctx, tenantId, searchParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.InvDevice)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, model.SearchParams) int); ok {
		r1 = rf(ctx, tenantId, searchParams)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, model.SearchParams) error); ok {
		r2 = rf(ctx, tenantId, searchParams)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
