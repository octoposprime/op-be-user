package application

import (
	"context"

	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
)

// GetUsersByFilter returns the users that match the given filter.
func (a QueryAdapter) GetUsersByFilter(ctx context.Context, userFilter me.UserFilter) (me.Users, error) {
	return a.Service.GetUsersByFilter(ctx, userFilter)
}
