package auth_ports

import "github.com/khanakia/mangobp/pkg/auth/auth_domain"

type UserRepo interface {
	// Get(id string) (auth_domain.User, error)
	FindByEmail(email string) (*auth_domain.User, error)
	Register(user *auth_domain.User) (*auth_domain.User, error)
	// Save(auth_domain.User) error
}
