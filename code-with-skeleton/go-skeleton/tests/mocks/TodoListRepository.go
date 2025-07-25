// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/repository/mysql/entity"
	mock "github.com/stretchr/testify/mock"

	mysql "github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/repository/mysql"
)

// TodoListRepository is an autogenerated mock type for the TodoListRepository type
type TodoListRepository struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *TodoListRepository) Begin() (mysql.TrxObj, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 mysql.TrxObj
	var r1 error
	if rf, ok := ret.Get(0).(func() (mysql.TrxObj, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() mysql.TrxObj); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mysql.TrxObj)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, dbTrx, params, nonZeroVal
func (_m *TodoListRepository) Create(ctx context.Context, dbTrx mysql.TrxObj, params *entity.TodoList, nonZeroVal bool) error {
	ret := _m.Called(ctx, dbTrx, params, nonZeroVal)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, *entity.TodoList, bool) error); ok {
		r0 = rf(ctx, dbTrx, params, nonZeroVal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: ctx, dbTrx, id
func (_m *TodoListRepository) DeleteByID(ctx context.Context, dbTrx mysql.TrxObj, id int64) error {
	ret := _m.Called(ctx, dbTrx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, int64) error); ok {
		r0 = rf(ctx, dbTrx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *TodoListRepository) GetByID(ctx context.Context, ID int64) (*entity.TodoList, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *entity.TodoList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.TodoList, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.TodoList); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TodoList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: ctx, ID
func (_m *TodoListRepository) GetByUserID(ctx context.Context, ID int64) ([]*entity.TodoList, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserID")
	}

	var r0 []*entity.TodoList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*entity.TodoList, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*entity.TodoList); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.TodoList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockByID provides a mock function with given fields: ctx, dbTrx, ID
func (_m *TodoListRepository) LockByID(ctx context.Context, dbTrx mysql.TrxObj, ID int64) (*entity.TodoList, error) {
	ret := _m.Called(ctx, dbTrx, ID)

	if len(ret) == 0 {
		panic("no return value specified for LockByID")
	}

	var r0 *entity.TodoList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, int64) (*entity.TodoList, error)); ok {
		return rf(ctx, dbTrx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, int64) *entity.TodoList); ok {
		r0 = rf(ctx, dbTrx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TodoList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, mysql.TrxObj, int64) error); ok {
		r1 = rf(ctx, dbTrx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, dbTrx, params, changes
func (_m *TodoListRepository) Update(ctx context.Context, dbTrx mysql.TrxObj, params *entity.TodoList, changes *entity.TodoList) error {
	ret := _m.Called(ctx, dbTrx, params, changes)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.TrxObj, *entity.TodoList, *entity.TodoList) error); ok {
		r0 = rf(ctx, dbTrx, params, changes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTodoListRepository creates a new instance of TodoListRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTodoListRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TodoListRepository {
	mock := &TodoListRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
