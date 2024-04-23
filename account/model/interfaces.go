package model

import (
	"context"

	"github.com/google/uuid"
)

// UserService defines methods the handler layer expects
// any service it interacts with to implement
// Service is expected to interact with Repository level
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
}

// UserRepository defines methods the service layer expects
// any repository it interacts with to implement
// Repository is expected to interact with DB layer
type UserRepository interface {
	FindByID(cts context.Context, uid uuid.UUID) (*User, error)
}
