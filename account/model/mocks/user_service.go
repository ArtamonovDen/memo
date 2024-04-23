package mocks

import (
	"context"

	"github.com/ArtamonovDen/memo/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock type for model.UserService
type MockUserService struct {
	mock.Mock
}

// Get is mock Get for UserService
func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	// args that will be passed to "Return" in the tests, when function
	// is called with a uid. Hence the name "ret"
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
