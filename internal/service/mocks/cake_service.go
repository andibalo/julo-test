// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "julo-test/internal/storage/entity"

	mock "github.com/stretchr/testify/mock"

	request "julo-test/internal/request"

	response "julo-test/internal/response"
)

// CakeService is an autogenerated mock type for the CakeService type
type CakeService struct {
	mock.Mock
}

// CreateCake provides a mock function with given fields: createUserReq
func (_m *CakeService) CreateCake(createUserReq *request.CreateCakeRequest) (response.Code, error) {
	ret := _m.Called(createUserReq)

	var r0 response.Code
	if rf, ok := ret.Get(0).(func(*request.CreateCakeRequest) response.Code); ok {
		r0 = rf(createUserReq)
	} else {
		r0 = ret.Get(0).(response.Code)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*request.CreateCakeRequest) error); ok {
		r1 = rf(createUserReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCakeByID provides a mock function with given fields: id
func (_m *CakeService) DeleteCakeByID(id string) (response.Code, error) {
	ret := _m.Called(id)

	var r0 response.Code
	if rf, ok := ret.Get(0).(func(string) response.Code); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(response.Code)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCake provides a mock function with given fields: page, perPage
func (_m *CakeService) GetAllCake(page int, perPage int) (response.Code, []entity.Cake, int64, error) {
	ret := _m.Called(page, perPage)

	var r0 response.Code
	if rf, ok := ret.Get(0).(func(int, int) response.Code); ok {
		r0 = rf(page, perPage)
	} else {
		r0 = ret.Get(0).(response.Code)
	}

	var r1 []entity.Cake
	if rf, ok := ret.Get(1).(func(int, int) []entity.Cake); ok {
		r1 = rf(page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]entity.Cake)
		}
	}

	var r2 int64
	if rf, ok := ret.Get(2).(func(int, int) int64); ok {
		r2 = rf(page, perPage)
	} else {
		r2 = ret.Get(2).(int64)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(int, int) error); ok {
		r3 = rf(page, perPage)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetCakeByID provides a mock function with given fields: id
func (_m *CakeService) GetCakeByID(id string) (response.Code, *entity.Cake, error) {
	ret := _m.Called(id)

	var r0 response.Code
	if rf, ok := ret.Get(0).(func(string) response.Code); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(response.Code)
	}

	var r1 *entity.Cake
	if rf, ok := ret.Get(1).(func(string) *entity.Cake); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*entity.Cake)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateCakeByID provides a mock function with given fields: req, id
func (_m *CakeService) UpdateCakeByID(req *request.UpdateCakeRequest, id string) (response.Code, error) {
	ret := _m.Called(req, id)

	var r0 response.Code
	if rf, ok := ret.Get(0).(func(*request.UpdateCakeRequest, string) response.Code); ok {
		r0 = rf(req, id)
	} else {
		r0 = ret.Get(0).(response.Code)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*request.UpdateCakeRequest, string) error); ok {
		r1 = rf(req, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCakeService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCakeService creates a new instance of CakeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCakeService(t mockConstructorTestingTNewCakeService) *CakeService {
	mock := &CakeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}