package mocks

import (
	"context"

	"github.com/ArtamonovDen/memo/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserReposotory struct {
	mock.Mock
}

func (m *MockUserReposotory) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	ret := m.Called(ctx, uid)

	// first value passed to "Return"
	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
