package model

import "github.com/google/uuid"

// UserService defines methods the handler layer expects
// any service it interacts with to implement
// Service is expected to interact with Repository level
type UserService interface {
	Get(uid uuid.UUID) (*User, error)
}

// UserRepository defines methods the service layer expects
// any repository it interacts with to implement
// Repository is expected to interact with DB layer
type UserRepository interface {
	FindByID(uid uuid.UUID) (*User, error)
}
