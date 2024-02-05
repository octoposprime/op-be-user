package application

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	tseed "github.com/octoposprime/op-be-user/tool/config"
)

// This is the migration process of the seed layer.
func (a Service) Migrate() Service {
	for _, seedUser := range tseed.GetSeedConfigInstance().Users {
		var seedUserFilter me.UserFilter
		seedUserFilter.UserName = seedUser.User.UserName
		inUsers, err := a.GetUsersByFilter(context.TODO(), seedUserFilter)
		if err != nil {
			panic(err)
		}
		if inUsers.TotalRows == 0 {
			inUser, err := a.SaveUser(context.TODO(), seedUser.User)
			if err != nil {
				panic(err)
			}
			userPasswordEntity := me.NewUserPassword(
				uuid.UUID{},
				inUser.Id,
				seedUser.UserPassword,
			)
			err = a.ChangePassword(context.TODO(), *userPasswordEntity)
			if err != nil {
				panic(err)
			}
		}
	}
	return a
}
