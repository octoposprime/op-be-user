package application

import (
	as "github.com/octoposprime/op-be-user/internal/application/service"
)

// QueryAdapter is an adapter for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type QueryAdapter struct {
	Service *as.Service
}

// NewQueryAdapter creates a new *QueryAdapter.
func NewQueryAdapter(s *as.Service) QueryAdapter {
	return QueryAdapter{
		s,
	}
}
