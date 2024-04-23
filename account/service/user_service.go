package service

import (
	"context"

	"github.com/ArtamonovDen/memo/model"
	"github.com/google/uuid"
)

// UserService acts as a struct for injecting implementation of UserRepository
// for use in methods
type UserService struct {
	UserRepository model.UserRepository
}

func (s *UserService) Get(ctx context.Context, uuid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uuid)
	return u, err
}

// USConfig will hold repositories that will eventually be injected into this
// service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory to create new user service by config
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}
