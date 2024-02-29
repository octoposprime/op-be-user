package infrastructure

import (
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

var UserSortMap map[mo.UserSortField]string = map[mo.UserSortField]string{
	mo.UserSortFieldId:        "id",
	mo.UserSortFieldName:      "user_name",
	mo.UserSortFieldCreatedAt: "created_at",
	mo.UserSortFieldUpdatedAt: "updated_at",
}
