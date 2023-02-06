package user

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
)

type User struct {
	collection repositories.IUserRepository
}

// NewUserService creates a service use to manage the user logic.
func NewUserService(repo repositories.IUserRepository) *User {
	return &User{
		collection: repo,
	}
}

// Creates a new User.
func (c *User) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	return c.collection.Create(ctx, user)
}
