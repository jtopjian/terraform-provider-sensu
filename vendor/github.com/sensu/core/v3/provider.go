package v3

import (
	"context"

	corev2 "github.com/sensu/core/v2"
)

// AuthProvider represents an abstracted authentication provider
type AuthProvider interface {
	Resource

	// Authenticate attempts to authenticate a user with its username and password
	Authenticate(ctx context.Context, username, password string) (*corev2.Claims, error)
	// Refresh renews the user claims with the provider claims
	Refresh(ctx context.Context, claims *corev2.Claims) (*corev2.Claims, error)

	// Name returns the provider name (e.g. default)
	Name() string
	// Type returns the provider type (e.g. ldap)
	Type() string
}
