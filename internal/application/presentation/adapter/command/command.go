package application

import (
	as "github.com/octoposprime/op-be-user/internal/application/service"
)

// CommandAdapter is an adapter for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type CommandAdapter struct {
	Service *as.Service
}

// NewCommandAdapter creates a new *CommandAdapter.
func NewCommandAdapter(s *as.Service) CommandAdapter {
	return CommandAdapter{
		s,
	}
}
