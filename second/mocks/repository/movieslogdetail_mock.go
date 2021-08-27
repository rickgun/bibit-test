package mocks

import (
	"bibit-test/models"

	"github.com/stretchr/testify/mock"
)

// MoviesLogDetail is an autogenerated mock type for the MoviesLogDetail type
type MoviesLogDetail struct {
	mock.Mock
}

// Detail provides a mock function with given fields: models.MoviesLogDetail
func (_m *MoviesLogDetail) Detail(mld *models.MoviesLogDetail) error {
	ret := _m.Called(mld)

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.MoviesLogDetail) error); ok {
		r1 = rf(mld)
	} else {
		r1 = ret.Error(1)
	}

	return r1
}
